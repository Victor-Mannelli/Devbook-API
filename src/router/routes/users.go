package routes

import (
	"api/src/controllers"
	"net/http"
)

var usersRoutes = []Route{
	{
		URI:      "/users",
		Method:   http.MethodPost,
		Function: controllers.CreateUser,
		Auth:     false,
	},
	{
		URI:      "/users",
		Method:   http.MethodGet,
		Function: controllers.FindUsers,
		Auth:     true,
	},
	{
		URI:      "/users/{userId}",
		Method:   http.MethodGet,
		Function: controllers.FindUserById,
		Auth:     true,
	},
	{
		URI:      "/users/changePassword",
		Method:   http.MethodPost,
		Function: controllers.ChangePassword,
		Auth:     true,
	},
	{
		URI:      "/users/{userId}",
		Method:   http.MethodPut,
		Function: controllers.UpdateUser,
		Auth:     true,
	},
	{
		URI:      "/users/{userId}",
		Method:   http.MethodDelete,
		Function: controllers.DeleteUser,
		Auth:     true,
	},
}
