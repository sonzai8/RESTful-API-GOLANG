package app

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"main/internal/config"
	"main/internal/routers"
)

type Module interface {
	Router() routers.Route
}
type Application struct {
	config *config.Config
	router *gin.Engine
	module []Module
}

func NewApplication(cfg *config.Config) *Application {
	loadEnv()
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hello": "world",
		})
	})

	modules := []Module{
		New
	}

	return &Application{
		config: cfg,
		router: r,
	}
}

func (a *Application) Run() error {
	return a.router.Run(a.config.ServerAddress)
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
}
