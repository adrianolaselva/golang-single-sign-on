package oauth

import (
	"encoding/base64"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/schema"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"net/url"
	"oauth2/src/common"
	"oauth2/src/dto"
	"oauth2/src/enums"
	"oauth2/src/models"
	"oauth2/src/repository"
	"strings"
	"time"
)

type AuthTokenClaim struct {
	*jwt.StandardClaims
	Profile 	models.User 	`json:"profile,omitempty"`
	ClientID 	string			`json:"client_id,omitempty"`
	Scope		[]string		`json:"scope,omitempty"`
}

type AuthFlow interface {
	SetRequest(r *http.Request) error
	SetExpiresAt(minutes int)
	GetAccessToken() (*dto.AccessTokenResponseDTO, error)
	GetAuthorizationCode() (*string, error)
	verifyCredentials() error
	verifyClientAndSecret() error
	verifyClient() error
	validateScopes() error
	getAccessTokenBase() (*models.AccessToken, error)
	getRefreshTokenBase(accessToken *models.AccessToken, accessTokenResponseDTO *dto.AccessTokenResponseDTO) error
	grantTypePassword() (*dto.AccessTokenResponseDTO, error)
	grantTypeClientCredentials() (*dto.AccessTokenResponseDTO, error)
	grantTypeRefreshToken() (*dto.AccessTokenResponseDTO, error)
}

type authFlow struct {
	token *jwt.Token
	accessTokenRequest *dto.AccessTokenRequestDTO
	accessTokenExpiresAt int64
	refreshTokenExpiresAt int64
	signature string
	tokenType string
	accessToken string
	hash common.Hash
	user *models.User
	client *models.Client
	userRepository repository.UserRepository
	clientRepository repository.ClientRepository
	authCodeRepository repository.AuthCodeRepository
	refreshTokenRepository repository.RefreshTokenRepository
	accessTokenRepository repository.AccessTokenRepository
}

// NewAuthFlow: create instance authentication flow
func NewAuthFlow(
		hmac *jwt.SigningMethodHMAC,
		signature string,
		userRepository repository.UserRepository,
		clientRepository repository.ClientRepository,
		authCodeRepository repository.AuthCodeRepository,
		refreshTokenRepository repository.RefreshTokenRepository,
		accessTokenRepository repository.AccessTokenRepository) *authFlow {
	return &authFlow{
		token: jwt.New(hmac),
		signature: signature,
		hash: common.NewHash(),
		userRepository: userRepository,
		clientRepository: clientRepository,
		authCodeRepository: authCodeRepository,
		refreshTokenRepository: refreshTokenRepository,
		accessTokenRepository: accessTokenRepository,
	}
}

// SetRequest: set request
func (o *authFlow) SetRequest(r *http.Request) error {

	if err := r.ParseForm(); err != nil {
		log.Println(err.Error())
	}

	values := make(url.Values)

	for key := range r.Form {
		values.Set(key, r.Form.Get(key))
	}

	for key := range r.PostForm {
		values.Set(key, r.PostForm.Get(key))
	}

	for key := range r.URL.Query() {
		values.Set(key, r.URL.Query().Get(key))
	}

	for key := range values {
		r.Form.Set(key, values.Get(key))
	}

	auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	if len(auth) == 2 || auth[0] == "Basic" {
		payload, _ := base64.StdEncoding.DecodeString(auth[1])
		pair := strings.SplitN(string(payload), ":", 2)
		o.user = &models.User{
			Username: pair[0],
			Password: &pair[1],
		}
	}

	if len(auth) == 2 || auth[0] == "Bearer" {
		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(auth[1], claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(o.signature), nil
		})
		if err == nil {
			if err = claims.Valid(); err != nil {
				return errors.Errorf(err.Error())
			}

			o.token = token

			o.tokenType = auth[0]
			o.accessToken = auth[1]
		}
	}

	accessTokenRequest := dto.AccessTokenRequestDTO{}
	if err := schema.NewDecoder().Decode(&accessTokenRequest, values); err != nil {
		return err
	}
	o.accessTokenRequest = &accessTokenRequest

	if accessTokenRequest.GrantType == enums.GrantTypeClientCredentials.String() {
		if o.accessTokenRequest.ClientID == "" {
			o.accessTokenRequest.ClientID = o.user.Username
			o.accessTokenRequest.ClientSecret = *o.user.Password
		}
	}

	return nil
}

// SetExpiresAt: set time expires
func (o *authFlow) SetExpiresAt(minutes int) {
	o.accessTokenExpiresAt = time.Now().Add(time.Minute * time.Duration(minutes)).Unix()
	o.refreshTokenExpiresAt = time.Now().Add(time.Minute * time.Duration(minutes * 120)).Unix()
}

// validateScopes: verify scopes
func (o *authFlow) validateScopes() error {
	if len(o.accessTokenRequest.Scope) > 0 {
		o.client.Scopes = strings.Replace(o.client.Scopes, " ", ",", 10000)
		clientScopes := strings.Split(o.client.Scopes, ",")
		requestScopes := strings.Split(o.accessTokenRequest.Scope, " ")

		for _, requestScope := range requestScopes {
			checked := false

			for _, scope := range clientScopes {
				if scope == requestScope {
					checked = true
					break
				}
			}

			if !checked {
				return errors.Errorf("scope %v not found for client", requestScope)
			}
		}
	}

	return nil
}

// verifyClient: verify client
func (o *authFlow) verifyClient() error {
	client, err := o.clientRepository.FindById(o.accessTokenRequest.ClientID)
	if err != nil {
		return errors.Errorf("invalid client_id")
	}

	if client.Revoked {
		return errors.Errorf("revoked client_id")
	}

	o.client = client

	return nil
}

// verifyCredentials: verify credentials
func (o *authFlow) verifyCredentials() error {
	user, err := o.userRepository.FindByUsername(o.user.Username)
	if err != nil {
		return errors.Errorf("username not found")
	}

	result, err := o.hash.BCryptCompare(*user.Password, *o.user.Password)
	if err != nil || !result {
		return  errors.Errorf("invalid credentials")
	}

	o.user = user

	return nil
}

// verifyClientAndSecret: verify client and secret
func (o *authFlow) verifyClientAndSecret() error {
	client, err := o.clientRepository.FindById(o.accessTokenRequest.ClientID)
	if err != nil {
		return errors.Errorf("invalid client_id")
	}

	if client.Secret != o.accessTokenRequest.ClientSecret  {
		return errors.Errorf("client_secret is invalid")
	}

	if client.Revoked {
		return errors.Errorf("revoked client_id")
	}

	o.client = client
	o.user = client.User

	return nil
}

// getAccessTokenBase: load base access token
func (o *authFlow) getAccessTokenBase() (*models.AccessToken, error) {
	o.token.Claims = &AuthTokenClaim{
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: o.accessTokenExpiresAt,
		},
		Profile: *o.user,
		ClientID: o.client.ID,
		Scope: strings.Split(o.accessTokenRequest.Scope, " "),
	}

	token, err := o.token.SignedString([]byte(o.signature))
	if err != nil {
		return nil, errors.Errorf("failed to generate token %s", err.Error())
	}

	expiresAt := time.Unix(o.accessTokenExpiresAt, 0)
	accessToken := models.AccessToken{
		User: o.user,
		Client: o.client,
		AccessToken: token,
		ExpiresAt: &expiresAt,
		Revoked: false,
		Scopes: o.accessTokenRequest.Scope,
	}

	return &accessToken, nil
}

// getRefreshTokenBase: generate refresh_token
func (o *authFlow) getRefreshTokenBase(accessToken *models.AccessToken, accessTokenResponseDTO *dto.AccessTokenResponseDTO) error {
	refreshTokenAux := o.token

	refreshTokenAux.Claims = &AuthTokenClaim{
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: o.refreshTokenExpiresAt,
		},
	}

	token, err := refreshTokenAux.SignedString([]byte(o.signature))
	if err != nil {
		return errors.Errorf("failed to generate refresh_token %s", err.Error())
	}

	expiresAt := time.Unix(o.refreshTokenExpiresAt, 0)
	refreshToken := models.RefreshToken{
		RefreshToken: token,
		Revoked:      false,
		ExpiresAt:    &expiresAt,
		AccessToken:  accessToken,
	}

	err = o.refreshTokenRepository.Create(&refreshToken)
	if err != nil {
		return errors.Errorf("failed to generate refresh_token")
	}

	accessTokenResponseDTO.RefreshToken = &refreshToken.RefreshToken

	return nil
}

// GetAccessToken: get access_token for grant_type
func (o *authFlow) GetAccessToken() (*dto.AccessTokenResponseDTO, error) {
	switch o.accessTokenRequest.GrantType {
		case enums.GrantTypePassword.String():
			return o.grantTypePassword()
		case enums.GrantTypeClientCredentials.String():
			return o.grantTypeClientCredentials()
		case enums.GrantTypeRefreshToken.String():
			return o.grantTypeRefreshToken()
		case enums.GrantTypeAuthorizationCode.String():
			return o.grantTypeAuthorizationCode()
		case enums.GrantTypeImplicit.String():
			return o.grantTypeImplicit()
		default:
			return nil, errors.Errorf("invalid grant_type %s", o.accessTokenRequest.GrantType)
	}
}

// grantTypePassword: flow password
func (o *authFlow) grantTypePassword() (*dto.AccessTokenResponseDTO, error) {

	err := o.verifyCredentials()
	if err != nil {
		return nil, err
	}

	err = o.verifyClient()
	if err != nil {
		return nil, err
	}

	err = o.validateScopes()
	if err != nil {
		return nil, err
	}

	accessToken, err := o.getAccessTokenBase()
	if err != nil {
		return nil, err
	}

	err = o.accessTokenRepository.Create(accessToken)
	if err != nil {
		return nil, errors.Errorf("failed to generate access_token")
	}

	accessTokenResponseDTO := &dto.AccessTokenResponseDTO{
		TokenType:   "Bearer",
		ExpiresIn:   o.accessTokenExpiresAt,
		AccessToken: accessToken.AccessToken,
	}

	err = o.getRefreshTokenBase(accessToken, accessTokenResponseDTO)
	if err != nil {
		return nil, errors.Errorf("failed to generate access_token")
	}

	return accessTokenResponseDTO, nil
}

// grantTypeClientCredentials: flow client_credentials
func (o *authFlow) grantTypeClientCredentials() (*dto.AccessTokenResponseDTO, error) {

	err := o.verifyClientAndSecret()
	if err != nil {
		return nil, err
	}

	err = o.validateScopes()
	if err != nil {
		return nil, err
	}

	accessToken, err := o.getAccessTokenBase()
	if err != nil {
		return nil, err
	}

	err = o.accessTokenRepository.Create(accessToken)
	if err != nil {
		return nil, errors.Errorf("failed to generate access_token")
	}

	accessTokenResponseDTO := dto.AccessTokenResponseDTO{
		TokenType:   "Bearer",
		ExpiresIn:   o.accessTokenExpiresAt,
		AccessToken: accessToken.AccessToken,
	}

	return &accessTokenResponseDTO, nil
}

// grantTypeRefreshToken: flow refresh_token
func (o *authFlow) grantTypeRefreshToken() (*dto.AccessTokenResponseDTO, error) {

	refreshToken, err := o.refreshTokenRepository.FindByRefreshToken(o.accessTokenRequest.RefreshToken)
	if err != nil {
		return nil, errors.Errorf("refresh_token not found")
	}


	requestAccessToken, err := o.accessTokenRepository.FindByAccessToken(o.accessToken)
	if err != nil {
		return nil, errors.Errorf("access_token not found")
	}

	if requestAccessToken.Revoked {
		return nil, errors.Errorf("access_token has revoked")
	}

	if refreshToken.Revoked {
		return nil, errors.Errorf("refresh_token has revoked")
	}

	o.user = refreshToken.AccessToken.User
	o.client = refreshToken.AccessToken.Client

	refreshToken.Revoked = true
	refreshToken.AccessToken.Revoked = true
	o.accessTokenRequest.Scope = refreshToken.AccessToken.Scopes

	err = o.refreshTokenRepository.Update(refreshToken)
	if err != nil {
		return nil, errors.Errorf("failed to revoke token and refresh_token")
	}

	accessToken, err := o.getAccessTokenBase()
	if err != nil {
		return nil, err
	}

	err = o.accessTokenRepository.Create(accessToken)
	if err != nil {
		return nil, errors.Errorf("failed to generate access_token")
	}

	accessTokenResponseDTO := dto.AccessTokenResponseDTO{
		TokenType:   "Bearer",
		ExpiresIn:   o.accessTokenExpiresAt,
		AccessToken: accessToken.AccessToken,
	}

	err = o.getRefreshTokenBase(accessToken, &accessTokenResponseDTO)
	if err != nil {
		return nil, err
	}

	return &accessTokenResponseDTO, nil
}




func (o *authFlow) GetAuthorizationCode() (*string, error) {
	return nil, nil
}

func (o *authFlow) grantTypeAuthorizationCode() (*dto.AccessTokenResponseDTO, error) {
	o.token.Claims = &AuthTokenClaim{
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: o.accessTokenExpiresAt,
		},
	}

	accessToken, err := o.token.SignedString([]byte(o.signature))
	if err != nil {
		return nil, errors.Errorf("failed to generate token %s", err.Error())
	}

	return &dto.AccessTokenResponseDTO{
		TokenType:   "Bearer",
		ExpiresIn:   o.accessTokenExpiresAt,
		AccessToken: accessToken,
	}, nil
}

func (o *authFlow) grantTypeImplicit() (*dto.AccessTokenResponseDTO, error) {
	o.token.Claims = &AuthTokenClaim{
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: o.accessTokenExpiresAt,
		},
	}

	accessToken, err := o.token.SignedString([]byte(o.signature))
	if err != nil {
		return nil, errors.Errorf("failed to generate token %s", err.Error())
	}

	return &dto.AccessTokenResponseDTO{
		TokenType:   "Bearer",
		ExpiresIn:   o.accessTokenExpiresAt,
		RefreshToken: &accessToken,
		AccessToken: accessToken,
	}, nil
}
