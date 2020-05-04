package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/larrygf02/go-blog/models"
	send_response "github.com/larrygf02/go-blog/response"
	"github.com/larrygf02/go-blog/utils"
)

type Body struct {
	ID        int     `json:"id"`
	StoriesID []int64 `json:"stories_id"`
	Type      string  `json:"type"`
}

func (s *Server) StoriesFavorites(w http.ResponseWriter, r *http.Request) {
	s.DB.LogMode(true)
	var body Body
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		send_response.ERROR(w, http.StatusInternalServerError, err)
	}
	var user models.User
	user.ID = body.ID
	user_found, _ := user.GetByID(s.DB)
	var favorites []int64
	switch body.Type {
	case "delete":
		fmt.Println("Delete item")
		favorites = utils.DeleteItemsInt(user_found.Favorites, body.StoriesID)
	case "add":
		fmt.Println("Add Item")
		favorites = utils.AppendInt(user_found.Favorites, body.StoriesID)
	default:
		send_response.ERROR(w, http.StatusInternalServerError, nil)
		return
	}
	user.Favorites = favorites
	updated, err := user.SaveFavorites(s.DB, body)
	if err != nil {
		send_response.ERROR(w, http.StatusInternalServerError, err)
	}
	send_response.JSON(w, http.StatusOK, updated)
}

func (s *Server) GetStoriesFavorites(w http.ResponseWriter, r *http.Request) {
	s.DB.LogMode(true)
	var user models.User
	parameters := mux.Vars(r)
	id, _ := strconv.Atoi(parameters["id"])
	user.ID = id
	userFound, exists := user.GetByID(s.DB)
	if !exists {
		send_response.ERROR(w, http.StatusNotFound, nil)
		return
	}
	var stories []models.Storie
	stories, err := userFound.GetFavorites(s.DB)
	if err != nil {
		send_response.ERROR(w, http.StatusNotFound, nil)
		return
	}
	send_response.JSON(w, http.StatusOK, stories)
}

func (s *Server) StoriesArchiveds(w http.ResponseWriter, r *http.Request) {
	s.DB.LogMode(true)
	var body Body
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		send_response.ERROR(w, http.StatusInternalServerError, err)
	}
	var user models.User
	user.ID = body.ID
	user_found, _ := user.GetByID(s.DB)
	var archiveds []int64
	switch body.Type {
	case "delete":
		fmt.Println("Delete item")
		archiveds = utils.DeleteItemsInt(user_found.Archiveds, body.StoriesID)
	case "add":
		fmt.Println("Add Item")
		archiveds = utils.AppendInt(user_found.Archiveds, body.StoriesID)
	default:
		send_response.ERROR(w, http.StatusInternalServerError, nil)
		return
	}
	user.Archiveds = archiveds
	updated, err := user.SaveArchiveds(s.DB, body)
	if err != nil {
		send_response.ERROR(w, http.StatusInternalServerError, err)
	}
	send_response.JSON(w, http.StatusOK, updated)
}

func (s *Server) GetStoriesArchiveds(w http.ResponseWriter, r *http.Request) {
	s.DB.LogMode(true)
	var user models.User
	parameters := mux.Vars(r)
	id, _ := strconv.Atoi(parameters["id"])
	user.ID = id
	userFound, exists := user.GetByID(s.DB)
	if !exists {
		send_response.ERROR(w, http.StatusNotFound, nil)
		return
	}
	var stories []models.Storie
	stories, err := userFound.GetArchiveds(s.DB)
	if err != nil {
		send_response.ERROR(w, http.StatusNotFound, nil)
		return
	}
	send_response.JSON(w, http.StatusOK, stories)
}
