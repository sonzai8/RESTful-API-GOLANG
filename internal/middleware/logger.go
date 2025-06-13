package middleware

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

func LoggerMiddleware() gin.HandlerFunc {

	logPath := "logs/http.log"
	if err := os.MkdirAll(filepath.Dir(logPath), os.ModePerm); err != nil {
		panic(err)
	}
	logFile, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	logger := zerolog.New(logFile).With().Timestamp().Logger()

	return func(ctx *gin.Context) {
		start := time.Now().Second()

		bodyBytes, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			logger.Error().Err(err).Msg("Failed to read request body")
		}

		ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		fmt.Printf("%+v", string(bodyBytes))

		fmt.Printf("-----------------")

		fmt.Printf("%+v", ctx.Request.Body)

		ctx.Next()

		end := time.Now().Second()
		duration := end - start
		statusCode := ctx.Writer.Status()
		logEvent := logger.Info()

		if statusCode >= 400 {
			logEvent = logger.Warn()
		} else if statusCode >= 500 {
			logEvent = logger.Error()
		}
		log.Println(statusCode, ctx.ClientIP())

		logEvent.Str("method", ctx.Request.Method).
			Str("path", ctx.Request.URL.Path).
			Str("ip", ctx.ClientIP()).
			Str("referer", ctx.Request.Referer()).
			Int("duration", duration).Msg("Http Request Log")

		log.Println("da xong phan ghi log ra file")
	}
}
