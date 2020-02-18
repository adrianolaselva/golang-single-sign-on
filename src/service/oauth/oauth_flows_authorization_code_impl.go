package oauth

import (
	"net/http"
	"oauth2/src/models"
)

type authFlowAuthorizationCode struct {}

func (o * authFlowAuthorizationCode) GetToken(w http.ResponseWriter, r *http.Request) (*models.AccessToken, error) {
	return nil, nil
}

func (o * authFlowAuthorizationCode) GetAuthorizeCode(w http.ResponseWriter, r *http.Request) (*models.AuthCode, error) {
	return nil, nil
}