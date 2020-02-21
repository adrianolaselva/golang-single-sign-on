package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"oauth2/src/common"
	"oauth2/src/dto"
	"oauth2/src/service"
)

type RoleController interface {
	GetRole(w http.ResponseWriter, r *http.Request)
	GetRoles(w http.ResponseWriter, r *http.Request)
	PostRole(w http.ResponseWriter, r *http.Request)
	PutRole(w http.ResponseWriter, r *http.Request)
	DeleteRole(w http.ResponseWriter, r *http.Request)
}

type roleControllerImpl struct {
	roleService service.RoleService
}

func NewRoleController(roleService service.RoleService) *roleControllerImpl {
	return &roleControllerImpl{roleService}
}

func (h *roleControllerImpl) GetRole(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	pathParams := mux.Vars(r)

	if pathParams["uuid"] == "" {
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(&dto.DefaultResponseDto{
			Message: "Not found",
		})
		return
	}

	role, err := h.roleService.FindById(pathParams["uuid"])
	if  err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(&dto.DefaultResponseDto{
			Message: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(role)
}

func (h *roleControllerImpl) GetRoles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	httpRequestCommon := common.NewHTTPRequestCommon(r)
	pagingParameters, _ := httpRequestCommon.GetPaginateParameters()

	data, err := h.roleService.Paginate(
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

func (h *roleControllerImpl) PostRole(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var roleDto dto.RoleDto
	err := json.NewDecoder(r.Body).Decode(&roleDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(&dto.DefaultResponseDto{
			Message: err.Error(),
		})
		return
	}

	role := roleDto.ToRole()
	err = h.roleService.Create(role)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(&dto.DefaultResponseDto{
			Message: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(role)
}

func (h *roleControllerImpl) PutRole(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	pathParams := mux.Vars(r)
	if pathParams["uuid"] == "" {
		response := make(map[string]interface{})
		response["message"] = "id role is not defined"
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(response)
		return
	}

	var roleDto dto.RoleDto
	err := json.NewDecoder(r.Body).Decode(&roleDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(&dto.DefaultResponseDto{
			Message: err.Error(),
		})
		return
	}

	role := roleDto.ToRole()
	role.ID = pathParams["uuid"]
	err = h.roleService.Update(role)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(&dto.DefaultResponseDto{
			Message: err.Error(),
		})
		return
	}

	role, _ = h.roleService.FindById(role.ID)

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(role)
}

func (h *roleControllerImpl) DeleteRole(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	pathParams := mux.Vars(r)
	if pathParams["uuid"] == "" {
		response := make(map[string]interface{})
		response["message"] = "id role is not defined"
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(response)
		return
	}

	err := h.roleService.Delete(pathParams["uuid"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(&dto.DefaultResponseDto{
			Message:err.Error(),
		})
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(&dto.DefaultResponseDto{})
}