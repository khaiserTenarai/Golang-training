package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	graphql "github.com/graphql-go/graphql"

	"order-management/services/order/controller"
	orderresolver "order-management/services/order/graphql/resolver"
	"order-management/services/order/service"
)

type graphQLRequest struct {
	Query         string                 `json:"query"`
	Variables     map[string]interface{} `json:"variables"`
	OperationName string                 `json:"operationName"`
}

func Register(
	router *gin.Engine,
	orderController *controller.OrderController,
	orderService service.OrderService,
) {
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"service": "order-service",
			"status":  "healthy",
		})
	})

	router.POST("/orders", orderController.Create)
	router.GET("/orders", orderController.GetAll)
	router.GET("/orders/:id", orderController.GetByID)
	router.PATCH("/orders/:id/status", orderController.UpdateStatus)

	// Build the schema once at startup instead of rebuilding it for every request.
	graphQLSchema, schemaErr := orderresolver.New(orderService)
	router.POST("/graphql", func(c *gin.Context) {
		var request graphQLRequest

		if err := json.NewDecoder(c.Request.Body).Decode(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": []gin.H{
					{"message": "Invalid JSON request body"},
				},
			})
			return
		}

		if request.Query == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": []gin.H{
					{"message": "GraphQL query is required"},
				},
			})
			return
		}

		if schemaErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"errors": []gin.H{{"message": "Failed to initialize GraphQL schema"}}})
			return
		}

		result := graphql.Do(graphql.Params{
			Schema:         graphQLSchema,
			RequestString:  request.Query,
			VariableValues: request.Variables,
			OperationName:  request.OperationName,
			Context:        c.Request.Context(),
		})

		status := http.StatusOK
		if len(result.Errors) > 0 {
			status = http.StatusBadRequest
		}

		c.JSON(status, result)
	})
}
