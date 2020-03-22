package controllers

import (
	"encoding/json"
	"fmt"
	_ "github.com/dgrijalva/jwt-go"
	"net/http"
	"net/url"
	"oauth2/src/dto"
	_ "oauth2/src/dto"
	"oauth2/src/service/oauth"
	"strings"
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
	values := make(url.Values)
	for key := range r.URL.Query() {
		values.Set(key, r.URL.Query().Get(key))
	}

	// validar parâmetros

	parameters := ""
	for key, value := range values {
		if len(parameters) > 0 {
			parameters += "&"
		}
		parameters += fmt.Sprintf("%s=%s", key, strings.Join(value, ","))
	}

	http.Redirect(w, r, fmt.Sprintf("/app/#/auth/login?%s", parameters), http.StatusTemporaryRedirect)
}

func (h *oAuthControllerImpl) GetHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	data := make(map[string]string)
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(data)
}

func (h *oAuthControllerImpl) GetLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var loginDTO dto.LoginDTO
	err := json.NewDecoder(r.Body).Decode(&loginDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(&dto.DefaultResponseDto{
			Message: err.Error(),
		})
		return
	}

	err = h.authFlow.ValidateParameters()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(&dto.AccessTokenErrorResponseDTO{
			Error:	err.Error(),
		})
		return
	}

	h.authFlow.SetExpiresAt(120)

	loginResponse, err := h.authFlow.Login(loginDTO)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode(&dto.AccessTokenErrorResponseDTO{
			Status: -2,
			Error:	err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(loginResponse)
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

	err = h.authFlow.ValidateParameters()
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

	err := h.authFlow.ValidateParameters()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(&dto.AccessTokenErrorResponseDTO{
			Error:	err.Error(),
		})
		return
	}

	data := make(map[string]string)
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(data)
}