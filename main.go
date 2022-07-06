package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nicopv-dev/go-first-api/db"
	"github.com/nicopv-dev/go-first-api/routes"
)

func main() {
	router := mux.NewRouter()
	// database
	db.DBConnection()

	// migrate
	db.Migrate()

	// routes
	routes.SetRoutes(router)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("Server started on port http://localhost:8080")
	log.Println(server.ListenAndServe())
}