package routers

import (
	"github.com/gin-gonic/gin"
	"main/internal/middleware"
)

type Route interface {
	Register(r *gin.RouterGroup)
}

func RegisterRouter(r *gin.Engine, routers ...Route) {
	r.Use(
		middleware.LoggerMiddleware(),
		middleware.ApiKeyMiddleware(),
		middleware.RateLimitingMiddleware(),
	)

	api := r.Group("/api/v1")

	for _, router := range routers {
		router.Register(api)
	}
}
