package views

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Nickname string
	Name     string
	Email    string
	Password string
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("postgres", "host=localhost sslmode=disable port=5433 user=postgres dbname=bloggo password=123")
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()
	vars := mux.Vars(r)
	nickname := vars["nickname"]
	name := vars["name"]
	email := vars["email"]
	password := vars["password"]

	db.Create(&User{Name: name, Email: email, Nickname: nickname, Password: password})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "New user successfully created")
}
