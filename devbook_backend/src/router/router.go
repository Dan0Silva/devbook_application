package router

import (
	"devbook_backend/src/router/routes"
	"fmt"

	"github.com/gorilla/mux"
)

func Generate() *mux.Router {
	r := mux.NewRouter()
	routes.GenerateRoutes(r)

	fmt.Printf("  Generating router successfuly\n")
	return r
}
