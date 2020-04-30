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

type StorieVisit struct {
	gorm.Model
	User     User `gorm:"foreignkey:UserId;not null"`
	UserId   uint
	Storie   Storie `gorm:"foreignkey:StorieId;not null"`
	StorieId uint
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
	json.NewEncoder(w).Encode(stories)
}

func NewStorieVisit(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("postgres", "host=localhost sslmode=disable port=5433 user=postgres dbname=bloggo password=123")
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()
	var storie_visit StorieVisit
	err = json.NewDecoder(r.Body).Decode(&storie_visit)
	if err != nil {
		fmt.Fprintf(w, "Error en la data de la visita")
	}
	err = db.Create(&storie_visit).Error
	if err != nil {
		fmt.Fprintf(w, "No se agrego la visita")
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "New storie created")
}

func GetAllStorieVisit(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("postgres", "host=localhost sslmode=disable port=5433 user=postgres dbname=bloggo password=123")
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()
	var stories_visit []Storie
	db.Find(&stories_visit)
	json.NewEncoder(w).Encode(stories_visit)
}
