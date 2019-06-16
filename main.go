package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/avi/go-contacts/app"
	"github.com/avi/go-contacts/controllers"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")

	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/user/{id}/contacts", controllers.GetContactsFor).Methods("GET")
	router.HandleFunc("/api/contacts/new", controllers.CreateContact).Methods("POST")

	router.HandleFunc("/api/me/contacts", controllers.GetContactsFor).Methods("GET")

	router.Use(app.JwtAuthentication)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Println(err)
	}

}
