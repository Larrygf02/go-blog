package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/larrygf02/go-blog/models"
	send_response "github.com/larrygf02/go-blog/response"
)

func (s *Server) NewUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Fprintf(w, "Inserte un Usuario valido")
	}
	fmt.Println(user)
	// Validate User
	if user.Nickname == "" {
		err := errors.New("Nickname is required")
		send_response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	if user.Name == "" {
		err := errors.New("Name is required")
		send_response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	if user.Email == "" {
		err := errors.New("Email is required")
		send_response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	if user.Password == "" {
		err := errors.New("Password is required")
		send_response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	userCreated, err := user.SaveUser(s.DB)
	if err != nil {
		send_response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	send_response.JSON(w, http.StatusCreated, userCreated)
}

func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		send_response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	userFind, isLogin := user.Login(s.DB)
	fmt.Println(isLogin)
	if isLogin {
		// Return JWT
	}
	// Las propiedas de las structuras deben estar en mayusculas
	// debe tener las dobles comillas en el json
	type resp struct {
		IsLogin bool        `json:"is_login"`
		User    models.User `json:"user"`
	}
	response := resp{
		IsLogin: isLogin,
		User:    *userFind,
	}
	send_response.JSON(w, http.StatusOK, response)
}

func (s *Server) UserNameValid(w http.ResponseWriter, r *http.Request) {
	s.DB.LogMode(true)
	var user models.User
	parameters := mux.Vars(r)
	fmt.Println(parameters)
	//nickname, _ := strconv.Atoi(parameters["nickname"])
	//user.Nickname = username
	err := s.DB.Where("nickname = ?", parameters["nickname"]).First(&user).Error
	type resp struct {
		IsValid bool `json:"is_valid"`
	}
	response := resp{}
	// No existe
	if err != nil {
		response = resp{
			IsValid: true,
		}
		send_response.JSON(w, http.StatusOK, response)
		return
	}
	// Ya existe
	response = resp{
		IsValid: false,
	}
	send_response.JSON(w, http.StatusOK, response)
}
