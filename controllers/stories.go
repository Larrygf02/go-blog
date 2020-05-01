package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/larrygf02/go-blog/models"
	send_response "github.com/larrygf02/go-blog/response"
)

func (s *Server) NewStorie(w http.ResponseWriter, r *http.Request) {
	var storie models.Storie
	err := json.NewDecoder(r.Body).Decode(&storie)
	if err != nil {
		fmt.Fprintf(w, "Inserte una historia valida")
	}
	storieCreated, err := storie.SaveStorie(s.DB)
	if err != nil {
		send_response.ERROR(w, http.StatusInternalServerError, err)
	}
	send_response.JSON(w, http.StatusCreated, storieCreated)
}

func (s *Server) StorieByUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	stories, err := user.GetStories(s.DB)
	if err != nil {
		send_response.ERROR(w, http.StatusInternalServerError, err)
	}
	send_response.JSON(w, http.StatusOK, stories)
}
