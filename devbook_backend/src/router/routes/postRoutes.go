package routes

import (
	postController "devbook_backend/src/controllers"
	"net/http"
)

var PostRoutes = []Routes{
	{
		Uri:         "/posts",
		Method:      http.MethodPost,
		Function:    postController.CreatePost,
		RequireAuth: true,
	},
	{
		Uri:         "/posts/{userId}",
		Method:      http.MethodGet,
		Function:    postController.GetUserPosts,
		RequireAuth: false,
	},
}
