package app

import (
	"main/internal/db"
	"main/internal/handler"
	"main/internal/repository"
	"main/internal/routes"
	"main/internal/services"
)

type UserModule struct {
	routes routes.Route
}

func NewUserModule() *UserModule {
	db.InitDB()
	uRepo := repository.NewSQLUserRepository(db.DB)
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
