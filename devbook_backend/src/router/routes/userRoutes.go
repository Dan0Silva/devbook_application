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
		Uri:         "/follow/{userId}",
		Method:      http.MethodPost,
		Function:    userController.FollowUser,
		RequireAuth: true,
	},
	{
		Uri:         "/follow/{userId}",
		Method:      http.MethodDelete,
		Function:    userController.UnfollowUser,
		RequireAuth: true,
	},
	{
		Uri:         "/following/{userId}",
		Method:      http.MethodGet,
		Function:    userController.GetUserFollowing,
		RequireAuth: true,
	},
	{
		Uri:         "/followers/{userId}",
		Method:      http.MethodGet,
		Function:    userController.GetUserFollowers,
		RequireAuth: true,
	},
}
