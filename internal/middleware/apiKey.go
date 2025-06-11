package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func ApiKeyMiddleware() gin.HandlerFunc {

	expectedKey := os.Getenv("API_KEY")
	log.Println("API_KEY is", expectedKey)
	if expectedKey == "" {
		expectedKey = "secret-key"
	}

	return func(ctx *gin.Context) {
		apiKey := ctx.GetHeader("X-Api-Key")
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
