package custom_validator

import (
	"github.com/go-playground/validator/v10"
	"rest_api/domain"
	"unicode"
)

func newCustomValidator() *validator.Validate {
	validate := validator.New()
	validate.RegisterValidation("isDifficultPassword", isDifficultPassword)
	return validate
}

func IsValid[T domain.UserRegisterRequest | domain.UserAuthorizeRequest | domain.AdvertPostRequest](t T) error {
	validate := newCustomValidator()
	err := validate.Struct(t)
	if err != nil {
		// TODO: parse valid errors
		// validationErrors := err.(validator.ValidationErrors)
		return err
	}
	return nil
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
