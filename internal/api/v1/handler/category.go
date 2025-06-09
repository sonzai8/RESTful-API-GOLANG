package v1handler

import (
	"github.com/gin-gonic/gin"
)

var validCategories = map[string]bool{
	"php":    true,
	"python": true,
	"golang": true,
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
	category := ctx.Param("category")
	if !validCategories[category] {
		ctx.JSON(200, gin.H{
			"message": "can not parse category",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "get Category " + category,
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
