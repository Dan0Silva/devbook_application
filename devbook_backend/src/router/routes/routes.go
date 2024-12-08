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
	for _, route := range UserRoutes {
		router.HandleFunc(route.Uri, route.Function).Methods(route.Method)
	}
}
