package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/larrygf02/go-blog/views"
)

var db *gorm.DB
var err error

func InitialMigration() {
	db, err := gorm.Open("postgres", "host=localhost sslmode=disable port=5433 user=postgres dbname=bloggo password=123")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}
	fmt.Println("You are connect")
	defer db.Close()
	db.AutoMigrate(&views.User{})
}
