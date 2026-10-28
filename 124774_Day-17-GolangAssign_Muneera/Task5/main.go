package main

import (
	"log"

	"employee-management/config"
	gqlserver "employee-management/graphql"
	"employee-management/middleware"
	"employee-management/model"
	"employee-management/repository"
	"employee-management/service"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg := config.Load()

	db, err := gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{})
	if err != nil {
		log.Fatalf("connect to PostgreSQL: %v", err)
	}

	if err := db.AutoMigrate(&model.Employee{}); err != nil {
		log.Fatalf("migrate database: %v", err)
	}

	employeeRepo := repository.NewEmployeeRepository(db)
	employeeService := service.NewEmployeeService(employeeRepo)

	graphQLServer, err := gqlserver.NewServer(employeeService)
	if err != nil {
		log.Fatalf("initialize GraphQL: %v", err)
	}

	router := gin.New()
	router.Use(gin.Recovery(), middleware.Logging())

	graphQLHandler := gqlserver.NewHTTPHandler(*graphQLServer)
	router.Any("/graphql", gin.WrapH(graphQLHandler))
	router.GET("/playground", gin.WrapH(graphQLHandler))

	log.Printf("GraphQL endpoint: http://localhost:%s/graphql", cfg.ServerPort)
	log.Printf("GraphQL Playground: http://localhost:%s/playground", cfg.ServerPort)

	if err := router.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("start server: %v", err)
	}
}
