package routes

import (
	"api/src/controllers"
	"net/http"
)

var routerUser = []Route{
	{
		URI:         "/usuarios",
		Method:      http.MethodPost,
		Function:    controllers.CreateUser,
		RequireAuth: false,
	},
	{
		URI:         "/usuarios",
		Method:      http.MethodGet,
		Function:    controllers.SelectUsers,
		RequireAuth: false,
	},
	{
		URI:         "/usuarios/{usuarioId}",
		Method:      http.MethodGet,
		Function:    controllers.SelectUser,
		RequireAuth: false,
	},
	{
		URI:         "/usuarios/{usuarioId}",
		Method:      http.MethodPut,
		Function:    controllers.UpdateUser,
		RequireAuth: false,
	},
	{
		URI:         "/usuarios/{usuarioId}",
		Method:      http.MethodDelete,
		Function:    controllers.DeleteUser,
		RequireAuth: false,
	},
}
