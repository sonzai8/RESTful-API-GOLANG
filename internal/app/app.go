package app

import (
	"log"
	"main/internal/config"
	"main/internal/validation"

	"main/internal/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Module interface {
	Routes() routes.Route
}
type Application struct {
	config *config.Config
	router *gin.Engine
	module []Module
}

func NewApplication(cfg *config.Config) *Application {
	loadEnv()
	validation.InitValidator()
	r := gin.Default()

	modules := []Module{
		NewUserModule(),
	}
	routes.RegisterRouter(r, getModuleRoutes(modules)...)

	return &Application{
		config: cfg,
		router: r,
	}
}

func (a *Application) Run() error {
	return a.router.Run(a.config.ServerAddress)
}

func loadEnv() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Println("Error loading .env file")
	}
}

func getModuleRoutes(modules []Module) []routes.Route {

	var routes []routes.Route
	for _, m := range modules {
		routes = append(routes, m.Routes())
	}
	return routes
}
