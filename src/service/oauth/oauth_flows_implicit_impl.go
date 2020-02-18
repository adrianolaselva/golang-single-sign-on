package oauth

import (
	"net/http"
	"oauth2/src/models"
)

type authFlowImplicit struct {}

func (o * authFlowImplicit) GetToken(w http.ResponseWriter, r *http.Request) (*models.AccessToken, error) {
	return nil, nil
}

func (o * authFlowImplicit) GetAuthorizeCode(w http.ResponseWriter, r *http.Request) (*models.AuthCode, error) {
	return nil, nil
}
