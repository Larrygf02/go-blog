package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/larrygf02/go-blog/models"
	send_response "github.com/larrygf02/go-blog/response"
)

type Body struct {
	ID        int     `json:"id"`
	StoriesID []int64 `json:"stories_id"`
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
	// agregar
	user_found, _ := user.GetByID(s.DB)
	fmt.Println(user_found.Favorites)
	fmt.Println(body.StoriesID)
	//updated, err := user.SaveFavorites(s.DB, body)
	if err != nil {
		send_response.ERROR(w, http.StatusInternalServerError, err)
	}
	send_response.JSON(w, http.StatusOK, user_found)
}
