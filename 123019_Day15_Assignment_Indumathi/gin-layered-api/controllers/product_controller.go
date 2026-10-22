package controllers

import (
	"net/http"
	"strconv"
	"gin-layered-api/models"
	"gin-layered-api/services"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productService services.ProductService
}

func NewProductController(productService services.ProductService) *ProductController {
	return &ProductController{productService: productService}
}

func (c *ProductController) GetProducts(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "2"))

	response := c.productService.GetProducts(page, limit)
	ctx.JSON(http.StatusOK, response)
}

func (c *ProductController) CreateProduct(ctx *gin.Context) {
	var req models.CreateProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := c.productService.CreateProduct(req)
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Product created successfully",
		"product": product,
	})
}
