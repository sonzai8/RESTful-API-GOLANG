package v1handler

import (
	"github.com/gin-gonic/gin"
	"main/internal/utils"
	"regexp"
)

var slugRegex = regexp.MustCompile(`^[a-z0-9]+(?:[-,.][a-z0-9]+)*$`)

type ProductHandler struct{}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}

func (p *ProductHandler) GetProducts(ctx *gin.Context) {
	search := ctx.Query("search")
	if err := utils.ValidationRequired("search", search); err != nil {
		ctx.JSON(200, gin.H{
			"message": err.Error(),
		})
		return
	}
	if err := utils.ValidationStringLength("search", search, 50, 3); err != nil {
		ctx.JSON(200, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "get list of Product",
	})
}

func (p *ProductHandler) GetProductById(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "get Product by id",
	})
}

func (p *ProductHandler) GetProductBySlug(ctx *gin.Context) {
	slug := ctx.Param("slug")

	if err := utils.ValidationRegex("slug", slug, slugRegex); err != nil {
		ctx.JSON(200, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "get Product by slug" + slug,
	})

}

func (p *ProductHandler) UpdateProduct(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "update Product",
	})
}

func (p *ProductHandler) CreateProduct(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Create Product",
	})
}

func (p *ProductHandler) DeleteProduct(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Delete Product",
	})
}
