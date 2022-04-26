package Public

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	en_translations "github.com/go-playground/validator/v10/translations/en"

	"golang-restAPI-JWT/Core/Structs"
	"golang-restAPI-JWT/Core/Validator"
	cv "golang-restAPI-JWT/Core/Validator/CustomValidation"
	"golang-restAPI-JWT/Core/Api"
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

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "welcome to restAPI",
			"userInfo": "Hello World!!!",
		})
	})

	// DEFAULT ROUTE
	router.POST("/register", func(c *gin.Context) {
		// using BindJson method to serialize body with struct
		if err := c.BindJSON(&reg_req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"responseCode": 201,
				"error": err.Error(),
			})
			reg_req = Structs.RegisterRequest{}
			return
		}

		if err := validate.Struct(reg_req); err != nil {
			errs := Validator.ToErrResponse(err, trans)
			c.JSON(http.StatusBadRequest, gin.H{
				"responseCode": 202,
				"error": errs,
			})
			reg_req = Structs.RegisterRequest{}
			return
		}

		reg_res = Api.RegisterUser(reg_req)
		c.JSON(reg_res.ResponseCode,&reg_res)
		reg_req = Structs.RegisterRequest{}
	})


	router.POST("/login", func(c *gin.Context) {
		// using BindJson method to serialize body with struct
		if err := c.BindJSON(&log_req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"responseCode": 201,
				"error": err.Error(),
			})
			log_req = Structs.LoginRequest{}
			return
		}

		if err := validate.Struct(log_req); err != nil {
			errs := Validator.ToErrResponse(err, trans)
			c.JSON(http.StatusBadRequest, gin.H{
				"responseCode": 202,
				"error": errs,
			})
			log_req = Structs.LoginRequest{}
			return
		}

		log_res = Api.LoginUser(log_req)
		c.JSON(log_res.ResponseCode,&log_res)
		log_req = Structs.LoginRequest{}
	})
	// // END DEFAULT ROUTE
}
