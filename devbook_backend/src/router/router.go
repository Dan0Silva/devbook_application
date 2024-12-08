package router

import (
	"devbook_backend/src/router/routes"

	"github.com/gorilla/mux"
)

func Generate() *mux.Router {
	r := mux.NewRouter()
	routes.GenerateRoutes(r)

	return r
}
