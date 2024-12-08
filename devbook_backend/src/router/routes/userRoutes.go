package routes

import (
	userController "devbook_backend/src/controllers"
	"net/http"
)

var UserRoutes = []Routes{
	{
		Uri:         "/users", // create
		Method:      http.MethodPost,
		Function:    userController.CreateUser,
		RequireAuth: false,
	},
	{
		Uri:         "/users", // read all
		Method:      http.MethodGet,
		Function:    userController.GetAllUsers,
		RequireAuth: false,
	},
	{
		Uri:         "/users/{userId}", // read one user
		Method:      http.MethodGet,
		Function:    userController.GetUser,
		RequireAuth: false,
	},
	{
		Uri:         "/users/{userId}", // update
		Method:      http.MethodPut,
		Function:    userController.UpdateUser,
		RequireAuth: false,
	},
	{
		Uri:         "/users/{userId}", // delete
		Method:      http.MethodDelete,
		Function:    userController.DeleteUser,
		RequireAuth: false,
	},
}
