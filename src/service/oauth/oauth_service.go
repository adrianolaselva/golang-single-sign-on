package oauth

import (
	"encoding/base64"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/schema"
	"github.com/pkg/errors"
	"log"
	"math/rand"
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
	ValidateParameters() error
	SetExpiresAt(minutes int)
	GetAccessToken() (*dto.AccessTokenResponseDTO, error)
	GetAuthorizationCode() (*string, error)
	Login(loginDTO dto.LoginDTO) (*dto.LoginResponseDTO, error)
	verifyCredentials() error
	verifyClientAndSecret() error
	verifyClient() error
	verifyAuthCode() error
	validateScopes() error
	getAccessTokenBase() (*models.AccessToken, error)
	getRefreshTokenBase(accessToken *models.AccessToken, accessTokenResponseDTO *dto.AccessTokenResponseDTO) error
	grantTypePassword() (*dto.AccessTokenResponseDTO, error)
	grantTypeClientCredentials() (*dto.AccessTokenResponseDTO, error)
	grantTypeRefreshToken() (*dto.AccessTokenResponseDTO, error)
	generateCode() string
}

type authFlow struct {
	token *jwt.Token
	accessTokenRequest *dto.AccessTokenRequestDTO
	accessTokenExpiresAt int64
	refreshTokenExpiresAt int64
	authCodeExpiresAt int64
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

	accessTokenRequest := dto.AccessTokenRequestDTO{}
	if err := schema.NewDecoder().Decode(&accessTokenRequest, values); err != nil {
		return err
	}
	o.accessTokenRequest = &accessTokenRequest
	o.user = &models.User{
		Username: o.accessTokenRequest.Username,
		Password: &o.accessTokenRequest.Password,
	}


	auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	if len(auth) == 2 || auth[0] == "Basic" {
		payload, _ := base64.StdEncoding.DecodeString(auth[1])
		pair := strings.SplitN(string(payload), ":", 2)

		switch accessTokenRequest.GrantType {
			case enums.GrantTypeAuthorizationCode.String():
				o.accessTokenRequest.ClientID = pair[0]
				o.accessTokenRequest.ClientSecret = pair[1]
				break
			case enums.GrantTypeImplicit.String():
				o.accessTokenRequest.ClientID = pair[0]
				o.accessTokenRequest.ClientSecret = pair[1]
				break
			case enums.GrantTypePassword.String():
				o.accessTokenRequest.ClientID = pair[0]
				o.accessTokenRequest.ClientSecret = pair[1]
				o.user = &models.User{
					Username: accessTokenRequest.Username,
					Password: &accessTokenRequest.Password,
				}
				break
			default:
				o.user = &models.User{
					Username: pair[0],
					Password: &pair[1],
				}
				break
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

	if accessTokenRequest.GrantType == enums.GrantTypeClientCredentials.String() {
		if o.accessTokenRequest.ClientID == "" {
			o.accessTokenRequest.ClientID = o.user.Username
			o.accessTokenRequest.ClientSecret = *o.user.Password
		}
	}

	return nil
}

// ValidateParameters: validate parameters by authentication flow
func (o *authFlow) ValidateParameters() error {
	switch o.accessTokenRequest.GrantType {
		case enums.GrantTypeAuthorizationCode.String():
			if &o.accessTokenRequest.ClientID == nil {
				return errors.Errorf("client_id not defined")
			}

			if &o.accessTokenRequest.ClientSecret == nil {
				return errors.Errorf("client_secret not defined")
			}

			if &o.accessTokenRequest.Code == nil {
				return errors.Errorf("code not defined")
			}

			if &o.accessTokenRequest.RedirectURI == nil {
				return errors.Errorf("redirect_uri not defined")
			}
			break
		case enums.GrantTypeImplicit.String():
			if &o.accessTokenRequest.ClientID == nil {
				return errors.Errorf("client_id not defined")
			}

			if &o.accessTokenRequest.RedirectURI == nil {
				return errors.Errorf("redirect_uri not defined")
			}
			break
		case enums.GrantTypePassword.String():
			if &o.accessTokenRequest.ClientID == nil {
				return errors.Errorf("client_id not defined")
			}

			if &o.accessTokenRequest.ClientSecret == nil {
				return errors.Errorf("client_secret not defined")
			}

			if &o.accessTokenRequest.Username == nil {
				return errors.Errorf("username not defined")
			}

			if &o.accessTokenRequest.Password == nil {
				return errors.Errorf("password not defined")
			}
			break
		case enums.GrantTypeRefreshToken.String():
			if &o.accessTokenRequest.RefreshToken == nil {
				return errors.Errorf("refresh_token not defined")
			}
			if &o.accessToken == nil {
				return errors.Errorf("access_token not defined")
			}
			break
		case enums.GrantTypeClientCredentials.String():
			if &o.accessTokenRequest.ClientID == nil {
				return errors.Errorf("client_id not defined")
			}

			if &o.accessTokenRequest.ClientSecret == nil {
				return errors.Errorf("client_secret not defined")
			}
			break
	}
	return nil
}

// SetExpiresAt: set time expires
func (o *authFlow) SetExpiresAt(minutes int) {
	o.accessTokenExpiresAt = time.Now().Add(time.Minute * time.Duration(minutes)).Unix()
	o.refreshTokenExpiresAt = time.Now().Add(time.Minute * time.Duration(minutes * 120)).Unix()
	o.authCodeExpiresAt = time.Now().Add(time.Minute * time.Duration(5)).Unix()
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

func (o *authFlow) Login(loginDTO dto.LoginDTO) (*dto.LoginResponseDTO, error) {
	o.user = &models.User{
		Username:  loginDTO.Username,
		Password:  &loginDTO.Password,
	}

	o.accessTokenRequest = &dto.AccessTokenRequestDTO{
		ClientID:     loginDTO.ClientID,
	}

	loginResponseDTO := &dto.LoginResponseDTO{
		ResponseType: loginDTO.ResponseType,
		ClientID: loginDTO.ClientID,
		Username: o.user.Username,
		Scope: loginDTO.Scope,
		State: loginDTO.State,
		RedirectUri: &loginDTO.RedirectURI,
	}

	err := o.verifyCredentials()
	if err != nil {
		return nil, err
	}

	err = o.verifyClient()
	if err != nil {
		return nil, err
	}

	if o.client.Redirect != loginDTO.RedirectURI {
		return nil, errors.Errorf("invalid redirect_uri")
	}

	switch loginDTO.ResponseType {
		case enums.ResponseTypeToken.String():
			accessToken, err := o.getAccessTokenBase()
			if err != nil {
				return nil, errors.Errorf("failed to generate access_token: %s", err.Error())
			}


			err = o.accessTokenRepository.Create(accessToken)
			if err != nil {
				return nil, errors.Errorf("failed to generate access_token")
			}

			loginResponseDTO.AccessToken = &dto.AccessTokenResponseDTO{
				TokenType:    "Bearer",
				ExpiresIn:   o.accessTokenExpiresAt,
				AccessToken:  accessToken.AccessToken,
				State: &loginDTO.State,
			}

			err = o.getRefreshTokenBase(accessToken, loginResponseDTO.AccessToken)
			if err != nil {
				return nil, errors.Errorf("failed to generate access_token")
			}

			break
		case enums.ResponseTypeAuthorizationCode.String():
			loginResponseDTO.Code = o.generateCode()
			expiresAt := time.Unix(o.authCodeExpiresAt, 0)
			authCode := &models.AuthCode{
				Code:      loginResponseDTO.Code,
				Scopes:    o.accessTokenRequest.Scope,
				ExpiresAt: &expiresAt,
				User:      o.user,
				Client:    o.client,
			}

			err = o.authCodeRepository.Create(authCode)
			if err != nil {
				return nil, errors.Errorf("failed to generate authorization code")
			}
			break
		default:
			return nil, errors.Errorf("invalid response_type: %s", loginDTO.ResponseType)
	}

	return loginResponseDTO, nil
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

// verifyAuthCode: verify authorization code
func (o *authFlow) verifyAuthCode() error {
	authCode, err := o.authCodeRepository.FindByCode(o.accessTokenRequest.Code)
	if err != nil {
		return errors.Errorf("invalid authorization code")
	}

	if authCode.Revoked {
		return errors.Errorf("revoked authorization code")
	}

	o.client = authCode.Client
	o.user = authCode.User
	o.accessTokenRequest.Scope = strings.Replace(authCode.Scopes, " ", ",", 1000)

	authCode.Revoked = true
	err = o.authCodeRepository.Update(authCode)
	if err != nil {
		return errors.Errorf("failed to revoked authorization code")
	}

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

// GetAuthorizationCode: generate code for obtain access_token, flow `authorization_code`
func (o *authFlow) GetAuthorizationCode() (*string, error) {
	return nil, nil
}

// grantTypeAuthorizationCode: obtain access_token from de `code` of the flow `authorization_code`
func (o *authFlow) grantTypeAuthorizationCode() (*dto.AccessTokenResponseDTO, error) {

	err := o.verifyClientAndSecret()
	if err != nil {
		return nil, err
	}

	err = o.verifyAuthCode()
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

func (o *authFlow) generateCode() string {
	b := make([]byte, 64)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}