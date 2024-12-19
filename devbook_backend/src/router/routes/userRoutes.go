package routes

import (
	userController "devbook_backend/src/controllers"
	"net/http"
)

var UserRoutes = []Routes{
	{
		Uri:         "/users",
		Method:      http.MethodPost,
		Function:    userController.CreateUser,
		RequireAuth: false,
	},
	{
		Uri:         "/users",
		Method:      http.MethodGet,
		Function:    userController.SearchUsers,
		RequireAuth: false,
	},
	{
		Uri:         "/users/{userId}",
		Method:      http.MethodGet,
		Function:    userController.GetUserById,
		RequireAuth: false,
	},
	{
		Uri:         "/users/{userId}",
		Method:      http.MethodPut,
		Function:    userController.UpdateUser,
		RequireAuth: true,
	},
	{
		Uri:         "/users/{userId}",
		Method:      http.MethodDelete,
		Function:    userController.DeleteUser,
		RequireAuth: true,
	},
	{
		Uri:         "/users/update-password",
		Method:      http.MethodPatch,
		Function:    userController.UpdatePassword,
		RequireAuth: true,
	},
}
