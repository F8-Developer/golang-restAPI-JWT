package public

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	en_translations "github.com/go-playground/validator/v10/translations/en"

	"intrajasa-merchant-api-gateway/core/structs"
	"intrajasa-merchant-api-gateway/core/validator"
	cv "intrajasa-merchant-api-gateway/core/validator/custom-validation"
	"intrajasa-merchant-api-gateway/core/api/vaonline"
)

var (
	dv_req structs.DisableVaRequest
	dv_res structs.DisableVaResponse
	gt_req structs.GetTokenRequest
	gt_res structs.GetTokenResponse
	gv_req structs.GenerateVaRequest
	gv_res structs.GenerateVaResponse
	gvp_req structs.GetVaPaymentStatusRequest
	gvp_res structs.GetVaPaymentStatusResponse
	uv_req structs.UpdateVaRequest
	uv_res structs.UpdateVaResponse
)

// APIRouter define router from here, you can add new api about your new services.
func APIRouter(router *gin.Engine) {
	// set validator
	validate := validator.InitValidator()
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	validate.RegisterValidation("int-length-1", cv.IntLen1)
	validate.RegisterValidation("int-length-5", cv.IntLen5)
	validate.RegisterValidation("int-length-10", cv.IntLen10)
	validate.RegisterValidation("uint64-length-18", cv.Uint64Len18)
	validate.RegisterValidation("float-decimal-val", cv.FloatDecimalVal)
	validate.RegisterValidation("date-format", cv.DateFormat)
	_ = en_translations.RegisterDefaultTranslations(validate, trans)
	// end set validator

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "welcome to apigateway, you can find you want here!!!", "userInfo": "Hello World!!!"})
	})

	// INTRAJASA DEFAULT ROUTE
	router.POST("/vaonline/rest/json/gettoken", func(c *gin.Context) {
		// using BindJson method to serialize body with struct
		if err := c.BindJSON(&gt_req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			gt_req = structs.GetTokenRequest{}
			return
		}
		if err := validate.Struct(gt_req); err != nil {
			errs := validator.ToErrResponse(err, trans)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errs,
			})
			gt_req = structs.GetTokenRequest{}
			return
		}
		
		gt_res = vaonline.GenerateToken(gt_req)
		c.JSON(http.StatusOK,&gt_res)
		gt_req = structs.GetTokenRequest{}
	})

	router.POST("/vaonline/rest/json/generateva", func(c *gin.Context) {
		// using BindJson method to serialize body with struct
		if err := c.BindJSON(&gv_req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			gv_req = structs.GenerateVaRequest{}
			return
		}
		if err := validate.Struct(gv_req); err != nil {
			errs := validator.ToErrResponse(err, trans)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errs,
			})
			gv_req = structs.GenerateVaRequest{}
			return
		}

		gv_res = vaonline.GenerateVa(gv_req)
		c.JSON(http.StatusOK,&gv_res)
		gv_req = structs.GenerateVaRequest{}
	})

	router.POST("/vaonline/rest/json/getstatus", func(c *gin.Context) {
		// using BindJson method to serialize body with struct
		if err := c.BindJSON(&gvp_req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			gvp_req = structs.GetVaPaymentStatusRequest{}
			return
		}
		if err := validate.Struct(gvp_req); err != nil {
			errs := validator.ToErrResponse(err, trans)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errs,
			})
			gvp_req = structs.GetVaPaymentStatusRequest{}
			return
		}

		gvp_res = vaonline.GetVaPaymentStatus(gvp_req)
		c.JSON(http.StatusOK,&gvp_res)
		gvp_req = structs.GetVaPaymentStatusRequest{}
	})

	router.POST("/vaonline/rest/json/disableVA", func(c *gin.Context) {
		// using BindJson method to serialize body with struct
		if err := c.BindJSON(&dv_req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			dv_req = structs.DisableVaRequest{}
			return
		}
		if err := validate.Struct(dv_req); err != nil {
			errs := validator.ToErrResponse(err, trans)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errs,
			})
			dv_req = structs.DisableVaRequest{}
			return
		}

		dv_res = vaonline.DisableVa(dv_req)
		c.JSON(http.StatusOK,&dv_res)
		dv_req = structs.DisableVaRequest{}
	})

	router.POST("/vaonline/rest/json/updateVA", func(c *gin.Context) {
		// using BindJson method to serialize body with struct
		if err := c.BindJSON(&uv_req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			uv_req = structs.UpdateVaRequest{}
			return
		}
		if err := validate.Struct(uv_req); err != nil {
			errs := validator.ToErrResponse(err, trans)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errs,
			})
			uv_req = structs.UpdateVaRequest{}
			return
		}

		uv_res = vaonline.UpdateVa(uv_req)
		c.JSON(http.StatusOK,&uv_res)
		uv_req = structs.UpdateVaRequest{}
	})
	// END INTRAJASA DEFAULT ROUTE
}
