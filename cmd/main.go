package main

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"log"
	"rest_api/controllers"
	"rest_api/internal/repository"
	"rest_api/internal/service"
)

func main() {
	db, err := sqlx.Connect("postgres",
		"postgres://postgres:postgres@localhost:5432/rest_api_db?sslmode=disable")
	if err != nil {
		log.Fatal("Failed to create database connection")
	}
	defer db.Close()

	ur := repository.NewUserRepository(db)
	us := service.NewUserService(ur)
	e := echo.New()
	defer e.Close()
	controllers.NewUserController(e, us)
	if err := e.Start(":8080"); err != nil {
		log.Println("Failed to start server")
		return
	}
}
