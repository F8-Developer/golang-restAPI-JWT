package validator

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	ut "github.com/go-playground/universal-translator"
)

func TranslateError(err error, trans ut.Translator) (errs []string) {
	if err == nil {
	  	return nil
	}
	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		translatedErr := fmt.Errorf(e.Translate(trans))
		errs = append(errs, translatedErr.Error())
	}
	return errs
}