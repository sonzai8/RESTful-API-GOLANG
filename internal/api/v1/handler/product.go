package v1handler

import (
	"github.com/gin-gonic/gin"
	"regexp"
)

var slugRegex = regexp.MustCompile(`^[a-z0-9]+(?:[-,.][a-z0-9]+)*$`)

type ProductHandler struct{}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}

func (p *ProductHandler) GetProducts(ctx *gin.Context) {
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

	if !slugRegex.MatchString(slug) {
		ctx.JSON(200, gin.H{
			"message": "can not parse slug",
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
