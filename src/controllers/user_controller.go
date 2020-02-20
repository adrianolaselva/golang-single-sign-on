package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"oauth2/src/service"
)

type UserController interface {
	GetUser(w http.ResponseWriter, r *http.Request)
	GetUsers(w http.ResponseWriter, r *http.Request)
	PostUser(w http.ResponseWriter, r *http.Request)
	PutUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
}

type userControllerImpl struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *userControllerImpl {
	return &userControllerImpl{userService}
}

func (h *userControllerImpl) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	pathParams := mux.Vars(r)

	if pathParams["uuid"] == "" {
		response := make(map[string]interface{})
		response["message"] = "id user is not defined"
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(response)
		return
	}


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

	pathParams := mux.Vars(r)
	if pathParams["uuid"] == "" {
		response := make(map[string]interface{})
		response["message"] = "id user is not defined"
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(data)
}

func (h *userControllerImpl) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	pathParams := mux.Vars(r)
	if pathParams["uuid"] == "" {
		response := make(map[string]interface{})
		response["message"] = "id user is not defined"
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(data)
}