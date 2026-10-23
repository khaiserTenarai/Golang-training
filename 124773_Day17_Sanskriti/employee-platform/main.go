package main

import (
	"log"

	"employee-platform/config"
	"employee-platform/controller"
	"employee-platform/database"
	"employee-platform/graphql"
	"employee-platform/middleware"
	"employee-platform/repository"
	"employee-platform/service"
	//"employee-platform/utility"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	pool, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	repo := repository.New(pool)
	svc := service.New(repo)
	rest := controller.New(svc)
	gql := graphql.New(svc)

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery(), middleware.RequestID())

	r.GET("/health", func(c *gin.Context) { c.JSON(200, gin.H{"status": "UP"}) })
	r.POST("/auth/login", rest.Login)
	api := r.Group("")
	api.Use(middleware.Authenticate(cfg.JWTSecret))
	api.GET("/employees", rest.ListEmployees)
	api.GET("/employees/:id", rest.GetEmployee)
	api.POST("/employees", middleware.RequireRole("admin"), rest.CreateEmployee)
	api.PUT("/employees/:id", middleware.RequireRole("admin"), rest.UpdateEmployee)
	api.DELETE("/employees/:id", middleware.RequireRole("admin"), rest.DeleteEmployee)
	api.GET("/departments", rest.ListDepartments)

	// GraphQL supports the same authenticated user context and shared service.
	r.POST("/graphql", middleware.Authenticate(cfg.JWTSecret), gql.Serve)
	r.GET("/graphql/playground", func(c *gin.Context) {
		c.Data(200, "text/html; charset=utf-8", []byte(graphql.PlaygroundHTML))
	})

	log.Printf("Employee platform listening on :%s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
