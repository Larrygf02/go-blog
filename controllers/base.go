package controllers

import (
	"fmt"
	"log"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/larrygf02/go-blog/views"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (s *Server) Initialize(DBDriver, DBUser, DBPassword, DBPort, DBHost, DBName string) {
	var err error
	// Conectando con base de datos POSTGRES
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", DBHost, DBPort, DBUser, DBName, DBPassword)
	fmt.Println(DBURL)
	s.DB, err = gorm.Open(DBDriver, DBURL)
	if err != nil {
		log.Fatal("This is the error:", err)
	}
	s.DB.Debug().AutoMigrate(&views.User{}, &views.Storie{}, &views.StorieVisit{})
	s.Router = mux.NewRouter()
	s.InitializeRoutes()
}
