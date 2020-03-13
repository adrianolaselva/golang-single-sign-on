package controllers

import (
	"encoding/json"
	_ "github.com/dgrijalva/jwt-go"
	"net/http"
	"oauth2/src/dto"
	_ "oauth2/src/dto"
	"oauth2/src/service/oauth"
)

type OAuthController interface {
	GetAuthorize(w http.ResponseWriter, r *http.Request)
	GetLogin(w http.ResponseWriter, r *http.Request)
	PostToken(w http.ResponseWriter, r *http.Request)
	GetUserInfo(w http.ResponseWriter, r *http.Request)
}

type oAuthControllerImpl struct {
	authFlow oauth.AuthFlow
}



func NewOAuthController(authFlow oauth.AuthFlow) *oAuthControllerImpl {
	return &oAuthControllerImpl{authFlow: authFlow}
}

func (h *oAuthControllerImpl) GetAuthorize(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	data := make(map[string]string)
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(data)
}

func (h *oAuthControllerImpl) GetHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	data := make(map[string]string)
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(data)
}

func (h *oAuthControllerImpl) GetLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	data := make(map[string]string)
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(data)
}

func (h *oAuthControllerImpl) PostToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	err := h.authFlow.SetRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(&dto.AccessTokenErrorResponseDTO{
			Error:	err.Error(),
		})
		return
	}

	h.authFlow.SetExpiresAt(120)

	accessToken, err := h.authFlow.GetAccessToken()
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode(&dto.AccessTokenErrorResponseDTO{
			Error:	err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(accessToken)
}

func (h *oAuthControllerImpl) GetUserInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	data := make(map[string]string)
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(data)
}