package Router

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"merchant-api-gateway/Core/Router/Public"
	"merchant-api-gateway/Middleware"
)

// Varible define to here
var (
	LisAddr string
)

func init() {
	LisAddr = os.Getenv("ADDRESS")
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
