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
		Uri:         "/posts",
		Method:      http.MethodGet,
		Function:    postController.GetAllPosts,
		RequireAuth: false,
	},
	{
		Uri:         "/posts/{userId}",
		Method:      http.MethodGet,
		Function:    postController.GetUserPosts,
		RequireAuth: false,
	},
	{
		Uri:         "/posts/{postId}",
		Method:      http.MethodPatch,
		Function:    postController.UpdatePost,
		RequireAuth: true,
	},
	{
		Uri:         "/posts/{postId}",
		Method:      http.MethodDelete,
		Function:    postController.DeletePost,
		RequireAuth: true,
	},
}
