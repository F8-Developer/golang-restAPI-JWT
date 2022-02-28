package Public

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	en_translations "github.com/go-playground/validator/v10/translations/en"

	"intrajasa-merchant-api-gateway/Core/Structs"
	"intrajasa-merchant-api-gateway/Core/Validator"
	cv "intrajasa-merchant-api-gateway/Core/Validator/CustomValidation"
	"intrajasa-merchant-api-gateway/Core/Api/VaOnline"
	"intrajasa-merchant-api-gateway/Core/Utils"
)

var (
	dv_req Structs.DisableVaRequest
	dv_res Structs.DisableVaResponse
	gt_req Structs.GetTokenRequest
	gt_res Structs.GetTokenResponse
	gv_req Structs.GenerateVaRequest
	gv_res Structs.GenerateVaResponse
	gvp_req Structs.GetVaPaymentStatusRequest
	gvp_res Structs.GetVaPaymentStatusResponse
	uv_req Structs.UpdateVaRequest
	uv_res Structs.UpdateVaResponse
)

// APIRouter define router from here, you can add new api about your new services.
func APIRouter(router *gin.Engine) {
	// set validator
	validate := Validator.InitValidator()
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
			gt_req = Structs.GetTokenRequest{}
			return
		}
		if err := validate.Struct(gt_req); err != nil {
			errs := Validator.ToErrResponse(err, trans)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errs,
			})
			gt_req = Structs.GetTokenRequest{}
			return
		}
		
		gt_res = VaOnline.GenerateToken(gt_req)
		c.JSON(gt_res.ResponseCode,&gt_res)
		gt_req = Structs.GetTokenRequest{}
	})

	router.POST("/vaonline/rest/json/generateva", func(c *gin.Context) {
		// validate token
		request := c.Request.URL.Query()
		if !Utils.ValidateToken(request.Get("t")) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Not Authorized!",
			})
			gv_req = Structs.GenerateVaRequest{}
			return
		}

		// using BindJson method to serialize body with struct
		if err := c.BindJSON(&gv_req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			gv_req = Structs.GenerateVaRequest{}
			return
		}
		if err := validate.Struct(gv_req); err != nil {
			errs := Validator.ToErrResponse(err, trans)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errs,
			})
			gv_req = Structs.GenerateVaRequest{}
			return
		}

		gv_res = VaOnline.GenerateVa(gv_req)
		c.JSON(http.StatusOK,&gv_res)
		gv_req = Structs.GenerateVaRequest{}
	})

	router.POST("/vaonline/rest/json/getstatus", func(c *gin.Context) {
		// using BindJson method to serialize body with struct
		if err := c.BindJSON(&gvp_req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			gvp_req = Structs.GetVaPaymentStatusRequest{}
			return
		}
		if err := validate.Struct(gvp_req); err != nil {
			errs := Validator.ToErrResponse(err, trans)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errs,
			})
			gvp_req = Structs.GetVaPaymentStatusRequest{}
			return
		}

		gvp_res = VaOnline.GetVaPaymentStatus(gvp_req)
		c.JSON(http.StatusOK,&gvp_res)
		gvp_req = Structs.GetVaPaymentStatusRequest{}
	})

	router.POST("/vaonline/rest/json/disableVA", func(c *gin.Context) {
		// using BindJson method to serialize body with struct
		if err := c.BindJSON(&dv_req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			dv_req = Structs.DisableVaRequest{}
			return
		}
		if err := validate.Struct(dv_req); err != nil {
			errs := Validator.ToErrResponse(err, trans)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errs,
			})
			dv_req = Structs.DisableVaRequest{}
			return
		}

		dv_res = VaOnline.DisableVa(dv_req)
		c.JSON(http.StatusOK,&dv_res)
		dv_req = Structs.DisableVaRequest{}
	})

	router.POST("/vaonline/rest/json/updateVA", func(c *gin.Context) {
		// using BindJson method to serialize body with struct
		if err := c.BindJSON(&uv_req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			uv_req = Structs.UpdateVaRequest{}
			return
		}
		if err := validate.Struct(uv_req); err != nil {
			errs := Validator.ToErrResponse(err, trans)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errs,
			})
			uv_req = Structs.UpdateVaRequest{}
			return
		}

		uv_res = VaOnline.UpdateVa(uv_req)
		c.JSON(http.StatusOK,&uv_res)
		uv_req = Structs.UpdateVaRequest{}
	})
	// END INTRAJASA DEFAULT ROUTE
}
