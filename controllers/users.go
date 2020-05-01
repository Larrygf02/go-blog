package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/larrygf02/go-blog/models"
	send_response "github.com/larrygf02/go-blog/response"
)

// Las propiedas de las structuras deben estar en mayusculas
// debe tener las dobles comillas en el json
type resp struct {
	IsLogin bool        `json:"is_login"`
	User    models.User `json:"user"`
}

func (s *Server) NewUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Fprintf(w, "Inserte un Usuario valido")
	}
	userCreated, err := user.SaveUser(s.DB)
	if err != nil {
		send_response.ERROR(w, http.StatusInternalServerError, err)
	}
	send_response.JSON(w, http.StatusCreated, userCreated)
}

func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Fprintf(w, "Hubo un problema")
	}

	userFind, isLogin := user.Login(s.DB)
	fmt.Println(isLogin)
	response := resp{
		IsLogin: isLogin,
		User:    *userFind,
	}
	json.NewEncoder(w).Encode(response)
	//send_response.JSON(w, http.StatusOK, response)
}
