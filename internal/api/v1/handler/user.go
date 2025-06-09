package v1handler

import "github.com/gin-gonic/gin"

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (u *UserHandler) GetUsers(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "get list of user",
	})
}
func (u *UserHandler) GetUserById(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "get user by id",
	})
}

func (u *UserHandler) UpdateUser(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "update user",
	})
}

func (u *UserHandler) CreateUser(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Create user",
	})
}

func (u *UserHandler) DeleteUser(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "DELETE user",
	})
}
