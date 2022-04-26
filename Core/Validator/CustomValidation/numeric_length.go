package Validator

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

func UintLen10(fl validator.FieldLevel) bool {
	inter := fl.Field()
	slice, ok := inter.Interface().(uint)
	if !ok {
		fmt.Println("NumLength1 failed")
	}

	count := 0
	for slice > 0 {
		slice = slice/10
		count++
	}
	return count <= 10
}