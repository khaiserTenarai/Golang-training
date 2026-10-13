package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Employee struct {
	ID         uint    `json:"id" gorm:"primaryKey"`
	Name       string  `json:"name" binding:"required" gorm:"not null"`
	Email      string  `json:"email" binding:"required,email" gorm:"uniqueIndex;not null"`
	Department string  `json:"department" binding:"required" gorm:"not null"`
	Salary     float64 `json:"salary" binding:"required,gte=0"`
}

func main() {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "postgres"),
		getEnv("DB_NAME", "employee_db"),
		getEnv("DB_SSLMODE", "disable"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("connect to PostgreSQL: %v", err)
	}
	if err := db.AutoMigrate(&Employee{}); err != nil {
		log.Fatalf("migrate database: %v", err)
	}

	router := gin.Default()
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	api := router.Group("/employees")

	api.POST("", func(c *gin.Context) {
		var employee Employee
		if err := c.ShouldBindJSON(&employee); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Create(&employee).Error; err != nil {
			c.JSON(http.StatusConflict, gin.H{"error": "could not create employee; email may already exist"})
			return
		}
		c.JSON(http.StatusCreated, employee)
	})

	api.GET("", func(c *gin.Context) {
		var employees []Employee
		if err := db.Order("id ASC").Find(&employees).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch employees"})
			return
		}
		c.JSON(http.StatusOK, employees)
	})

	api.GET("/:id", func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil || id == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid employee id"})
			return
		}
		var employee Employee
		if err := db.First(&employee, id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "employee not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch employee"})
			return
		}
		c.JSON(http.StatusOK, employee)
	})

	api.PUT("/:id", func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil || id == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid employee id"})
			return
		}
		var employee Employee
		if err := db.First(&employee, id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "employee not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch employee"})
			return
		}
		var input Employee
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		employee.Name = input.Name
		employee.Email = input.Email
		employee.Department = input.Department
		employee.Salary = input.Salary
		if err := db.Save(&employee).Error; err != nil {
			c.JSON(http.StatusConflict, gin.H{"error": "could not update employee; email may already exist"})
			return
		}
		c.JSON(http.StatusOK, employee)
	})

	api.DELETE("/:id", func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil || id == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid employee id"})
			return
		}
		result := db.Delete(&Employee{}, id)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not delete employee"})
			return
		}
		if result.RowsAffected == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "employee not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "employee deleted"})
	})

	port := getEnv("PORT", "8080")
	log.Printf("Employee API listening on :%s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("start server: %v", err)
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
