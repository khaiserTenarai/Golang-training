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
	// Load configuration.
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("configuration error: ", err)
	}

	// Connect to PostgreSQL.
	db, err := gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{
		TranslateError: true,
	})
	if err != nil {
		log.Fatal("database connection failed: ", err)
	}

	// Create database tables if necessary.
	if err := db.AutoMigrate(&model.Employee{}); err != nil {
		log.Fatal("database migration failed: ", err)
	}

	// Initialize application layers.
	employeeRepo := repository.NewEmployeeRepository(db)
	employeeService := service.NewEmployeeService(employeeRepo)
	employeeController := controller.NewEmployeeController(employeeService)

	// Create the Gin router.
	router := gin.New()

	// Register middleware in execution order.
	router.Use(gin.Recovery())
	router.Use(middleware.First())
	router.Use(middleware.Second())
	router.Use(middleware.Logging())

	// Health check.
	router.GET("/health", func(c *gin.Context) {
		sqlDB, err := db.DB()
		if err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"status": "unhealthy",
			})
			return
		}

		if err := sqlDB.Ping(); err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"status": "unhealthy",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "healthy",
		})
	})

	// Endpoint to demonstrate middleware execution order.
	router.GET("/middleware-order", func(c *gin.Context) {
		log.Println("Handler: Processing request")

		c.JSON(http.StatusOK, gin.H{
			"message": "Middleware execution order demonstrated",
		})
	})

	// Employee REST API routes.
	employees := router.Group("/employees")
	{
		employees.GET("", employeeController.GetAll)
		employees.POST("", employeeController.Create)
		employees.GET("/:id", employeeController.GetByID)
		employees.PUT("/:id", employeeController.Update)
		employees.DELETE("/:id", employeeController.Delete)
	}

	// Start the server.
	address := ":" + cfg.ServerPort
	log.Printf("server starting on http://localhost%s", address)

	if err := router.Run(address); err != nil {
		log.Fatal("server failed: ", err)
	}
}