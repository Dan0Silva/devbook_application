package main

import (
	"devbook_backend/src/router"

	"fmt"
	"log"
	"net/http"
)

func main() {
	port := "5000"

	router := router.Generate()

	fmt.Printf("API rodando na porta %s\n\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
