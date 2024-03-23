package controllers

import (
	"context"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"rest_api/domain"
)

type UserService interface {
	GetUserByEmail(context.Context, string) (domain.User, error)
	Register(context.Context, domain.UserRegisterRequest) error
}

type UserController struct {
	us UserService
}

func NewUserController(e *echo.Echo, us UserService) {
	uc := &UserController{
		us: us,
	}
	e.POST("/authorize", uc.Authorize)
	e.POST("/register", uc.Register)
}

func (uc *UserController) Register(c echo.Context) error {
	var userReg domain.UserRegisterRequest
	if err := c.Bind(&userReg); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if err := domain.IsValid[domain.UserRegisterRequest](userReg); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	if err := uc.us.Register(ctx, userReg); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, "Success")
}

func (uc *UserController) Authorize(c echo.Context) error {
	var req domain.UserAuthorizeRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	if err := domain.IsValid[domain.UserAuthorizeRequest](req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	user, err := uc.us.GetUserByEmail(ctx, req.Login)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, "Invalid password or email")
	}
	//TODO: JWT-token
	return c.JSON(http.StatusOK, "Success")
}
