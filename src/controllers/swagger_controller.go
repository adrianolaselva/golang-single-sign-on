package controllers

import (
	"encoding/json"
	"github.com/go-openapi/swag"
	"log"
	"net/http"
)

type SwaggerController interface {
	GetSwaggerJSON(w http.ResponseWriter, r *http.Request)
}

type swaggerControllerImpl struct {

}

func NewSwaggerController() *swaggerControllerImpl {
	return &swaggerControllerImpl{}
}

func (h *swaggerControllerImpl) GetSwaggerJSON(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	swaggerJson, err := swag.YAMLDoc("doc/swagger.yaml")
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(swaggerJson)
}