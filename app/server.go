package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

type AppConfig struct {
	AppName string
	AppEnv  string
	AppPort string
}

func (server *Server) Initialize(appConfig AppConfig) {
	fmt.Println("Welcome to " + appConfig.AppName)

	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", getEnv("DB_USERNAME", ""), getEnv("DB_PASSWORD", ""), getEnv("DB_HOST",  "localhost"), getEnv("DB_PORT", "3336"), getEnv("DB_NAME", ""))
	server.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed on connecting to the database server")
	}

	server.Router = mux.NewRouter()
	server.InitializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Printf("Listening to port %s", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func Run() {
	var server    = Server{}
	var appConfig = AppConfig{}

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error on loading .env file")
	}

	appConfig.AppName = getEnv("APP_NAME", "GoShop")
	appConfig.AppEnv  = getEnv("APP_ENV", "development")
	appConfig.AppPort = getEnv("APP_PORT", "9999")

	server.Initialize(appConfig)
	server.Run(":" + appConfig.AppPort)
}