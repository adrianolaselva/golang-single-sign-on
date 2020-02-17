package controllers

import (
	"encoding/json"
	"net/http"
	"oauth2/src/repository"
)

type UserController interface {
	GetUser(w http.ResponseWriter, r *http.Request)
	GetUsers(w http.ResponseWriter, r *http.Request)
	PostUser(w http.ResponseWriter, r *http.Request)
	PutUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
}

type userControllerImpl struct {
	userRepository repository.UserRepository
}

func NewUserController(userRepository repository.UserRepository) *userControllerImpl {
	return &userControllerImpl{userRepository}
}

func (h *userControllerImpl) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	data := make(map[string]string)
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(data)
}

func (h *userControllerImpl) GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	data := make(map[string]string)
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(data)
}

func (h *userControllerImpl) PostUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	data := make(map[string]string)
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(data)
}

func (h *userControllerImpl) PutUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	data := make(map[string]string)
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(data)
}

func (h *userControllerImpl) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	data := make(map[string]string)
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(data)
}