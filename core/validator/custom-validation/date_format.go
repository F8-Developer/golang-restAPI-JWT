package validator

import (
	"fmt"
	"time"
	"github.com/go-playground/validator/v10"
)

func DateFormat(fl validator.FieldLevel) bool {
	inter := fl.Field()
	slice, ok := inter.Interface().(string)
	if !ok {
		fmt.Println("string date-format failed")
	}

	_, err := time.Parse("2006-01-02", slice)
	if err != nil {
		return false
	}
	return true
}