package v1handler

import (
	"github.com/gin-gonic/gin"
	"main/internal/utils"
	"net/http"
)

type GetCategoriesByMapParams struct {
	Category string `uri:"category" binding:"oneof=php python golang"`
}
type CategoryHandler struct{}

func NewCategoryHandler() *CategoryHandler {
	return &CategoryHandler{}
}

func (p *CategoryHandler) GetCategories(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "get list of Category",
	})
}

func (p *CategoryHandler) GetCategoriesByMap(ctx *gin.Context) {
	var params GetCategoriesByMapParams
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
		return
	}
	ctx.JSON(200, gin.H{
		"message": "get Category ",
	})
}

func (p *CategoryHandler) UpdateCategory(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "update Category",
	})
}

func (p *CategoryHandler) CreateCategory(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Create Category",
	})
}

func (p *CategoryHandler) DeleteCategory(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Delete Category",
	})
}
