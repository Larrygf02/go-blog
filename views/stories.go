package views

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
)

type Storie struct {
	gorm.Model
	Title   string `gorm:"type:varchar(100);not null"`
	Content string
	User    User `gorm:"foreignkey:UserId"`
	UserId  uint
}

func NewStorie(w http.ResponseWriter, r *http.Request) {

	db, err := gorm.Open("postgres", "host=localhost sslmode=disable port=5433 user=postgres dbname=bloggo password=123")
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()
	var storie Storie
	err = json.NewDecoder(r.Body).Decode(&storie)
	if err != nil {
		fmt.Fprintf(w, "Inserte una Historia Valida")
	}
	err = db.Create(&storie).Error
	if err != nil {
		fmt.Fprintf(w, "No se inserto")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "New storie created")
}

func StorieByUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("postgres", "host=localhost sslmode=disable port=5433 user=postgres dbname=bloggo password=123")
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()
	var user User
	var stories []Storie
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Fprintf(w, "Operacion no valida")
	}
	db.Model(&user).Related(&stories)
	//fmt.Println(stories)
	w.Header().Set("Content-Type", "application/json")
	//fmt.Fprintf(w, "Data get obtein successfully")
	json.NewEncoder(w).Encode(stories)
}
