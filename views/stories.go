package views

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
)

type Storie struct {
	gorm.Model
	Title     string `gorm:"type:varchar(100);not null"`
	Content   string
	UserID    User `gorm:"foreignkey:UserRef"`
	UserRefer uint
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
