package controllers

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"rest_api/domain"
	"rest_api/lib"
	"rest_api/lib/custom_validator"
	"strconv"
)

type AdvertService interface {
	Post(context.Context, domain.AdvertToPost) (domain.AdvertPostResponse, error)
	List(context.Context, lib.AdvertQueryParams, uint64) ([]domain.AdvertListResponse, error)
}

type AdvertController struct {
	as AdvertService
}

func NewAdvertController(e *echo.Echo, as AdvertService) {
	a := e.Group("/advert")
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(domain.JWTCustomClaims)
		},
		SigningKey: []byte(domain.JWTKey),
	}
	a.Use(echojwt.WithConfig(config))

	ac := &AdvertController{
		as: as,
	}

	a.POST("", ac.Post)
	a.GET("", ac.List)
}

func (ac *AdvertController) Post(c echo.Context) error {
	user, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "failed to get token",
		})
	}
	claims, ok := user.Claims.(*domain.JWTCustomClaims)
	if !ok {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "failed to get claims",
		})
	}
	userID := claims.UserID
	log.Println(userID)

	var ad domain.AdvertPostRequest
	if err := c.Bind(&ad); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, fmt.Sprintf("Failed to parse JSON: %s", err.Error()))
	}

	if err := custom_validator.IsValid[domain.AdvertPostRequest](ad); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"error": fmt.Sprintf("Failed to falidate fields: %s", err.Error()),
		})
	}

	if err := lib.ValidatePicture(ad.ImageURL); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"error": fmt.Sprintf("failed to validate picture by url: %s", err.Error()),
		})
	}

	price, err := strconv.Atoi(ad.Price)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Failed to parse price: %s", err.Error()))
	}
	advert := domain.AdvertToPost{
		UserID:   userID,
		Title:    ad.Title,
		Text:     ad.Text,
		ImageURL: ad.ImageURL,
		Price:    uint64(price),
	}
	ctx := c.Request().Context()
	adResp, err := ac.as.Post(ctx, advert)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Failed to post advert: %s", err.Error()))
	}
	return c.JSON(http.StatusOK, adResp)
}

func (ac *AdvertController) List(c echo.Context) error {
	user, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "failed to get token",
		})
	}
	claims, ok := user.Claims.(*domain.JWTCustomClaims)
	if !ok {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "failed to get claims",
		})
	}
	userID := claims.UserID

	query := c.QueryParams()
	params, err := lib.ParseAdvertParams(query)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": fmt.Sprintf("Failed to get query params: %s", err.Error()),
		})
	}

	ctx := c.Request().Context()
	list, err := ac.as.List(ctx, params, userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": fmt.Sprintf("Failed to get list: %s", err.Error()),
		})
	}

	return c.JSON(http.StatusOK, list)
}
