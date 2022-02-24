package public

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	en_translations "github.com/go-playground/validator/v10/translations/en"

	"intrajasa-merchant-api-gateway/core/structs"
	"intrajasa-merchant-api-gateway/core/validator"
	"intrajasa-merchant-api-gateway/core/api/users"
	"intrajasa-merchant-api-gateway/core/utils/consts"
)

var (
	gtr structs.GetTokenRequest
	gvr structs.GenerateVaRequest
	uis users.UserInfoService
	flag int
)

// APIRouter define router from here, you can add new api about your new services.
func APIRouter(router *gin.Engine) {
	// set validator
	validate := validator.InitValidator()
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	validate.RegisterValidation("is-awesome", validator.ValidateMyVal)
	_ = en_translations.RegisterDefaultTranslations(validate, trans)
	// end set validator

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "welcome to apigateway, you can find you want here!!!", "userInfo": "Hello World!!!"})
	})

	// INTRAJASA DEFAULT ROUTE
	router.POST("/vaonline/rest/json/gettoken", func(c *gin.Context) {
		// using BindJson method to serialize body with struct
		if err := c.BindJSON(&gtr); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			gtr = structs.GetTokenRequest{}
			return
		}
		if err := validate.Struct(gtr); err != nil {
			errs := validator.TranslateError(err, trans)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errs,
			})
			gtr = structs.GetTokenRequest{}
			return
		}
		c.JSON(http.StatusOK,&gtr)
		gtr = structs.GetTokenRequest{}
	})

	router.POST("/vaonline/rest/json/generateva", func(c *gin.Context) {
		// using BindJson method to serialize body with struct
		if err := c.BindJSON(&gvr); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			gvr = structs.GenerateVaRequest{}
			return
		}
		if err := validate.Struct(gvr); err != nil {
			errs := validator.ToErrResponse(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errs,
			})
			gvr = structs.GenerateVaRequest{}
			return
		}
		c.JSON(http.StatusOK,&gvr)
		gvr = structs.GenerateVaRequest{}
	})

	router.POST("/vaonline/rest/json/getstatus", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"merchantRefCode": "JS008sKs", 
			"merchantId": "001", 
			"totalAmount": "200000.0", 
			"vaNumber": "8228006200100634 ", 
			"paymentStatus": "1", 
			"responseCode": "200", 
			"responseMsg": "Success"})
	})

	router.POST("/vaonline/rest/json/disableVA", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"responseCode": "200", "responseMsg": "Success"})
	})

	router.POST("/vaonline/rest/json/updateVA", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"responseCode": "200", "responseMsg": "Success"})
	})
	// END INTRAJASA DEFAULT ROUTE

	router.POST("/api/users/login", func(c *gin.Context) {
		c.BindJSON(&uis)
		/*
			// TO-DO: cache user login session.
			session := sessions.Default(c)
			if session.Get(uis.USERNAME) == nil {
				flag = uis.Login()
				session.Set(uis.USERNAME, uis.USERNAME)
				session.Save()
				log.Println("Try login and save session in session store.")
			} else {
				flag = consts.SUCCESS
				log.Println("Have a session in session store.")
			}
		*/
		flag = uis.Login()
		switch flag {
		case consts.SUCCESS:
			c.JSON(http.StatusOK, gin.H{"code": consts.SUCCESS, "Message": "Login Successful", "tooken": ""})
		case consts.NOACCOUNT:
			c.JSON(http.StatusOK, gin.H{"code": consts.NOACCOUNT, "Message": "Not found your account"})
		case consts.SYSERROR:
			c.JSON(http.StatusOK, gin.H{"code": consts.SYSERROR, "Message": "System error!!!"})
		case consts.WRONGPASSWD:
			c.JSON(http.StatusOK, gin.H{"code": consts.WRONGPASSWD, "Message": "Wrong password..."})
		default:
			c.JSON(http.StatusOK, gin.H{"code": consts.SYSERROR, "Message": "Unknow error.."})
		}
	})

	router.POST("/api/users/register", func(c *gin.Context) {
		c.BindJSON(&uis)
		if err := uis.Register(); err == nil {
			c.JSON(http.StatusOK, gin.H{"statusCode": http.StatusOK, "message": "Welcome " + uis.USERNAME + ",you have login successful!"})
		} else {
			c.JSON(http.StatusExpectationFailed, gin.H{"errorMessage": "Rigster failed "})
		}
	})
}
