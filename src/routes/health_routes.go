package routes

import (
	"net/http"
	"oauth2/src/common"
	"oauth2/src/controllers"
)

type HealthRoutes interface {
	Routes() []*common.Route
}

type healthRoutesImpl struct {
	healthController controllers.HealthController
}

func NewHealthRoutes(healthController controllers.HealthController) *healthRoutesImpl {
	return &healthRoutesImpl{healthController}
}

func (h *healthRoutesImpl) Routes() []*common.Route {
	return []*common.Route{
		{
			http.MethodGet,
			"/health",
			http.HandlerFunc(h.healthController.GetHealth),
		},
	}
}