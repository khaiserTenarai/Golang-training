package main

import (
	"log"
	"net/http"

	"employee-management/config"
	graphserver "employee-management/graphql"
	"employee-management/middleware"
	"employee-management/model"
	"employee-management/repository"
	"employee-management/service"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Load environment configuration
	cfg := config.LoadConfig()

	// Connect to PostgreSQL
	db, err := gorm.Open(
		postgres.Open(cfg.DatabaseURL()),
		&gorm.Config{},
	)
	if err != nil {
		log.Fatal("Database connection failed: ", err)
	}

	log.Println("Database connected successfully")

	// Create employees table if it does not exist
	if err := db.AutoMigrate(&model.Employee{}); err != nil {
		log.Fatal("Database migration failed: ", err)
	}

	// Create repository
	employeeRepo := repository.NewEmployeeRepository(db)

	// Create service
	employeeService := service.NewEmployeeService(employeeRepo)

	// Create GraphQL server
	gqlServer, err := graphserver.NewServer(employeeService)
	if err != nil {
		log.Fatal("GraphQL schema creation failed: ", err)
	}

	// Create Gin server
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.Logging())

	// Health check
	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// GraphQL endpoint and GraphiQL interface
	router.Any("/graphql", gin.WrapH(gqlServer.NewHTTPHandler()))

	// Start the server
	address := ":" + cfg.ServerPort
	log.Println("Server starting at http://localhost" + address)
	log.Println("GraphQL endpoint: http://localhost" + address + "/graphql")

	if err := router.Run(address); err != nil {
		log.Fatal("Server failed: ", err)
	}
}
