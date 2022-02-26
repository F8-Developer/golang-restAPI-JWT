package validator

import (
	"fmt"
	"strconv"
	"strings"
	"github.com/go-playground/validator/v10"
)

func IntLen1(fl validator.FieldLevel) bool {
	inter := fl.Field()
	slice, ok := inter.Interface().(int)
	if !ok {
		fmt.Println("NumLength1 failed")
	}

	count := 0
	for slice > 0 {
		slice = slice/10
		count++
	}
	return count <= 1
}

func IntLen5(fl validator.FieldLevel) bool {
	inter := fl.Field()
	slice, ok := inter.Interface().(int)
	if !ok {
		fmt.Println("NumLength1 failed")
	}

	count := 0
	for slice > 0 {
		slice = slice/10
		count++
	}
	return count <= 5
}

func IntLen10(fl validator.FieldLevel) bool {
	inter := fl.Field()
	slice, ok := inter.Interface().(int)
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

func Uint64Len18(fl validator.FieldLevel) bool {
	inter := fl.Field()
	slice, ok := inter.Interface().(uint64)
	if !ok {
		fmt.Println("NumLength1 failed")
	}

	count := 0
	for slice > 0 {
		slice = slice/10
		count++
	}
	return count <= 18
}

func FloatDecimalVal(fl validator.FieldLevel) bool {
	inter := fl.Field()
	slice, ok := inter.Interface().(float64)
	if !ok {
		fmt.Println("NumLength1 failed")
	}
	if !NumDecPlaces(slice) {
		return false
	}

	count := 0
	slice_int := uint64(slice)
	for slice_int > 0 {
		slice_int = slice_int/10
		count++
	}
	return count <= 15
}

func NumDecPlaces(v float64) bool {
	s := strconv.FormatFloat(v, 'f', -1, 64)
	i := strings.IndexByte(s, '.')

	if i > -1 {
		count_length := len(s) - i - 1
		return count_length <= 2
	}
	return true
}