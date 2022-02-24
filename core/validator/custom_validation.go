package validator

import (
	"github.com/go-playground/validator/v10"
)

func ValidateMyVal(fl validator.FieldLevel) bool {
	return fl.Field().String() == "awesome"
}