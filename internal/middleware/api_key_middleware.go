package middleware

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func ApiKeyMiddleware() gin.HandlerFunc {

	expectedKey := os.Getenv("X_API_KEY")
	log.Println("X_API_KEY is", expectedKey)
	if expectedKey == "" {
		expectedKey = "secret-key"
	}

	return func(ctx *gin.Context) {
		apiKey := ctx.GetHeader("X-API-KEY")
		log.Println("API KEY: ", apiKey)
		if apiKey == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}
		if apiKey != expectedKey {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}
		ctx.Next()

	}

}
