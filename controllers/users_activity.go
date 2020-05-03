package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

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
