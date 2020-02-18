package oauth

import (
	"net/http"
	"oauth2/src/models"
)

type AuthFlow interface {
	GetToken(w http.ResponseWriter, r *http.Request) (*models.AccessToken, error)
	GetAuthorizeCode(w http.ResponseWriter, r *http.Request) (*models.AuthCode, error)
}
