package controllers

import (
	"github.com/gorilla/mux"
	"github.com/larrygf02/go-blog/middlewares"
)

func (s *Server) InitializeRoutes() {
	s.Router.HandleFunc("/user", s.NewUser).Methods("POST")
	s.Router.HandleFunc("/login", s.Login).Methods("POST")
	s.Router.HandleFunc("/storie", s.NewStorie).Methods("POST")
	s.Router.HandleFunc("/storie/user/{id}", s.StorieByUser).Methods("GET")
	s.Router.HandleFunc("/storievisit", s.NewStorieVisit).Methods("POST")
	s.Router.HandleFunc("/storievisit", s.GetAllStorieVisit).Methods("GET")
	s.Router.HandleFunc("/storieapplause", s.SaveStorieApplause).Methods("POST")
	s.Router.HandleFunc("/storiecomment", s.SaveStorieComment).Methods("POST")
	s.Router.HandleFunc("/storiecomment/{id}", s.UpdateStorieComment).Methods("PUT")
	s.Router.HandleFunc("/draft", s.NewDraft).Methods("POST")
	s.Router.HandleFunc("/draft/{id}", s.UpdateDraft).Methods("PUT")
	s.Router.HandleFunc("/draft/user/{user_id}", s.DraftByUser).Methods("GET")
	s.Router.HandleFunc("/user/storie/favorites", s.StoriesFavorites).Methods("POST")
	// Habilitar CORS
	s.Router.Use(mux.CORSMethodMiddleware(s.Router))
	s.Router.Use(middlewares.SetMiddlewareJSON)
}
