package v1handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

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
	idStr := ctx.Param("id")
	_, err := uuid.Parse(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Id must be a uuid",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "get user by id " + idStr,
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
