package main

import (
	"log"
	"net/http"

	"employee-management/config"
	"employee-management/controller"

	"employee-management/graphql"

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
		log.Fatal("configuration error: ", err)
	}

	// Connect to PostgreSQL.
	db, err := gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{
		TranslateError: true,
	})
	if err != nil {
		log.Fatal("database connection failed: ", err)
	}

	// Create the employee table if it does not exist.
	if err := db.AutoMigrate(&model.Employee{}); err != nil {
		log.Fatal("database migration failed: ", err)
	}

	// Initialize the repository, service, and controller layers.
	employeeRepo := repository.NewEmployeeRepository(db)
	employeeService := service.NewEmployeeService(employeeRepo)
	employeeController := controller.NewEmployeeController(employeeService)

	// Initialize GraphQL.
	graphQLServer, err := graphql.NewServer(employeeService)
	if err != nil {
		log.Fatal("GraphQL initialization failed: ", err)
	}
	graphQLHandler := graphql.NewHTTPHandler(graphQLServer)

	// Configure the Gin HTTP server.
	router := gin.New()
	router.Use(gin.Recovery())
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

	// REST API routes.
	employees := router.Group("/employees")
	{
		employees.GET("", employeeController.GetAll)
		employees.POST("", employeeController.Create)
		employees.GET("/:id", employeeController.GetByID)
		employees.PUT("/:id", employeeController.Update)
		employees.DELETE("/:id", employeeController.Delete)
	}

	// GraphQL endpoint. GraphiQL is available in the browser.
	router.Any("/graphql", gin.WrapH(graphQLHandler))

	// Start the server.
	address := ":" + cfg.ServerPort
	log.Printf("server starting on http://localhost%s", address)
	log.Printf("GraphQL endpoint: http://localhost%s/graphql", address)

	if err := router.Run(address); err != nil {
		log.Fatal("server failed: ", err)
	}
}