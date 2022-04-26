package Private

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	en_translations "github.com/go-playground/validator/v10/translations/en"

	"golang-restAPI-JWT/Middleware"
	"golang-restAPI-JWT/Core/Structs"
	"golang-restAPI-JWT/Core/Validator"
	cv "golang-restAPI-JWT/Core/Validator/CustomValidation"
)

var (
	reg_req Structs.RegisterRequest
	reg_res Structs.RegisterResponse
	log_req Structs.LoginRequest
	log_res Structs.LoginResponse
)

// APIRouter define router from here, you can add new api about your new services.
func APIRouter(router *gin.Engine) {
	// set validator
	validate := Validator.InitValidator()
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	validate.RegisterValidation("email", cv.Email)
	_ = en_translations.RegisterDefaultTranslations(validate, trans)
	// end set validator

	authorized := router.Group("/secure").Use(Middleware.Auth())
	// /admin/secrets endpoint
	authorized.GET("/category", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "welcome to restAPI",
			"userInfo": "Hello World!!!",
		})
	})
}