package routes

import (
	followersController "devbook_backend/src/controllers"
	"net/http"
)

var FollowersRoutes = []Routes{
	{
		Uri:         "/follow/{userId}",
		Method:      http.MethodPost,
		Function:    followersController.FollowUser,
		RequireAuth: true,
	},
	{
		Uri:         "/follow/{userId}",
		Method:      http.MethodDelete,
		Function:    followersController.UnfollowUser,
		RequireAuth: true,
	},
	{
		Uri:         "/following/{userId}",
		Method:      http.MethodGet,
		Function:    followersController.GetUserFollowing,
		RequireAuth: true,
	},
	{
		Uri:         "/followers/{userId}",
		Method:      http.MethodGet,
		Function:    followersController.GetUserFollowers,
		RequireAuth: true,
	},
}
