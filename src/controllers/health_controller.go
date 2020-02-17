package controllers

import (
	"encoding/json"
	"net/http"
)

type HealthController interface {
	GetHealth(w http.ResponseWriter, r *http.Request)
}

type healthControllerImpl struct {

}

func NewHealthController() *healthControllerImpl {
	return &healthControllerImpl{}
}

func (h *healthControllerImpl) GetHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	data := make(map[string]string)
	data["message"] = "health check ok"
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(data)
}