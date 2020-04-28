package views

import (
	"encoding/json"
	"fmt"
	"net/http"

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
	var user User
	err = json.NewDecoder(r.Body).Decode(&user)
	//reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Inserte un Usuario valido")
	}
	db.Create(&user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "New user successfully created")
}
