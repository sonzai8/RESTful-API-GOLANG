package app

import "main/internal/routers"

type UserModule struct {
	routers routers.Route
}

func NewUserModule() *UserModule {

	return &UserModule{}
}
