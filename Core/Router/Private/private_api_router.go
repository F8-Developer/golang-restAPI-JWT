package Private

import (
	"github.com/gin-gonic/gin"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	en_translations "github.com/go-playground/validator/v10/translations/en"

	"golang-restAPI-JWT/Middleware"
	"golang-restAPI-JWT/Core/Structs"
	"golang-restAPI-JWT/Core/Validator"
	cv "golang-restAPI-JWT/Core/Validator/CustomValidation"
	"golang-restAPI-JWT/Core/Api"
)

var (
	cas_res Structs.CategoriesResponse
	prs_res Structs.ProductsResponse
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
	//secure endpoint with jwt token
	authorized.POST("/category/list", func(c *gin.Context) {
		cas_res = Api.GetAllCategories()
		c.JSON(cas_res.ResponseCode,&cas_res)
	})
	
	authorized.POST("/product/list", func(c *gin.Context) {
		prs_res = Api.GetAllProducts()
		c.JSON(prs_res.ResponseCode,&prs_res)
	})
}