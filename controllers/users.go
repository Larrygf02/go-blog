package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/larrygf02/go-blog/models"
)

func (s *Server) NewUser(w http.ResponseWriter, r *http.Request) {
	/* db, err := gorm.Open("postgres", "host=localhost sslmode=disable port=5433 user=postgres dbname=bloggo password=123")
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close() */
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Fprintf(w, "Inserte un Usuario valido")
	}
	userCreated, err := user.SaveUser(s.DB)
	if err != nil {
		panic("Not inserted")
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "New user successfully created")
}
