package v1handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"main/internal/utils"
	"net/http"
	"regexp"
)

var slugRegex = regexp.MustCompile(`^[a-z0-9]+(?:[-,.][a-z0-9]+)*$`)

type GetProductsParams struct {
	Search  string `form:"search" binding:"required,search"`
	Date    string `form:"date" binding:"required,datetime=2006-01-02"`
	Display *bool  `json:"display" binding:"omitempty"`
}
type GetProductsBySlugParam struct {
	Slug string `uri:"slug" binding:"slug,min=5,max=100"`
	Page int    `uri:"test" binding:"gt=3,required"`
}

type CreateProductParams struct {
	Name             string                 `json:"name" binding:"required,min=5,max=100"`
	Price            int                    `json:"price" binding:"required,min_int=1000,max_int=10000000"`
	Display          *bool                  `json:"display" binding:"omitempty"`
	Image            ProductImage           `json:"image" binding:"required"`
	Tags             []string               `json:"tags" binding:"required,gt=3,lt=10"`
	ProductAttribute []ProductAttribute     `json:"product_attribute" binding:"required,gt=0,dive"`
	ProductInfo      map[string]ProductInfo `json:"product_info" binding:"required,gt=0,dive"`
	ProductMetaData  map[string]any         `json:"product_metadata" binding:"omitempty,gt=0,dive"`
}

type ProductInfo struct {
	InfoKey   string `json:"info_key" binding:"required"`
	InfoValue string `json:"info_value" binding:"required"`
}
type ProductAttribute struct {
	AttributeName  string `json:"attribute_name" binding:"required"`
	AttributeValue string `json:"attribute_value" binding:"required"`
}
type ProductImage struct {
	ImageName string `json:"image_name" binding:"required,file_ext=jpg"`
	ImageURL  string `json:"image_url" binding:"required"`
}
type ProductHandler struct{}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}

func (p *ProductHandler) GetProducts(ctx *gin.Context) {
	var params GetProductsParams
	if err := ctx.ShouldBindQuery(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
		return
	}
	if params.Display == nil {
		defaultDisplay := true
		params.Display = &defaultDisplay
	}
	limit := ctx.DefaultQuery("limit", "10")
	if limit == "" {
		limit = "10"
	}
	ctx.JSON(200, gin.H{
		"message": "get list of Product",
		"Display": params.Display,
	})
}

func (p *ProductHandler) GetProductById(ctx *gin.Context) {
	id := ctx.Param("id")
	value, err := utils.ValidationPositiveInt("id", id)
	if err != nil {
		ctx.JSON(200, gin.H{
			"message": err.Error(),
		})
		return
	}

	fmt.Println(value)
	ctx.JSON(200, gin.H{
		"message": "get Product by id",
	})
}

func (p *ProductHandler) GetProductBySlug(ctx *gin.Context) {

	var params GetProductsBySlugParam

	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
		return
	}
	ctx.JSON(200, gin.H{
		"message": "get Product by slug",
	})

}

func (p *ProductHandler) UpdateProduct(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "update Product",
	})
}

func (p *ProductHandler) CreateProduct(ctx *gin.Context) {
	var params CreateProductParams
	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
		return
	}

	for key := range params.ProductInfo {
		log.Println(key)
		if _, err := uuid.Parse(key); err != nil {
			ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
			return
		}

	}

	ctx.JSON(200, gin.H{
		"message": "Create Product",
	})
}

func (p *ProductHandler) DeleteProduct(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Delete Product",
	})
}
