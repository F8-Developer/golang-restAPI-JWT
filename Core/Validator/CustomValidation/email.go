package Validator

import (
	"github.com/go-playground/validator/v10"
	"net/mail"
)

func Email(fl validator.FieldLevel) bool {
	val := fl.Field().Interface().(string)
	_, err := mail.ParseAddress(val)
	if err != nil {
		return false
	}
	return true
}