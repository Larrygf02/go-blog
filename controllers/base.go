package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/larrygf02/go-blog/models"
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
	s.DB.Debug().AutoMigrate(&models.User{}, &models.Storie{}, &models.StorieVisit{},
		&models.StorieApplause{}, &models.StorieComment{}, &models.Draft{})
	s.Router = mux.NewRouter()
	s.InitializeRoutes()
}

func (s *Server) Run(port string) {
	log.Fatal(http.ListenAndServe(port, s.Router))
}
