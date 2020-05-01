package controllers

import (
	"github.com/gorilla/mux"
	"github.com/larrygf02/go-blog/middlewares"
)

func (s *Server) InitializeRoutes() {
	s.Router.HandleFunc("/user", s.NewUser).Methods("POST")
	s.Router.HandleFunc("/login", s.Login).Methods("POST")
	s.Router.HandleFunc("/storie", s.NewStorie).Methods("POST")
	s.Router.HandleFunc("/storiebyuser", s.StorieByUser).Methods("POST")
	s.Router.HandleFunc("/storievisit", s.NewStorieVisit).Methods("POST")
	s.Router.HandleFunc("/storievisit", s.GetAllStorieVisit).Methods("GET")
	s.Router.HandleFunc("/storieapplause", s.SaveStorieApplause).Methods("POST")
	s.Router.HandleFunc("/storiecomment", s.SaveStorieComment).Methods("POST")
	// Habilitar CORS
	s.Router.Use(mux.CORSMethodMiddleware(s.Router))
	s.Router.Use(middlewares.SetMiddlewareJSON)
}
