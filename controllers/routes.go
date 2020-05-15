package controllers

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/larrygf02/go-blog/middlewares"
)

func (s *Server) InitializeRoutes() {
	s.Router.HandleFunc("/user", s.NewUser).Methods("POST")
	s.Router.HandleFunc("/login", s.Login).Methods("POST", "OPTIONS")
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
	s.Router.HandleFunc("/user/storie/favorites/{id}", s.GetStoriesFavorites).Methods("GET")
	s.Router.HandleFunc("/user/storie/archiveds", s.StoriesArchiveds).Methods("POST")
	s.Router.HandleFunc("/user/storie/archiveds/{id}", s.GetStoriesArchiveds).Methods("GET")
	// Habilitar CORS
	s.Router.Use(mux.CORSMethodMiddleware(s.Router))
	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"content-type"}),
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"}),
		handlers.AllowCredentials(),
	)
	s.Router.Use(cors)
	s.Router.Use(middlewares.SetMiddlewareJSON)
}
