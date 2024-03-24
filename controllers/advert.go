package controllers

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"net/http"
	"rest_api/domain"
	"strconv"
)

type AdvertService interface {
	Post(context.Context, domain.AdvertToPost) (domain.AdvertPostResponse, error)
	List(context.Context, uint64, uint64) ([]domain.AdvertListResponse, error)
}

type AdvertController struct {
	as AdvertService
}

func NewAdvertController(e *echo.Echo, as AdvertService) {
	a := e.Group("/advert")
	{
		config := echojwt.Config{
			NewClaimsFunc: func(c echo.Context) jwt.Claims {
				return new(domain.JWTCustomClaims)
			},
			SigningKey: []byte(domain.JWTKey),
		}
		a.Use(echojwt.WithConfig(config))
	}
}

func (ac *AdvertController) Post(c echo.Context) error {
	var ad domain.AdvertPostRequest
	if err := c.Bind(&ad); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, fmt.Sprintf("Failed to parse JSON: %s", err.Error()))
	}
	price, err := strconv.Atoi(ad.Price)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Failed to parse price: %s", err.Error()))
	}
	advert := domain.AdvertToPost{
		Title:    ad.Title,
		Text:     ad.Text,
		ImageURL: ad.ImageURL,
		Price:    uint64(price),
	}
	ctx := c.Request().Context()
	// TODO: добавить проверку формата и размера картинки
	adResp, err := ac.as.Post(ctx, advert)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Failed to post advert: %s", err.Error()))
	}
	return c.JSON(http.StatusOK, adResp)
}

func (ac *AdvertController) List(c echo.Context) error {
	// TODO: передавать offset и limit в query
	// TODO: передавать тип сортирвки и цены по query
	// TODO: сделать
	return nil
}
