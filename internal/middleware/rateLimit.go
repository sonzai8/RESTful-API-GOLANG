package middleware

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"net/http"
	"sync"
	"time"
)

type Client struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

var (
	mu      sync.Mutex
	clients = make(map[string]Client)
)

func getClientIP(ctx *gin.Context) string {

	ip := ctx.ClientIP()
	if ip == "" {
		ip = ctx.Request.Header.Get("X-Forwarded-For")
	}
	return ip
}

func getRateLimiter(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()
	client, exists := clients[ip]
	if !exists {
		limiter := rate.NewLimiter(5, 10)
		newClient := &Client{limiter, time.Now()}
		clients[ip] = *newClient
		return limiter
	}

	client.lastSeen = time.Now()
	return client.limiter

}

func CleanUpClients() {
	for {
		time.Sleep(1 * time.Minute)
		mu.Lock()

		for ip, client := range clients {
			if time.Since(client.lastSeen) > 3*time.Minute {
				delete(clients, ip)
			}
		}
		mu.Unlock()
	}
}

// ab -n 20 -c 1 localhost:8084/api/v1/categories
// wrk -t4 -c100 -d10s localhost:8084/api/v1/categories
func RateLimitingMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ip := getClientIP(ctx)
		limiter := getRateLimiter(ip)
		if !limiter.Allow() {
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"message": "Too many requests",
			})
			return
		}
		ctx.Next()
	}
}
