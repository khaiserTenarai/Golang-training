package main

import (
	"fmt"
	"log"

	"employee-management/config"
	"employee-management/controller"
	graph "employee-management/graphql"
	"employee-management/middleware"
	"employee-management/model"
	"employee-management/repository"
	"employee-management/service"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// 1. Load application configuration.
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	// 2. Connect to PostgreSQL.
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}

	// 3. Create the employee table if it does not exist.
	if err := db.AutoMigrate(&model.Employee{}); err != nil {
		log.Fatal("failed to migrate database: ", err)
	}

	// 4. Create repository and service.
	employeeRepo := repository.NewEmployeeRepository(db)
	employeeService := service.NewEmployeeService(employeeRepo)

	// 5. Build the modular GraphQL schema.
	graphServer, err := graph.NewServer(employeeService)
	if err != nil {
		log.Fatal("failed to create GraphQL schema: ", err)
	}

	// 6. Create the REST controller.
	employeeController := controller.NewEmployeeController(employeeService)

	// 7. Configure Gin and logging middleware.
	router := gin.New()
	router.Use(middleware.Logging())
	router.Use(gin.Recovery())

	// REST routes.
	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
	router.GET("/employees", employeeController.GetAll)
	router.GET("/employees/:id", employeeController.GetByID)

	// GraphQL endpoint.
	router.Any("/graphql", gin.WrapH(graph.NewHTTPHandler(graphServer)))

	// 8. Start the server.
	log.Println("REST API: http://localhost:" + cfg.ServerPort + "/employees")
	log.Println("GraphQL:  http://localhost:" + cfg.ServerPort + "/graphql")

	if err := router.Run(":" + cfg.ServerPort); err != nil {
		log.Fatal(err)
	}
}
