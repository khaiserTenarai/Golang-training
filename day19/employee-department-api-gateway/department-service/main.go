package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"department-service/config"
	"department-service/dto"
	"department-service/middleware"
	"department-service/repository"
	"department-service/service"

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

	deptRepo := repository.NewDepartmentRepository(pool)
	deptService := service.NewDepartmentService(deptRepo)

	r := gin.Default()
	r.Use(middleware.CorrelationIDMiddleware())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"service": "department-service", "status": "UP"})
	})

	api := r.Group("/api/v1")
	{
		api.POST("/departments", func(c *gin.Context) {
			var req dto.CreateDepartmentRequest
			if err := c.ShouldBindJSON(&req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			resp, err := deptService.CreateDepartment(c.Request.Context(), req)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusCreated, resp)
		})

		api.GET("/departments/:id", func(c *gin.Context) {
			id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
			resp, err := deptService.GetDepartmentByID(c.Request.Context(), id)
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "Department not found"})
				return
			}
			c.JSON(http.StatusOK, resp)
		})

		api.GET("/departments", func(c *gin.Context) {
			resps, err := deptService.GetAllDepartments(c.Request.Context())
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, resps)
		})
	}

	r.Run(":" + cfg.AppPort)
}
