package routes

import (
	"api/src/controllers"
	"net/http"
)

var routesUsers = []Route{
	//* Cadastra um usuário
	{
		URI:         "/users",
		Method:      http.MethodPost,
		Function:    controllers.CreateUser,
		RequireAuth: false,
	},

	//* Busca todos os usuários
	{
		URI:         "/users",
		Method:      http.MethodGet,
		Function:    controllers.GetUsers,
		RequireAuth: false,
	},
	//* Busca um único usuário
	{
		URI:         "/users/{userId}",
		Method:      http.MethodGet,
		Function:    controllers.GetUserById,
		RequireAuth: false,
	},
	//* Atualiza um usuário
	{
		URI:         "/users/{userId}",
		Method:      http.MethodPut,
		Function:    controllers.UpdateUser,
		RequireAuth: false,
	},
	//* Deleta um usuário
	{
		URI:         "/users/{userId}",
		Method:      http.MethodDelete,
		Function:    controllers.DeleteUser,
		RequireAuth: false,
	},
}
