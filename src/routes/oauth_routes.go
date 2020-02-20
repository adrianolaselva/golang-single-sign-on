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
	authenticationMiddleware middlewares.AuthenticationMiddleware
}

func NewOAuthRoutes(
	oAuthController controllers.OAuthController,
	authenticationMiddleware middlewares.AuthenticationMiddleware) *oAuthImpl {
	return &oAuthImpl{oAuthController, authenticationMiddleware}
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
			o.authenticationMiddleware.
				ValidateClientIdAndSecret(http.HandlerFunc(o.oAuthController.PostToken)),
		},
		{
			http.MethodGet,
			"/oauth2/user-info",
			o.authenticationMiddleware.
				ValidateJWTToken(http.HandlerFunc(o.oAuthController.PostToken), []string{}),
		},
	}
}