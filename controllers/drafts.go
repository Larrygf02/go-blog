package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/larrygf02/go-blog/models"
	send_response "github.com/larrygf02/go-blog/response"
)

func (s *Server) NewDraft(w http.ResponseWriter, r *http.Request) {
	var draft models.Draft
	err := json.NewDecoder(r.Body).Decode(&draft)
	if err != nil {
		send_response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	created, err := draft.Save(s.DB)
	if err != nil {
		send_response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	send_response.JSON(w, http.StatusCreated, created)
}
