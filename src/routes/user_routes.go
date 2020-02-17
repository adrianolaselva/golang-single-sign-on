package routes

import (
	"net/http"
	"oauth2/src/common"
	"oauth2/src/controllers"
)

type User interface {
	Routes() []*common.Route
}

type userImpl struct {
	userController controllers.UserController
}

func NewUserRoutes(userController controllers.UserController) *userImpl {
	return &userImpl{userController}
}

func (o *userImpl) Routes() []*common.Route {
	return []*common.Route{
		{
			http.MethodGet,
			"/v1/users/{uuid}",
			http.HandlerFunc(o.userController.GetUser),
		},
		{
			http.MethodGet,
			"/v1/users",
			http.HandlerFunc(o.userController.GetUsers),
		},
		{
			http.MethodPost,
			"/v1/users",
			http.HandlerFunc(o.userController.PostUser),
		},
		{
			http.MethodPut,
			"/v1/users/{uuid}",
			http.HandlerFunc(o.userController.PutUser),
		},
		{
			http.MethodDelete,
			"/v1/users/{uuid}",
			http.HandlerFunc(o.userController.DeleteUser),
		},
	}
}