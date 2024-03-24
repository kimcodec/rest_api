package controllers

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"rest_api/domain"
)

type UserAuthService interface {
	GetUserByLogin(context.Context, string) (domain.UserGetByLogin, error)
	Register(context.Context, domain.UserRegisterRequest) (domain.UserRegisterResponse, error)
}

type JWTAuthService interface {
	CreateToken(id uint64) (string, error)
}

type UserController struct {
	us  UserAuthService
	jas JWTAuthService
}

func NewUserController(e *echo.Echo, us UserAuthService, jas JWTAuthService) {
	uc := &UserController{
		us:  us,
		jas: jas,
	}
	e.POST("/authorize", uc.Authorize)
	e.POST("/register", uc.Register)
}

func (uc *UserController) Register(c echo.Context) error {
	var userReg domain.UserRegisterRequest
	if err := c.Bind(&userReg); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"error": fmt.Sprintf("Failed to parse JSON: %s", err.Error()),
		})
	}
	if err := domain.IsValid[domain.UserRegisterRequest](userReg); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": fmt.Sprintf("Failed to validate fields: %s", err.Error()),
		})
	}

	hashPass, err := bcrypt.GenerateFromPassword([]byte(userReg.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": fmt.Sprintf("Failed to hash password: %s", err.Error()),
		})
	}
	userReg.Password = string(hashPass)

	ctx := c.Request().Context()
	user, err := uc.us.Register(ctx, userReg)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": fmt.Sprintf("failed to register user: %s", err.Error()),
		})
	}
	return c.JSON(http.StatusCreated, user)
}

func (uc *UserController) Authorize(c echo.Context) error {
	var req domain.UserAuthorizeRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"error": fmt.Sprintf("Failed to parse JSON: %s", err.Error()),
		})
	}
	if err := domain.IsValid[domain.UserAuthorizeRequest](req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": fmt.Sprintf("Failed to validate fields: %s", err.Error()),
		})
	}

	ctx := c.Request().Context()
	user, err := uc.us.GetUserByLogin(ctx, req.Login)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Failed to find a user: %s", err.Error()))
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "Invalid password or email",
		})
	}
	t, err := uc.jas.CreateToken(12)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": fmt.Sprintf("failed to create token: %s", err.Error()),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}
