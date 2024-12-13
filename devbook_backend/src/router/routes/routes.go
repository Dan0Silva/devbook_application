package routes

import (
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

	for _, route := range allRoutes {
		router.HandleFunc(route.Uri, route.Function).Methods(route.Method)
	}
}
