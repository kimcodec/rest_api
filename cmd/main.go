package main

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
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

	e := echo.New()
	defer e.Close()

	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	/*a := e.Group("/advert")
	{
		config := echojwt.Config{
			NewClaimsFunc: func(c echo.Context) jwt.Claims {
				return new(domain.JWTCustomClaims)
			},
			SigningKey: []byte(domain.JWTKey),
		}
		a.Use(echojwt.WithConfig(config))
	}*/

	ur := repository.NewUserAuthRepository(db)
	us := service.NewUserService(ur)
	jas := service.NewJWTAuthService()
	controllers.NewUserController(e, us, jas)

	if err := e.Start(":8080"); err != nil {
		log.Println("Failed to start server")
		return
	}
}
