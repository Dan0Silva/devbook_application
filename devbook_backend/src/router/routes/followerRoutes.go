package routes

import (
	followerController "devbook_backend/src/controllers"
	"net/http"
)

var FollowerRoutes = []Routes{
	{
		Uri:         "/follow/{userId}",
		Method:      http.MethodPost,
		Function:    followerController.FollowUser,
		RequireAuth: true,
	},
	{
		Uri:         "/follow/{userId}",
		Method:      http.MethodDelete,
		Function:    followerController.UnfollowUser,
		RequireAuth: true,
	},
	{
		Uri:         "/following/{userId}",
		Method:      http.MethodGet,
		Function:    followerController.GetUserFollowing,
		RequireAuth: true,
	},
	{
		Uri:         "/followers/{userId}",
		Method:      http.MethodGet,
		Function:    followerController.GetUserFollowers,
		RequireAuth: true,
	},
}
