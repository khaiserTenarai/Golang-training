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
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load configuration: ", err)
	}

	// Connect to PostgreSQL
	db, err := gorm.Open(
		postgres.Open(cfg.DSN()),
		&gorm.Config{TranslateError: true},
	)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	// Automatically create or update database tables
	if err := db.AutoMigrate(&model.Employee{}); err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	// Initialize repository, service, and controller
	employeeRepository := repository.NewEmployeeRepository(db)
	employeeService := service.NewEmployeeService(employeeRepository)
	employeeController := controller.NewEmployeeController(employeeService)

	// Create Gin router
	router := gin.New()

	// Register middleware
	router.Use(middleware.RequestID())
	router.Use(middleware.Logging())
	router.Use(middleware.Recovery())
	router.Use(middleware.CORSMiddleware())

	// Health check
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

	// Task 3: Demonstrate Next()
	router.GET("/middleware-next", middleware.NextDemo(), func(c *gin.Context) {
		log.Println("Handler: Executing route handler")

		c.JSON(http.StatusOK, gin.H{
			"message": "Next() demonstration completed",
		})
	})

	// Task 4: Demonstrate Abort()
	router.GET("/middleware-abort", middleware.AbortDemo(), func(c *gin.Context) {
		log.Println("Handler: This should not execute")

		c.JSON(http.StatusOK, gin.H{
			"message": "Handler executed",
		})
	})

	// Task 6: Demonstrate Recovery middleware
	router.GET("/panic", func(c *gin.Context) {
		panic("testing recovery middleware")
	})

	// Employee CRUD APIs
	router.GET("/employees", employeeController.GetAll)
	router.POST("/employees", employeeController.Create)
	router.GET("/employees/:id", employeeController.GetByID)
	router.PUT("/employees/:id", employeeController.Update)
	router.DELETE("/employees/:id", employeeController.Delete)

	// Start server
	log.Println("Server starting on http://localhost:" + cfg.ServerPort)

	if err := router.Run(":" + cfg.ServerPort); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
