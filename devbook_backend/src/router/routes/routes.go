package routes

import (
	"devbook_backend/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type Routes struct {
	Uri         string
	Method      string
	Function    func(http.ResponseWriter, *http.Request)
	RequireAuth bool
}

func GenerateRoutes(router *mux.Router) {
	allRoutes := []Routes{}

	allRoutes = append(allRoutes, UserRoutes...)
	allRoutes = append(allRoutes, LoginRoutes...)
	allRoutes = append(allRoutes, FollowerRoutes...)
	allRoutes = append(allRoutes, PostRoutes...)

	for _, route := range allRoutes {

		if route.RequireAuth {
			router.HandleFunc(route.Uri, middlewares.Authenticate(route.Function)).Methods(route.Method)
		} else {
			router.HandleFunc(route.Uri, route.Function).Methods(route.Method)
		}
	}
}
