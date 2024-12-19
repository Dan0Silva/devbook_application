package main

import (
	"devbook_backend/src/config"
	"devbook_backend/src/router"

	"fmt"
	"log"
	"net/http"
)

func main() {
	config.LoadEnvironment()

	// dbseed.PopulateDatabase(30, 20)

	port := config.Port

	router := router.Generate()

	fmt.Printf("\n  -> API running on port %s\n\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
