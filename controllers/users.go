package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/larrygf02/go-blog/models"
	"github.com/larrygf02/go-blog/response"
)

func (s *Server) NewUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Fprintf(w, "Inserte un Usuario valido")
	}
	userCreated, err := user.SaveUser(s.DB)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
	}
	response.JSON(w, http.StatusCreated, userCreated)
}
