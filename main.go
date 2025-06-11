package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	v1handler "main/internal/api/v1/handler"
	"main/internal/middleware"
	"main/internal/utils"
	"os"
)

func main() {
	r := gin.Default()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	go middleware.CleanUpClients()
	TestEnv := os.Getenv("TEST_ENV")
	log.Println(TestEnv)
	if err := utils.RegisterValidators(); err != nil {
		panic(err)
	}
	v1User := v1handler.NewUserHandler()
	v1Product := v1handler.NewProductHandler()
	v1Category := v1handler.NewCategoryHandler()
	v1 := r.Group("api/v1")
	v1.Use(middleware.ApiKeyMiddleware(), middleware.RateLimitingMiddleware())
	{
		users := v1.Group("/users")
		{
			users.GET("/", v1User.GetUsers)
			users.GET("/:id", v1User.GetUserById)
			users.PATCH("/:id", v1User.UpdateUser)
			users.POST("/:id", v1User.CreateUser)
			users.DELETE("/:id", v1User.DeleteUser)
		}

		products := v1.Group("/products")
		{
			products.GET("/", v1Product.GetProducts)
			//products.GET("/:id", v1Product.GetProductById)
			products.GET("/:slug", v1Product.GetProductBySlug)
			products.PATCH("/:id", v1Product.UpdateProduct)
			products.POST("/", v1Product.CreateProduct)
			products.DELETE("/:id", v1Product.DeleteProduct)
		}

		categories := v1.Group("/categories")
		{
			categories.GET("/", v1Category.GetCategories)
			//products.GET("/:id", v1Product.GetProductById)
			categories.GET("/:category", v1Category.GetCategoriesByMap)
			categories.PATCH("/:id", v1Category.UpdateCategory)
			categories.POST("/:id", v1Category.CreateCategory)
			categories.DELETE("/:id", v1Category.DeleteCategory)
		}

	}

	r.Run(":8084")
}
