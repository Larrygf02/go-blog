package views

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Nickname string `gorm:"type:varchar(70);unique;not null"`
	Name     string `gorm:"type:varchar(50)"`
	Email    string `gorm:"type:varchar(60);not null"`
	Password string `gorm:"type: varchar(80); not null"`
}

/* func (u *User) init() {
	fmt.Println("Init USER")
	db, err := gorm.Open("postgres", "host=localhost sslmode=disable port=5433 user=postgres dbname=bloggo password=123")
	if err != nil {
		panic("Could not connect to the database")
	}
	err = db.Model(&User{}).ModifyColumn("name", "varchar").Error
	if err != nil {
		panic("Error in migration")
	}
	err = db.Model(&User{}).ModifyColumn("email", "varchar(60)").Error
	if err != nil {
		panic("Error in migration")
	}
	defer db.Close()
} */

func Login(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("postgres", "host=localhost sslmode=disable port=5433 user=postgres dbname=bloggo password=123")
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()
	var userFind User
	var user User
	count := 0
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Fprintf(w, "There are a problem with data")
	}
	db.Where(&user).First(&userFind).Count(&count)
	if count != 0 {
		fmt.Fprintf(w, "User is login")
	} else {
		fmt.Fprintf(w, "User is not login")
	}
}
