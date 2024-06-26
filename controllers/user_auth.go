package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"

	"rest_api/domain"
	"rest_api/lib/custom_validator"
)

type UserAuthService interface {
	GetUserByLogin(context.Context, string) (domain.User, error)
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

// @summary		Register
// @tags			user_auth
// @Description	Регистрация пользователя
// @ID				register
// @Accept			json
// @Produce		json
// @Param			req	body		domain.UserRegisterRequest	true	"Данные пользователя"
// @Success		200	{object}	domain.UserRegisterResponse
// @Router			/register [post]
func (uc *UserController) Register(c echo.Context) error {
	var userReg domain.UserRegisterRequest
	if err := c.Bind(&userReg); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"error": fmt.Sprintf("Failed to parse JSON: %s", err.Error()),
		})
	}
	if err := custom_validator.IsValid[domain.UserRegisterRequest](userReg); err != nil {
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

// @summary		Authorize
// @tags			user_auth
// @Description	Авторизация пользователя
// @ID				authorize
// @Accept			json
// @Produce		json
// @Param			req	body	domain.UserAuthorizeRequest	true	"Данные пользователя"
// @Router			/authorize [post]
func (uc *UserController) Authorize(c echo.Context) error {
	var req domain.UserAuthorizeRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"error": fmt.Sprintf("Failed to parse JSON: %s", err.Error()),
		})
	}
	if err := custom_validator.IsValid[domain.UserAuthorizeRequest](req); err != nil {
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
	t, err := uc.jas.CreateToken(user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": fmt.Sprintf("failed to create token: %s", err.Error()),
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}
