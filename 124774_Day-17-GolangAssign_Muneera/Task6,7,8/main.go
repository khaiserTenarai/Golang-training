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

	if cfg.DBPassword == "" {
		log.Fatal("DB_PASSWORD must be set in the environment or .env file")
	}

	db, err := gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{})
	if err != nil {
		log.Fatal("database connection failed: ", err)
	}

	if err := db.AutoMigrate(&model.Employee{}); err != nil {
		log.Fatal("database migration failed: ", err)
	}

	repo := repository.NewEmployeeRepository(db)
	svc := service.NewEmployeeService(repo)

	gs, err := gqlserver.NewServer(svc)
	if err != nil {
		log.Fatal("GraphQL schema failed: ", err)
	}

	r := gin.New()
	r.Use(gin.Recovery(), middleware.Logging())

	handler := gqlserver.NewHTTPHandler(gs)

	// Authentication protects GraphQL requests.
	r.Any("/graphql", middleware.Auth(), gin.WrapH(handler))

	// Playground page is accessible; requests to /graphql
	// still require an Authorization header.
	r.GET("/playground", gin.WrapH(handler))

	log.Println("GraphQL: http://localhost:" + cfg.ServerPort + "/graphql")
	log.Println("Playground: http://localhost:" + cfg.ServerPort + "/playground")

	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatal("server failed: ", err)
	}
}
