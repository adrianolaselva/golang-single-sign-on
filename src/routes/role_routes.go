package routes

import (
	"net/http"
	"oauth2/src/common"
	"oauth2/src/controllers"
	"oauth2/src/middlewares"
)

type Role interface {
	Routes() []*common.Route
}

type roleImpl struct {
	roleController controllers.RoleController
	authenticationMiddleware middlewares.AuthenticationMiddleware
}

func NewRoleRoutes(
	roleController controllers.RoleController,
	authenticationMiddleware middlewares.AuthenticationMiddleware) *roleImpl {
	return &roleImpl{roleController, authenticationMiddleware}
}

func (o *roleImpl) Routes() []*common.Route {
	return []*common.Route{
		{
			http.MethodGet,
			"/v1/roles/{uuid}",
			o.authenticationMiddleware.
				ValidateToken(http.HandlerFunc(o.roleController.GetRole), []string{"roles:read"}),
		},
		{
			http.MethodGet,
			"/v1/roles",
			o.authenticationMiddleware.
				ValidateToken(http.HandlerFunc(o.roleController.GetRoles), []string{"roles:read"}),
		},
		{
			http.MethodPost,
			"/v1/roles",
			o.authenticationMiddleware.
				ValidateToken(http.HandlerFunc(o.roleController.PostRole), []string{"roles:write"}),
		},
		{
			http.MethodPut,
			"/v1/roles/{uuid}",
			o.authenticationMiddleware.
				ValidateToken(http.HandlerFunc(o.roleController.PutRole), []string{"roles:write"}),
		},
		{
			http.MethodDelete,
			"/v1/roles/{uuid}",
			o.authenticationMiddleware.
				ValidateToken(http.HandlerFunc(o.roleController.DeleteRole), []string{"roles:delete"}),
		},
	}
}