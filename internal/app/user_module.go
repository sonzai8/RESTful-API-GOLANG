package app

import (
	"main/internal/handler"
	"main/internal/repository"
	"main/internal/routes"
	"main/internal/services"
)

type UserModule struct {
	routes routes.Route
}

func NewUserModule() *UserModule {
	uRepo := repository.NewInMemoryUserRepository()
	uService := services.NewUserService(uRepo)
	uHandler := handler.NewUserHandler(uService)
	usersRouter := routes.NewUserRoutes(uHandler)

	return &UserModule{
		routes: usersRouter,
	}
}

func (um *UserModule) Routes() routes.Route {

	return um.routes
}
