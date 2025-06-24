package app

import (
	"fmt"
	"log"
	"main/internal/config"
	"main/internal/validation"

	"main/internal/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	db := connectToPostgres()
	fmt.Printf("%v", db)
	if err := validation.InitValidator(); err != nil {
		log.Fatalf("Failed to initialize validator: %v", err)
	}

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

func connectToPostgres() *gorm.DB {

	// https://github.com/jackc/pgx
	dsn := config.NewConfig().DbConNStr
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("connect to db failed")
	}

	fmt.Printf("db:", db)
	return db
}

func getModuleRoutes(modules []Module) []routes.Route {

	var routes []routes.Route
	for _, m := range modules {
		routes = append(routes, m.Routes())
	}
	return routes
}
