package controllers


import (
	"encoding/json"
	"net/http"
)

type OAuthController interface {
	GetAuthorize(w http.ResponseWriter, r *http.Request)
	GetLogin(w http.ResponseWriter, r *http.Request)
	PostToken(w http.ResponseWriter, r *http.Request)
	GetUserInfo(w http.ResponseWriter, r *http.Request)
}

type oAuthControllerImpl struct {

}



func NewOAuthController() *oAuthControllerImpl {
	return &oAuthControllerImpl{}
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
	data := make(map[string]string)
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(data)
}

func (h *oAuthControllerImpl) GetUserInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	data := make(map[string]string)
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(data)
}