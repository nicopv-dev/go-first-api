package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func SetRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/api").Subrouter()
	subRoute.HandleFunc("/hello", func (writter http.ResponseWriter, request *http.Request)  {
		writter.Write([]byte("Hello World"))
	}).Methods("GET")

	// http:localhost:8080/api/users
	subRoute.HandleFunc("/users", GetAllUsers).Methods("GET")
	subRoute.HandleFunc("/users/{id}", GetUser).Methods("GET")
	subRoute.HandleFunc("/users", UpdateUser).Methods("POST")

	// http:localhost:8080/api/posts
}