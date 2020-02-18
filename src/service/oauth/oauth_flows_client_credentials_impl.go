package oauth

import (
	"net/http"
	"oauth2/src/models"
)

type authFlowClientCredentials struct {}

func (o * authFlowClientCredentials) GetToken(w http.ResponseWriter, r *http.Request) (*models.AccessToken, error) {
	return nil, nil
}

func (o * authFlowClientCredentials) GetAuthorizeCode(w http.ResponseWriter, r *http.Request) (*models.AuthCode, error) {
	return nil, nil
}
