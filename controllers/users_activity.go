package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/larrygf02/go-blog/models"
	send_response "github.com/larrygf02/go-blog/response"
	"github.com/larrygf02/go-blog/utils"
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
	favorites := utils.AppendInt(user_found.Favorites, body.StoriesID)
	user.Favorites = favorites
	updated, err := user.SaveFavorites(s.DB, body)
	if err != nil {
		send_response.ERROR(w, http.StatusInternalServerError, err)
	}
	send_response.JSON(w, http.StatusOK, updated)
}
