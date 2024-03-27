package main

import (
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	"log"
	"os"

	"rest_api/controllers"
	_ "rest_api/docs"
	"rest_api/internal/repository"
	"rest_api/internal/service"
)

const (
	defaultAddress = ":8080"
	defaultDBURI   = "postgres://postgres:postgres@localhost:5432/rest_api_db?sslmode=disable"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

//	@title		Rest API
//	@version	1

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
//	@description				Bearer + JWT

func main() {
	dbURI := os.Getenv("DATABASE_URI")
	if dbURI == "" {
		dbURI = defaultDBURI
		log.Println("Failed to get database URI, used default value")
	}
	db, err := sqlx.Connect("postgres", dbURI)
	if err != nil {
		log.Fatal("Failed to create database connection")
	}
	defer db.Close()

	e := echo.New()
	defer e.Close()

	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	ur := repository.NewUserAuthRepository(db)
	us := service.NewUserService(ur)
	jas := service.NewJWTAuthService()
	controllers.NewUserController(e, us, jas)

	ar := repository.NewAdvertRepository(db)
	as := service.NewAdvertService(ar)
	controllers.NewAdvertController(e, as)

	address := os.Getenv("APP_ADDRESS")
	if address == "" {
		address = defaultAddress
		log.Println("Failed to get application address, used default value")
	}

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	if err := e.Start(address); err != nil {
		log.Println("Failed to start server")
		return
	}
}
