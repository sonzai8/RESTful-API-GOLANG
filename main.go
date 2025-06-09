package main

import (
	"github.com/gin-gonic/gin"
	v1handler "main/internal/api/v1/handler"
)

func main() {
	r := gin.Default()
	v1User := v1handler.NewUserHandler()
	v1Product := v1handler.NewProductHandler()
	v1 := r.Group("api/v1")
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
			products.GET("/:id", v1Product.GetProductById)
			products.PATCH("/:id", v1Product.UpdateProduct)
			products.POST("/:id", v1Product.CreateProduct)
			products.DELETE("/:id", v1Product.DeleteProduct)
		}

	}

	r.Run(":8084")
}
