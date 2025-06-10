package v1handler

import (
	"github.com/gin-gonic/gin"
)

type GetUsersByParam struct {
	ID int `uri:"id" biding:"gt=0"`
}

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
	var params GetUsersByParam
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(200, gin.H{
			"message": err,
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "get user by id ",
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
