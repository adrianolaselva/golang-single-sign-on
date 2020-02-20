package routes

import (
	"net/http"
	"oauth2/src/common"
	"oauth2/src/controllers"
	"oauth2/src/middlewares"
)

type User interface {
	Routes() []*common.Route
}

type userImpl struct {
	userController controllers.UserController
	authenticationMiddleware middlewares.AuthenticationMiddleware
}

func NewUserRoutes(
	userController controllers.UserController,
	authenticationMiddleware middlewares.AuthenticationMiddleware) *userImpl {
	return &userImpl{userController, authenticationMiddleware}
}

func (o *userImpl) Routes() []*common.Route {
	return []*common.Route{
		{
			http.MethodGet,
			"/v1/users/{uuid}",
			o.authenticationMiddleware.
				ValidateJWTToken(http.HandlerFunc(o.userController.GetUser), []string{"users:read"}),
		},
		{
			http.MethodGet,
			"/v1/users",
			o.authenticationMiddleware.
				ValidateJWTToken(http.HandlerFunc(o.userController.GetUsers), []string{"users:read"}),
		},
		{
			http.MethodPost,
			"/v1/users",
			o.authenticationMiddleware.
				ValidateJWTToken(http.HandlerFunc(o.userController.PostUser), []string{"users:write"}),
		},
		{
			http.MethodPut,
			"/v1/users/{uuid}",
			o.authenticationMiddleware.
				ValidateJWTToken(http.HandlerFunc(o.userController.PutUser), []string{"users:write"}),
		},
		{
			http.MethodDelete,
			"/v1/users/{uuid}",
			o.authenticationMiddleware.
				ValidateJWTToken(http.HandlerFunc(o.userController.DeleteUser), []string{"users:delete"}),
		},
	}
}