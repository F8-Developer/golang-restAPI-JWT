package Router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"golang-restAPI-JWT/Core/Router/Public"
	"golang-restAPI-JWT/Middleware"
	"golang-restAPI-JWT/Config"
)

// Varible define to here
var (
	LisAddr string
)

func init() {
	LisAddr = Config.GoDotEnvVariable("APP_ADDRESS")
	if LisAddr == "" {
		LisAddr = "0.0.0.0:8080"
	}
}

// Start application by load self-define router.
func Start(env string) {
	// enable debug/release mode
	switch env {
	case "development":
		gin.SetMode(gin.DebugMode)
	default:
		gin.SetMode(gin.ReleaseMode)
		fmt.Printf("Start prod mode...\nServer listen on: %v", LisAddr)
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(Middleware.CORSMiddleware())
	router.Use(Middleware.LoggerApp())
	//No Permission Validation
	Public.APIRouter(router)

	router.Run(LisAddr)
}
