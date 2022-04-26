package Validator
import (
    "fmt"
    "reflect"
    "strings"
    "github.com/go-playground/validator/v10"
	ut "github.com/go-playground/universal-translator"
)
type ErrResponse struct {
    Errors []string `json:"errors"`
}
func New() *validator.Validate {
    validate := validator.New()
    validate.SetTagName("form")
    // Using the names which have been specified for JSON representations of structs, rather than normal Go field names
    validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
        name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
        if name == "-" {
            return ""
        }
        return name
    })
    return validate
}
func ToErrResponse(err error, trans ut.Translator) *ErrResponse {
    if fieldErrors, ok := err.(validator.ValidationErrors); ok {
        resp := ErrResponse{
            Errors: make([]string, len(fieldErrors)),
        }
        for i, err := range fieldErrors {
            switch err.Tag() {
			case "email":
                resp.Errors[i] = fmt.Sprintf("Invalid email address")
            case "uint-len-10":
                resp.Errors[i] = fmt.Sprintf("%s numeric max length is 10", err.Field())
            case "int-len-11":
                resp.Errors[i] = fmt.Sprintf("%s numeric max length is 11", err.Field())
            default:
				translatedErr := fmt.Errorf(err.Translate(trans))
				resp.Errors[i] = translatedErr.Error()
            }
        }
        return &resp
    }
    return nil
}