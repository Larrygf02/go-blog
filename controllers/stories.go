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

func (s *Server) NewStorie(w http.ResponseWriter, r *http.Request) {
	var storie models.Storie
	err := json.NewDecoder(r.Body).Decode(&storie)
	if err != nil {
		fmt.Fprintf(w, "Inserte una historia valida")
	}
	storieCreated, err := storie.SaveStorie(s.DB)
	if err != nil {
		send_response.ERROR(w, http.StatusInternalServerError, err)
	}
	send_response.JSON(w, http.StatusCreated, storieCreated)
}

func (s *Server) StorieByUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	stories, err := user.GetStories(s.DB)
	if err != nil {
		send_response.ERROR(w, http.StatusInternalServerError, err)
	}
	send_response.JSON(w, http.StatusOK, stories)
}

/* STORIEVISIT*/
func (s *Server) NewStorieVisit(w http.ResponseWriter, r *http.Request) {
	var storie_visit models.StorieVisit
	err := json.NewDecoder(r.Body).Decode(&storie_visit)
	if err != nil {
		fmt.Fprintf(w, "Error en la data")
	}
	storieVisitCreated, err := storie_visit.SaveStorieVisit(s.DB)
	if err != nil {
		send_response.ERROR(w, http.StatusInternalServerError, err)
	}

	send_response.JSON(w, http.StatusOK, storieVisitCreated)
}

func (s *Server) GetAllStorieVisit(w http.ResponseWriter, r *http.Request) {
	var storie_visit models.StorieVisit
	stories_visit, count := storie_visit.GetAll(s.DB)
	type resp struct {
		Count int                  `json:"count"`
		Data  []models.StorieVisit `json:"data"`
	}
	response := resp{
		Count: count,
		Data:  *stories_visit,
	}
	send_response.JSON(w, http.StatusOK, response)
}

/* StorieApplauses */
func (s *Server) SaveStorieApplause(w http.ResponseWriter, r *http.Request) {
	var storie_applaus models.StorieApplause
	err := json.NewDecoder(r.Body).Decode(&storie_applaus)
	if err != nil {
		fmt.Fprintf(w, "Error en la data")
	}
	storieApplausCreated, err := storie_applaus.Save(s.DB)
	if err != nil {
		send_response.ERROR(w, http.StatusInternalServerError, err)
	}
	send_response.JSON(w, http.StatusCreated, storieApplausCreated)
}

/* StorieComment */
func (s *Server) SaveStorieComment(w http.ResponseWriter, r *http.Request) {
	var storie_comment models.StorieComment
	err := json.NewDecoder(r.Body).Decode(&storie_comment)
	if err != nil {
		fmt.Fprintf(w, "Error en la data")
	}
	storieCommentCreated, err := storie_comment.Save(s.DB)
	if err != nil {
		send_response.ERROR(w, http.StatusInternalServerError, err)
	}
	send_response.JSON(w, http.StatusCreated, storieCommentCreated)
}

func (s *Server) UpdateStorieComment(w http.ResponseWriter, r *http.Request) {
	s.DB.LogMode(true)
	parameters := mux.Vars(r)
	id, err := strconv.Atoi(parameters["id"])
	if err != nil {
		send_response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	// json.NewDecoder(r.Body).Decode(&body)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		send_response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	var storie_comment models.StorieComment
	json.Unmarshal(body, &storie_comment)
	storie_comment.ID = id
	err = storie_comment.Update(s.DB)
	if err != nil {
		send_response.ERROR(w, http.StatusInternalServerError, err)
	}
	updated, _ := storie_comment.Get(s.DB)
	/* var body models.StorieComment
	fmt.Println(&body)
	var test models.StorieComment
	s.DB.Model(&test).Where("id = ?", id).Updates(models.StorieComment{Content: body.Content}) */
	// test other method
	send_response.JSON(w, http.StatusOK, updated)
}
