package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
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
	// Where ORIGIN_ALLOWED is like `scheme://dns[:port]`, or `*` (insecure)
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	log.Fatal(http.ListenAndServe(port, handlers.CORS(originsOk, headersOk, methodsOk)(s.Router)))
}
