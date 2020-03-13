package oauth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/schema"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"net/url"
	"oauth2/src/dto"
	"oauth2/src/enums"
	"oauth2/src/repository"
	"time"
)

type AuthTokenClaim struct {
	*jwt.StandardClaims
	Data map[string]string
}

type AuthFlow interface {
	SetRequest(r *http.Request) error
	SetExpiresAt(minutes int)
	GetAccessToken() (*dto.AccessTokenResponseDTO, error)
	GetRefreshToken() (*dto.AccessTokenResponseDTO, error)
	GetAuthorizationCode() (*string, error)
}

type authFlow struct {
	token *jwt.Token
	accessTokenRequest *dto.AccessTokenRequestDTO
	expiresAt int64
	secret string
	clientRepository repository.ClientRepository
	authCodeRepository repository.AuthCodeRepository
	refreshTokenRepository repository.RefreshTokenRepository
	accessTokenRepository repository.AccessTokenRepository
}

func NewAuthFlow(
		hmac *jwt.SigningMethodHMAC,
		secret string,
		clientRepository repository.ClientRepository,
		authCodeRepository repository.AuthCodeRepository,
		refreshTokenRepository repository.RefreshTokenRepository,
		accessTokenRepository repository.AccessTokenRepository) *authFlow {
	return &authFlow{
		token: jwt.New(hmac),
		secret: secret,
		clientRepository: clientRepository,
		authCodeRepository: authCodeRepository,
		refreshTokenRepository: refreshTokenRepository,
		accessTokenRepository: accessTokenRepository,
	}
}

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

	return nil
}

func (o *authFlow) SetExpiresAt(minutes int) {
	o.expiresAt = time.Now().Add(time.Minute * time.Duration(minutes)).Unix()
}

func (o *authFlow) GetAccessToken() (*dto.AccessTokenResponseDTO, error) {
	switch o.accessTokenRequest.GrantType {
		case enums.GrantTypePassword.String():
			return o.grantTypePassword()
		case enums.GrantTypeRefreshToken.String():
			return o.grantTypeRefreshToken()
		case enums.GrantTypeAuthorizationCode.String():
			return o.grantTypeAuthorizationCode()
		case enums.GrantTypeClientCredentials.String():
			return o.grantTypeClientCredentials()
		case enums.GrantTypeImplicit.String():
			return o.grantTypeImplicit()
		default:
			return nil, errors.Errorf("invalid grant_type %s", o.accessTokenRequest.GrantType)
	}
}

func (o *authFlow) GetRefreshToken() (*dto.AccessTokenResponseDTO, error) {
	return nil, nil
}

func (o *authFlow) GetAuthorizationCode() (*string, error) {
	return nil, nil
}

func (o *authFlow) grantTypeAuthorizationCode() (*dto.AccessTokenResponseDTO, error) {
	data := make(map[string]string)
	o.token.Claims = &AuthTokenClaim{
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: o.expiresAt,
		},
		Data: data,
	}

	accessToken, err := o.token.SignedString([]byte(o.secret))
	if err != nil {
		return nil, errors.Errorf("failed to generate token %s", err.Error())
	}

	return &dto.AccessTokenResponseDTO{
		TokenType:   "Bearer",
		ExpiresIn:   o.expiresAt,
		AccessToken: accessToken,
	}, nil
}

func (o *authFlow) grantTypeClientCredentials() (*dto.AccessTokenResponseDTO, error) {
	data := make(map[string]string)
	o.token.Claims = &AuthTokenClaim{
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: o.expiresAt,
		},
		Data: data,
	}

	accessToken, err := o.token.SignedString([]byte(o.secret))
	if err != nil {
		return nil, errors.Errorf("failed to generate token %s", err.Error())
	}

	return &dto.AccessTokenResponseDTO{
		TokenType:   "Bearer",
		ExpiresIn:   o.expiresAt,
		AccessToken: accessToken,
	}, nil
}

func (o *authFlow) grantTypePassword() (*dto.AccessTokenResponseDTO, error) {

	// verify user and password

	// verify client_id verify client_id by user_id

	// verify scopes

	// verify if not revoked

	// generate access_token

	data := make(map[string]string)
	o.token.Claims = &AuthTokenClaim{
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: o.expiresAt,
		},
		Data: data,
	}

	accessToken, err := o.token.SignedString([]byte(o.secret))
	if err != nil {
		return nil, errors.Errorf("failed to generate token %s", err.Error())
	}

	return &dto.AccessTokenResponseDTO{
		TokenType:   "Bearer",
		ExpiresIn:   o.expiresAt,
		RefreshToken: &accessToken,
		AccessToken: accessToken,
	}, nil
}

func (o *authFlow) grantTypeRefreshToken() (*dto.AccessTokenResponseDTO, error) {
	data := make(map[string]string)
	o.token.Claims = &AuthTokenClaim{
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: o.expiresAt,
		},
		Data: data,
	}

	accessToken, err := o.token.SignedString([]byte(o.secret))
	if err != nil {
		return nil, errors.Errorf("failed to generate token %s", err.Error())
	}

	return &dto.AccessTokenResponseDTO{
		TokenType:   "Bearer",
		ExpiresIn:   o.expiresAt,
		RefreshToken: &accessToken,
		AccessToken: accessToken,
	}, nil
}

func (o *authFlow) grantTypeImplicit() (*dto.AccessTokenResponseDTO, error) {
	data := make(map[string]string)
	o.token.Claims = &AuthTokenClaim{
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: o.expiresAt,
		},
		Data: data,
	}

	accessToken, err := o.token.SignedString([]byte(o.secret))
	if err != nil {
		return nil, errors.Errorf("failed to generate token %s", err.Error())
	}

	return &dto.AccessTokenResponseDTO{
		TokenType:   "Bearer",
		ExpiresIn:   o.expiresAt,
		RefreshToken: &accessToken,
		AccessToken: accessToken,
	}, nil
}
