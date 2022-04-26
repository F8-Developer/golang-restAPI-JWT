package Middleware

import (
	"fmt"

	jwtAuth "golang-restAPI-JWT/Auth"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc{
	return func(context *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := context.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA):]
		fmt.Println(tokenString)
		if tokenString == "" {
			context.JSON(401, gin.H{"error": "Not Authorized"})
			context.Abort()
			return
		}
		err:= jwtAuth.ValidateToken(tokenString)
		if err != nil {
			context.JSON(401, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
		context.Next()
	}
}