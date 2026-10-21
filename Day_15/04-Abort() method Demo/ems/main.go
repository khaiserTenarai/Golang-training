package main

import (
	"log"
	"net/http"

	"employee-management/config"
	"employee-management/controller"
	"employee-management/middleware"
	"employee-management/model"
	"employee-management/repository"
	"employee-management/service"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Load application configuration.
cfg, err := config.Load()
if err != nil {
    log.Fatal("Failed to load configuration: ", err)
}

	// Connect to PostgreSQL database.
	db, err := gorm.Open(
		postgres.Open(cfg.DSN()),
		&gorm.Config{
			TranslateError: true,
		},
	)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	// Automatically create or update the Employee table.
	if err := db.AutoMigrate(&model.Employee{}); err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	// Initialize repository, service, and controller.
	employeeRepository := repository.NewEmployeeRepository(db)
	employeeService := service.NewEmployeeService(employeeRepository)
	employeeController := controller.NewEmployeeController(employeeService)

	// Create Gin router without default middleware.
	router := gin.New()

	// Recover from unexpected panics.
	router.Use(gin.Recovery())

	// Task 1: Register custom logging middleware globally.
	router.Use(middleware.Logging())

	// Health check endpoint.
	router.GET("/health", func(c *gin.Context) {
		sqlDB, err := db.DB()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "database connection error",
			})
			return
		}

		if err := sqlDB.Ping(); err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"status": "database unavailable",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "healthy",
		})
	})

	// Task 3: Demonstrate c.Next().
	router.GET(
		"/middleware-next",
		middleware.NextDemo(),
		func(c *gin.Context) {
			log.Println("Handler: Executing route handler")

			c.JSON(http.StatusOK, gin.H{
				"message": "Next() demonstration completed",
			})
		},
	)

	// Task 4: Demonstrate c.Abort().
	router.GET(
		"/middleware-abort",
		middleware.AbortDemo(),
		func(c *gin.Context) {
			// This handler should not execute because
			// AbortDemo() calls c.Abort().
			log.Println("Handler: This should not execute")

			c.JSON(http.StatusOK, gin.H{
				"message": "Handler executed",
			})
		},
	)

	// Employee CRUD endpoints.
	router.GET("/employees", employeeController.GetAll)
	router.POST("/employees", employeeController.Create)
	router.GET("/employees/:id", employeeController.GetByID)
	router.PUT("/employees/:id", employeeController.Update)
	router.DELETE("/employees/:id", employeeController.Delete)

	// Start the HTTP server.
	log.Println("Server starting on http://localhost:" + cfg.ServerPort)

	if err := router.Run(":" + cfg.ServerPort); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}