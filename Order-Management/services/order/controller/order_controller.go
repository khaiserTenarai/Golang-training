package controller

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"order-management/services/order/dto"
	"order-management/services/order/service"
)

type OrderController struct {
	service service.OrderService
}

func NewOrderController(s service.OrderService) *OrderController {
	return &OrderController{
		service: s,
	}
}

// Create godoc
// @Summary Create an order
// @Description Creates an order after validating the customer, product, and quantity.
// @Tags Orders
// @Accept json
// @Produce json
// @Param order body dto.CreateOrderRequest true "Order details"
// @Success 201 {object} map[string]interface{} "Order created successfully"
// @Failure 400 {object} map[string]string "Invalid request or order creation failed"
// @Router /orders [post]
func (c *OrderController) Create(ctx *gin.Context) {
	var req dto.CreateOrderRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	order, err := c.service.Create(
		ctx.Request.Context(),
		req.CustomerID,
		req.ProductID,
		req.Quantity,
	)
	if err != nil {
		ctx.Set("error", err.Error())
		status := http.StatusBadRequest
		message := strings.ToLower(err.Error())
		switch {
		case errors.Is(err, context.DeadlineExceeded):
			status = http.StatusGatewayTimeout
		case strings.Contains(message, "insufficient stock"), strings.Contains(message, "http 409"):
			status = http.StatusConflict
		case strings.Contains(message, "http 404"):
			status = http.StatusNotFound
		case strings.Contains(message, "create order failed"), strings.Contains(message, "create order:"):
			status = http.StatusInternalServerError
		case strings.Contains(message, "product service"), strings.Contains(message, "circuit breaker"):
			status = http.StatusBadGateway
		}
		ctx.JSON(status, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, order)
}

// GetAll godoc
// @Summary List all orders
// @Description Returns all orders.
// @Tags Orders
// @Produce json
// @Success 200 {array} map[string]interface{} "List of orders"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /orders [get]
func (c *OrderController) GetAll(ctx *gin.Context) {
	orders, err := c.service.GetAll(ctx.Request.Context())
	if err != nil {
		ctx.Set("error", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, orders)
}

// GetByID godoc
// @Summary Get an order by ID
// @Description Returns one order using its ID.
// @Tags Orders
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} map[string]interface{} "Order details"
// @Failure 400 {object} map[string]string "Invalid order ID"
// @Failure 404 {object} map[string]string "Order not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /orders/{id} [get]
func (c *OrderController) GetByID(ctx *gin.Context) {
	id, err := parseID(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid order ID",
		})
		return
	}

	order, err := c.service.GetByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.Set("error", err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "order not found",
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, order)
}

// UpdateStatus godoc
// @Summary Update order status
// @Description Updates an order status, for example to CONFIRMED or CANCELLED.
// @Tags Orders
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Param status body dto.UpdateOrderStatusRequest true "New order status"
// @Success 200 {object} map[string]interface{} "Order status updated successfully"
// @Failure 400 {object} map[string]string "Invalid request or status"
// @Failure 404 {object} map[string]string "Order not found"
// @Router /orders/{id}/status [patch]
func (c *OrderController) UpdateStatus(ctx *gin.Context) {
	id, err := parseID(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid order ID",
		})
		return
	}

	var req dto.UpdateOrderStatusRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	order, err := c.service.UpdateStatus(
		ctx.Request.Context(),
		id,
		req.Status,
	)
	if err != nil {
		ctx.Set("error", err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
			return
		}
		status := http.StatusBadRequest
		if errors.Is(err, context.DeadlineExceeded) {
			status = http.StatusGatewayTimeout
		}
		if strings.Contains(strings.ToLower(err.Error()), "product service") || strings.Contains(strings.ToLower(err.Error()), "circuit breaker") {
			status = http.StatusBadGateway
		}
		ctx.JSON(status, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, order)
}

func parseID(value string) (uint, error) {
	id, err := strconv.ParseUint(value, 10, 64)
	if err != nil || id == 0 {
		return 0, errors.New("invalid ID")
	}

	return uint(id), nil
}
