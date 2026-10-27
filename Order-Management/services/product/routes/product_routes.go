package routes

import (
	"crypto/subtle"
	"errors"
	"net/http"
	"os"
	"strconv"

	"gorm.io/gorm"

	"order-management/services/product/controller"
	productresolver "order-management/services/product/graphql/resolver"
	"order-management/services/product/service"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

// Register registers the Product Service HTTP and GraphQL routes.
func Register(
	router *gin.Engine,
	productController *controller.ProductController,
	productService service.ProductService,
) error {
	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "healthy",
			"service": "product-service",
		})
	})

	products := router.Group("/products")
	{
		products.POST("", productController.Create)
		products.GET("", productController.GetAll)
		products.GET("/:id", productController.GetByID)
		products.PUT("/:id", productController.Update)
		products.DELETE("/:id", productController.Delete)
	}

	// Internal stock APIs are service-to-service only; never expose them through the Gateway.
	internal := router.Group("/internal/products")
	internal.Use(func(ctx *gin.Context) {
		expected := os.Getenv("INTERNAL_SERVICE_TOKEN")
		provided := ctx.GetHeader("X-Internal-Token")
		if expected == "" || subtle.ConstantTimeCompare([]byte(expected), []byte(provided)) != 1 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized service request"})
			return
		}
		ctx.Next()
	})
	internal.POST("/:id/reserve", func(ctx *gin.Context) {
		id, err := parseID(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID"})
			return
		}
		var req struct {
			OrderReference string `json:"orderReference" binding:"required"`
			Quantity       int    `json:"quantity" binding:"required"`
		}
		if err := ctx.ShouldBindJSON(&req); err != nil || req.OrderReference == "" || req.Quantity <= 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "orderReference and positive quantity are required"})
			return
		}
		product, err := productService.ReserveStock(ctx.Request.Context(), id, req.OrderReference, req.Quantity)
		if err != nil {
			ctx.Set("error", err.Error())
			status := http.StatusConflict
			if errors.Is(err, gorm.ErrRecordNotFound) {
				status = http.StatusNotFound
			}
			ctx.JSON(status, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, product)
	})
	internal.POST("/:id/release", func(ctx *gin.Context) {
		id, err := parseID(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID"})
			return
		}
		var req struct {
			OrderReference string `json:"orderReference" binding:"required"`
			Quantity       int    `json:"quantity" binding:"required"`
		}
		if err := ctx.ShouldBindJSON(&req); err != nil || req.OrderReference == "" || req.Quantity <= 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "orderReference and positive quantity are required"})
			return
		}
		if err := productService.ReleaseStock(ctx.Request.Context(), id, req.OrderReference, req.Quantity); err != nil {
			ctx.Set("error", err.Error())
			status := http.StatusConflict
			if errors.Is(err, gorm.ErrRecordNotFound) {
				status = http.StatusNotFound
			}
			ctx.JSON(status, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"status": "released", "productId": id, "quantity": req.Quantity, "orderReference": req.OrderReference})
	})

	graphQLSchema, err := productresolver.New(productService)
	if err != nil {
		return err
	}

	router.POST("/graphql", func(ctx *gin.Context) {
		var request struct {
			Query         string                 `json:"query" binding:"required"`
			OperationName string                 `json:"operationName"`
			Variables     map[string]interface{} `json:"variables"`
		}

		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"errors": []gin.H{{"message": "Invalid GraphQL request"}},
			})
			return
		}

		result := graphql.Do(graphql.Params{
			Schema:         graphQLSchema,
			RequestString:  request.Query,
			OperationName:  request.OperationName,
			VariableValues: request.Variables,
			Context:        ctx.Request.Context(),
		})

		ctx.JSON(http.StatusOK, result)
	})

	return nil
}

func parseID(value string) (uint, error) {
	id, err := strconv.ParseUint(value, 10, 64)
	if err != nil || id == 0 {
		return 0, errors.New("invalid ID")
	}
	return uint(id), nil
}
