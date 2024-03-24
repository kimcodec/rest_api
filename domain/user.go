package domain

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

type User struct {
	Login    string `db:"login"`
	Password string `db:"password"`
}

type UserAuthorizeRequest struct {
	Login    string `json:"login" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// TODO: fix validation

type UserRegisterRequest struct {
	Login    string `json:"login" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func isDifficultPassword(fldLvl validator.FieldLevel) bool {
	// Minimum eight characters, at least one uppercase letter, one lowercase letter and one number
	reg := regexp.MustCompile("^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])[a-zA-Z0-9]{8,}$")
	pass, ok := fldLvl.Field().Interface().(string)
	if !ok {
		return false
	}
	return reg.MatchString(pass)
}

func IsValid[T UserRegisterRequest | UserAuthorizeRequest](t T) error {
	validate := validator.New()
	err := validate.Struct(t)
	if err != nil {
		// TODO: parse valid errors
		// validate := validator.New(validator.WithRequiredStructEnabled())
		return err
	}
	return nil
}
