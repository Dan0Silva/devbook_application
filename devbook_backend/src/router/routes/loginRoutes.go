package routes

import (
	loginController "devbook_backend/src/controllers"
	"net/http"
)

var LoginRoutes = []Routes{
	{
		Uri:         "/login", // login
		Method:      http.MethodPost,
		Function:    loginController.Login,
		RequireAuth: false,
	},
}
