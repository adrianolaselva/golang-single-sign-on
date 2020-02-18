package oauth

import (
	"net/http"
	"oauth2/src/models"
)

type authFlowPassword struct {}

func (o * authFlowPassword) GetToken(w http.ResponseWriter, r *http.Request) (*models.AccessToken, error) {
	return nil, nil
}

func (o * authFlowPassword) GetAuthorizeCode(w http.ResponseWriter, r *http.Request) (*models.AuthCode, error) {
	return nil, nil
}
