package routes

import (
	"net/http"
	"oauth2/src/common"
	"oauth2/src/controllers"
)

type SwaggerRoutes interface {
	Routes() []*common.Route
}

type swaggerRoutesImpl struct {
	swaggerController controllers.SwaggerController
}

func NewSwaggerRoutes(swaggerController controllers.SwaggerController) *swaggerRoutesImpl {
	return &swaggerRoutesImpl{swaggerController}
}

func (h *swaggerRoutesImpl) Routes() []*common.Route {
	return []*common.Route{
		{
			http.MethodGet,
			"/swagger/docs.json",
			http.HandlerFunc(h.swaggerController.GetSwaggerJSON),
		},
	}
}