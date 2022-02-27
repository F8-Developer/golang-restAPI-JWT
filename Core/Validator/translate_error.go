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
            case "int-length-1":
                resp.Errors[i] = fmt.Sprintf("%s numeric max length is 1%s", err.Field(), err.Param())
            case "int-length-5":
                resp.Errors[i] = fmt.Sprintf("%s numeric max length is 5%s", err.Field(), err.Param())
            case "int-length-10":
                resp.Errors[i] = fmt.Sprintf("%s numeric max length is 10%s", err.Field(), err.Param())
            case "uint64-length-18":
                resp.Errors[i] = fmt.Sprintf("%s numeric max length is 18%s", err.Field(), err.Param())
            case "float-decimal-val":
                resp.Errors[i] = fmt.Sprintf("%s numeric decimal format (two decimal), ex : 1000.00%s", err.Field(), err.Param())
            case "date-format":
                resp.Errors[i] = fmt.Sprintf("%s date format invalid, ex : YYYY-MM-DD%s", err.Field(), err.Param())
            default:
				translatedErr := fmt.Errorf(err.Translate(trans))
				resp.Errors[i] = translatedErr.Error()
            }
        }
        return &resp
    }
    return nil
}