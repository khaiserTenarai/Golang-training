package router

import (
	"context"
	"net/http"

	"employee-platform/controllers"
	"employee-platform/middleware"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

type GQLRequest struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}

func SetupRouter(empCtrl *controllers.EmployeeController, gqlSchema graphql.Schema, jwtSecret string) *gin.Engine {
	r := gin.Default()

	r.Use(middleware.AuthMiddleware(jwtSecret))

	// REST Routes
	api := r.Group("/api/v1")
	{
		api.GET("/employees", empCtrl.GetEmployees)
		api.POST("/employees", middleware.RequireRole("ADMIN"), empCtrl.CreateEmployee)
	}

	// GraphQL Endpoint
	r.POST("/graphql", func(c *gin.Context) {
		var req GQLRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Pass context with authorization to GraphQL resolvers
		ctx := context.WithValue(c.Request.Context(), "userCtx", c.MustGet("userCtx"))

		result := graphql.Do(graphql.Params{
			Schema:         gqlSchema,
			RequestString:  req.Query,
			VariableValues: req.Variables,
			Context:        ctx,
		})

		c.JSON(http.StatusOK, result)
	})

	return r
}
