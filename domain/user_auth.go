package domain

import (
	"github.com/go-playground/validator/v10"
	"unicode"
)

type UserAuthorizeRequest struct {
	Login    string `json:"login" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserRegisterRequest struct {
	Login    string `json:"login" validate:"required,min=4,max=50"`
	Password string `json:"password" validate:"required,isDifficultPassword"`
}

type UserRegisterResponse struct {
	ID       string `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

func isDifficultPassword(fldLvl validator.FieldLevel) bool {
	// Minimum eight characters, at least one uppercase letter, one lowercase letter and one number, one symbol
	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	pass := fldLvl.Field().String()
	if len(pass) >= 8 {
		hasMinLen = true
	}
	for _, char := range pass {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}

func IsValid[T UserRegisterRequest | UserAuthorizeRequest](t T) error {
	validate := validator.New()
	validate.RegisterValidation("isDifficultPassword", isDifficultPassword)
	err := validate.Struct(t)
	if err != nil {
		// TODO: parse valid errors
		// validationErrors := err.(validator.ValidationErrors)
		return err
	}
	return nil
}

type UserGetByLogin struct {
	Login    string `db:"login"`
	Password string `db:"password"`
}

type UserDB struct {
	ID       string `db:"id"`
	Login    string `db:"login"`
	Password string `db:"password"`
}
