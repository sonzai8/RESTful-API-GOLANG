package handler

import (
	"main/internal/services"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {

}

func (h *UserHandler) GetUserByUUID(c *gin.Context) {

}

func (h *UserHandler) CreateUser(c *gin.Context) {

}

func (h *UserHandler) UpdateUser(c *gin.Context) {

}

func (h *UserHandler) DeleteUser(c *gin.Context) {

}
