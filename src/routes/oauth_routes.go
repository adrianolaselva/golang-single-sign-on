package routes

import (
	"net/http"
	"oauth2/src/common"
	"oauth2/src/controllers"
	"oauth2/src/middlewares"
)

type OAuth interface {
	Routes() []*common.Route
}

type oAuthImpl struct {
	oAuthController controllers.OAuthController
}

func NewOAuthRoutes(oAuthController controllers.OAuthController) *oAuthImpl {
	return &oAuthImpl{oAuthController}
}

func (o *oAuthImpl) Routes() []*common.Route {
	return []*common.Route{
		{
			http.MethodGet,
			"/oauth2/authorize",
			http.HandlerFunc(o.oAuthController.GetAuthorize),
		},
		{
			http.MethodGet,
			"/oauth2/login",
			http.HandlerFunc(o.oAuthController.GetLogin),
		},
		{
			http.MethodPost,
			"/oauth2/token",
			middlewares.
				NewAuthenticationMiddleware().
				ValidateClientIdAndSecret(http.HandlerFunc(o.oAuthController.PostToken)),
		},
		{
			http.MethodGet,
			"/oauth2/user-info",
			middlewares.
				NewAuthenticationMiddleware().
				ValidateJWTToken(http.HandlerFunc(o.oAuthController.GetUserInfo)),
		},
	}
}