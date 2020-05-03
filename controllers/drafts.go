package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func (s *Server) UpdateDraft(w http.ResponseWriter, r *http.Request) {
	s.DB.LogMode(true)
	parameters := mux.Vars(r)
	id, err := strconv.Atoi(parameters["id"])
	if err != nil {
		send_response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		send_response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	var draft models.Draft
	json.Unmarshal(body, &draft)
	draft.ID = id
	updated, err := draft.Update(s.DB)
	if err != nil {
		send_response.ERROR(w, http.StatusInternalServerError, err)
	}
	send_response.JSON(w, http.StatusOK, updated)
}

func (s *Server) DraftByUser(w http.ResponseWriter, r *http.Request) {
	s.DB.LogMode(true)
	var user models.User
	parameters := mux.Vars(r)
	id, err := strconv.Atoi(parameters["user_id"])
	fmt.Println(id)
	if err != nil {
		send_response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	s.DB.First(&user, id)
	//err := json.NewDecoder(r.Body).Decode(&user)
	drafts, err := user.GetDrafts(s.DB)
	if err != nil {
		send_response.ERROR(w, http.StatusInternalServerError, err)
	}
	send_response.JSON(w, http.StatusOK, drafts)
}
