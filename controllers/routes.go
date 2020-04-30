package controllers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/larrygf02/go-blog/middlewares"
	"github.com/larrygf02/go-blog/views"
)

func (s *Server) InitializeRoutes() {
	// router := mux.NewRouter().StrictSlash(true)
	//s.Router.HandleFunc("/", home).Methods("GET")
	s.Router.HandleFunc("/user", views.NewUser).Methods("POST")
	s.Router.HandleFunc("/login", views.Login).Methods("POST")
	s.Router.HandleFunc("/storie", views.NewStorie).Methods("POST")
	s.Router.HandleFunc("/storiebyuser", views.StorieByUser).Methods("POST")
	s.Router.HandleFunc("/storievisit", views.NewStorieVisit).Methods("POST")
	s.Router.HandleFunc("/storievisit", views.GetAllStorieVisit).Methods("GET")
	// Habilitar CORS
	s.Router.Use(mux.CORSMethodMiddleware(s.Router))
	s.Router.Use(middlewares.SetMiddlewareJSON)
	log.Fatal(http.ListenAndServe(":5000", s.Router))
}
