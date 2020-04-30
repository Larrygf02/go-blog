package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/larrygf02/go-blog/config"
	"github.com/larrygf02/go-blog/middlewares"
	"github.com/larrygf02/go-blog/views"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Back")
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", home).Methods("GET")
	router.HandleFunc("/user", views.NewUser).Methods("POST")
	router.HandleFunc("/login", views.Login).Methods("POST")
	router.HandleFunc("/storie", views.NewStorie).Methods("POST")
	router.HandleFunc("/storiebyuser", views.StorieByUser).Methods("POST")
	router.HandleFunc("/storievisit", views.NewStorieVisit).Methods("POST")
	router.HandleFunc("/storievisit", views.GetAllStorieVisit).Methods("GET")
	// Habilitar CORS
	router.Use(mux.CORSMethodMiddleware(router))
	router.Use(middlewares.SetMiddlewareJSON)
	log.Fatal(http.ListenAndServe(":5000", router))
}

func main() {
	fmt.Println("Hello world")
	config.InitialMigration()
	handleRequests()
}
