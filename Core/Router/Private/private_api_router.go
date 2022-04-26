package Private

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	en_translations "github.com/go-playground/validator/v10/translations/en"

	"golang-restAPI-JWT/Auth"
	"golang-restAPI-JWT/Middleware"
	"golang-restAPI-JWT/Core/Structs"
	"golang-restAPI-JWT/Core/Validator"
	cv "golang-restAPI-JWT/Core/Validator/CustomValidation"
	"golang-restAPI-JWT/Core/Api"
)

var (
	cas_res Structs.CategoriesResponse
	prs_res Structs.ProductsResponse
	prd_req Structs.ProductDetailRequest
	prd_res Structs.ProductDetailResponse
	ca_req Structs.CartAddRequest
	ca_res Structs.CartAddResponse
)

// APIRouter define router from here, you can add new api about your new services.
func APIRouter(router *gin.Engine) {
	// set validator
	validate := Validator.InitValidator()
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	validate.RegisterValidation("uint-len-10", cv.UintLen10)
	validate.RegisterValidation("int-len-11", cv.IntLen11)
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

	authorized.POST("/product/detail", func(c *gin.Context) {
		// using BindJson method to serialize body with struct
		if err := c.BindJSON(&prd_req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"responseCode": 201,
				"error": err.Error(),
			})
			prd_req = Structs.ProductDetailRequest{}
			return
		}

		if err := validate.Struct(prd_req); err != nil {
			errs := Validator.ToErrResponse(err, trans)
			c.JSON(http.StatusBadRequest, gin.H{
				"responseCode": 202,
				"error": errs,
			})
			prd_req = Structs.ProductDetailRequest{}
			return
		}

		prd_res = Api.GetProductDetail(prd_req)
		c.JSON(prd_res.ResponseCode,&prd_res)
		prd_req = Structs.ProductDetailRequest{}
	})

	authorized.POST("/cart/add", func(c *gin.Context) {
		// get user email from jwt
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA):]
		email, err := Auth.GetEmail(tokenString)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"responseCode": 201,
				"error": err.Error(),
			})
			ca_req = Structs.CartAddRequest{}
			return
		}

		// using BindJson method to serialize body with struct
		if err := c.BindJSON(&ca_req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"responseCode": 201,
				"error": err.Error(),
			})
			ca_req = Structs.CartAddRequest{}
			return
		}

		if err := validate.Struct(ca_req); err != nil {
			errs := Validator.ToErrResponse(err, trans)
			c.JSON(http.StatusBadRequest, gin.H{
				"responseCode": 202,
				"error": errs,
			})
			ca_req = Structs.CartAddRequest{}
			return
		}

		ca_res = Api.AddToCart(email, ca_req)
		c.JSON(ca_res.ResponseCode,&ca_res)
		ca_req = Structs.CartAddRequest{}
	})
}