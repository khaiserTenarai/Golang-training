package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"employee-service/client"
	"employee-service/config"
	"employee-service/dto"
	"employee-service/middleware"
	"employee-service/repository"
	"employee-service/service"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	cfg := config.LoadConfig()

	dbURI := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	pool, err := pgxpool.New(context.Background(), dbURI)
	if err != nil {
		fmt.Printf("Database connection failed: %v\n", err)
		return
	}
	defer pool.Close()

	empRepo := repository.NewEmployeeRepository(pool)
	deptClient := client.NewDepartmentClient(cfg.DepartmentServiceURL)
	empService := service.NewEmployeeService(empRepo, deptClient)

	r := gin.Default()
	r.Use(middleware.CorrelationIDMiddleware())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"service": "employee-service", "status": "UP"})
	})

	api := r.Group("/api/v1")
	{
		api.POST("/employees", func(c *gin.Context) {
			var req dto.CreateEmployeeRequest
			if err := c.ShouldBindJSON(&req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			resp, err := empService.CreateEmployee(c.Request.Context(), req)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusCreated, resp)
		})

		api.GET("/employees/:id", func(c *gin.Context) {
			id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			corrID := c.GetString(middleware.CorrelationIDHeader)
			resp, err := empService.GetEmployeeByID(c.Request.Context(), id, corrID)
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
				return
			}
			c.JSON(http.StatusOK, resp)
		})

		api.GET("/employees", func(c *gin.Context) {
			resps, err := empService.GetAllEmployees(c.Request.Context())
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, resps)
		})

		api.PUT("/employees/:id", func(c *gin.Context) {
			id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			var req dto.UpdateEmployeeRequest
			if err := c.ShouldBindJSON(&req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			resp, err := empService.UpdateEmployee(c.Request.Context(), id, req)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, resp)
		})

		api.DELETE("/employees/:id", func(c *gin.Context) {
			id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			if err := empService.DeleteEmployee(c.Request.Context(), id); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusNoContent, nil)
		})
	}

	r.Run(":" + cfg.AppPort)
}
