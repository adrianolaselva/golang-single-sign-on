package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"oauth2/src/common"
	"oauth2/src/dto"
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
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(&dto.DefaultResponseDto{
			Message: "Not found",
		})
		return
	}

	user, err := h.userService.FindById(pathParams["uuid"])
	if  err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(&dto.DefaultResponseDto{
			Message: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(user)
}

func (h *userControllerImpl) GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	httpRequestCommon := common.NewHTTPRequestCommon(r)
	pagingParameters, _ := httpRequestCommon.GetPaginateParameters()

	data, err := h.userService.Paginate(
		&pagingParameters.Filters,
		&pagingParameters.OrderBy,
		&pagingParameters.OrderDir,
		&pagingParameters.Limit,
		&pagingParameters.Page)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(&dto.DefaultResponseDto{
			Message: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(data)
}

func (h *userControllerImpl) PostUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var userDto dto.UserDto
	err := json.NewDecoder(r.Body).Decode(&userDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(&dto.DefaultResponseDto{
			Message: err.Error(),
		})
		return
	}

	user, err := userDto.ToUser()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(&dto.DefaultResponseDto{
			Message: err.Error(),
		})
		return
	}

	err = h.userService.Create(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(&dto.DefaultResponseDto{
			Message: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(user)
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

	var userDto dto.UserDto
	err := json.NewDecoder(r.Body).Decode(&userDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(&dto.DefaultResponseDto{
			Message: err.Error(),
		})
		return
	}

	user, err := userDto.ToUser()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(&dto.DefaultResponseDto{
			Message: err.Error(),
		})
		return
	}

	user.ID = pathParams["uuid"]
	err = h.userService.Update(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(&dto.DefaultResponseDto{
			Message: err.Error(),
		})
		return
	}

	user, _ = h.userService.FindById(user.ID)

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(user)
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

	user, err := h.userService.FindById(pathParams["uuid"])
	if user == nil {
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(&dto.DefaultResponseDto{
			"user not found",
		})
	}

	err = h.userService.Delete(pathParams["uuid"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(&dto.DefaultResponseDto{
			Message:err.Error(),
		})
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(&dto.DefaultResponseDto{})
}