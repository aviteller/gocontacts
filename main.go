package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/avi/go-contacts/app"
	"github.com/avi/go-contacts/controllers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	router := mux.NewRouter()

	c := cors.New(cors.Options{
		AllowedHeaders: []string{"*"},
		AllowedOrigins: []string{"*"},                                     // All origins
		AllowedMethods: []string{"GET", "HEAD", "POST", "PUT", "OPTIONS"}, // Allowing only get, just an example
	})

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")

	router.HandleFunc("/api/user/{id}/contacts", controllers.GetContactsFor).Methods("GET")
	router.HandleFunc("/api/contacts/new", controllers.CreateContact).Methods("POST")
	router.HandleFunc("/api/contacts/delete/{id}", controllers.DeleteContact).Methods("GET")

	router.HandleFunc("/api/me/contacts", controllers.GetContactsFor).Methods("GET")

	router.Use(app.JwtAuthentication)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	err := http.ListenAndServe(":"+port, c.Handler(router))
	if err != nil {
		fmt.Println(err)
	}

}
