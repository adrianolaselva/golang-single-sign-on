package oauth

import (
	"net/http"
	"oauth2/src/models"
)

type authFlowRefreshToken struct {}

func (o * authFlowRefreshToken) GetToken(w http.ResponseWriter, r *http.Request) (*models.AccessToken, error) {
	return nil, nil
}

func (o * authFlowRefreshToken) GetAuthorizeCode(w http.ResponseWriter, r *http.Request) (*models.AuthCode, error) {
	return nil, nil
}
