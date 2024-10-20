package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/guyjof/nethttp-crud-api/routes"
)

func main() {
	fmt.Println("REST API in Go 1.23.1!")

	routes.RegisterRoutes()

	fmt.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
