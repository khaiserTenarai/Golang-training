package controller

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"order-management/services/product/dto"
	"order-management/services/product/model"
	"order-management/services/product/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ProductController handles product HTTP requests.
type ProductController struct {
	service service.ProductService
}

// NewProductController creates a ProductController.
func NewProductController(
	productService service.ProductService,
) *ProductController {
	return &ProductController{
		service: productService,
	}
}

// Create creates a product.
//
// @Summary Create a product
// @Tags Products
// @Accept json
// @Produce json
// @Param product body dto.CreateProductRequest true "Product details"
// @Success 201 {object} model.Product
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /products [post]
func (c *ProductController) Create(ctx *gin.Context) {
	var request dto.CreateProductRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid product request: " + err.Error(),
		})
		return
	}

	if err := request.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := &model.Product{
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		Stock:       request.Stock,
	}

	created, err := c.service.Create(ctx.Request.Context(), product)
	if err != nil {
		c.respondServiceError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, created)
}

// GetAll returns all products.
//
// @Summary List products
// @Tags Products
// @Produce json
// @Success 200 {array} model.Product
// @Failure 500 {object} map[string]string
// @Router /products [get]
func (c *ProductController) GetAll(ctx *gin.Context) {
	products, err := c.service.GetAll(ctx.Request.Context())
	if err != nil {
		c.respondServiceError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, products)
}

// GetByID returns a product by ID.
//
// @Summary Get a product by ID
// @Tags Products
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} model.Product
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /products/{id} [get]
func (c *ProductController) GetByID(ctx *gin.Context) {
	id, ok := parseProductID(ctx)
	if !ok {
		return
	}

	product, err := c.service.GetByID(ctx.Request.Context(), id)
	if err != nil {
		c.respondServiceError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, product)
}

// Update updates an existing product.
//
// @Summary Update a product
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body dto.UpdateProductRequest true "Updated product details"
// @Success 200 {object} model.Product
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /products/{id} [put]
func (c *ProductController) Update(ctx *gin.Context) {
	id, ok := parseProductID(ctx)
	if !ok {
		return
	}

	var request dto.UpdateProductRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid product request: " + err.Error(),
		})
		return
	}

	if err := request.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := &model.Product{
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		Stock:       request.Stock,
	}

	updated, err := c.service.Update(
		ctx.Request.Context(),
		id,
		product,
	)
	if err != nil {
		c.respondServiceError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, updated)
}

// Delete removes a product.
//
// @Summary Delete a product
// @Tags Products
// @Produce json
// @Param id path int true "Product ID"
// @Success 204
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /products/{id} [delete]
func (c *ProductController) Delete(ctx *gin.Context) {
	id, ok := parseProductID(ctx)
	if !ok {
		return
	}

	if err := c.service.Delete(ctx.Request.Context(), id); err != nil {
		c.respondServiceError(ctx, err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

// parseProductID validates the ID path parameter.
func parseProductID(ctx *gin.Context) (uint, bool) {
	value, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil || value == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "product ID must be a positive integer",
		})
		return 0, false
	}

	return uint(value), true
}

// respondServiceError converts service errors into HTTP responses.
func (c *ProductController) respondServiceError(
	ctx *gin.Context,
	err error,
) {
	ctx.Set("error", err.Error())
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "product not found",
		})
		return
	}

	if errors.Is(err, context.Canceled) {
		ctx.JSON(http.StatusRequestTimeout, gin.H{
			"error": "request was cancelled",
		})
		return
	}

	if errors.Is(err, context.DeadlineExceeded) {
		ctx.JSON(http.StatusGatewayTimeout, gin.H{
			"error": "request timed out",
		})
		return
	}

	ctx.JSON(http.StatusInternalServerError, gin.H{
		"error": err.Error(),
	})
}
