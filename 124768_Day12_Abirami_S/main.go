package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"day12/controller"
	"day12/middleware"
	"day12/repository"
	"day12/service"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found; using system environment variables")
	}

	port := env("PORT", "8080")
	databaseURL := os.Getenv("DATABASE_URL")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		log.Fatalf("could not configure database: %v", err)
	}
	defer db.Close()

	repo := repository.NewEmployeeRepository(db)
	svc := service.NewEmployeeService(repo)
	ctl := controller.NewEmployeeController(svc, db)

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.RequestID())
	r.Use(middleware.Logging())

	r.GET("/health", ctl.Health)
	r.GET("/ready", ctl.Ready)

	api := r.Group("/employees")
	api.Use(middleware.Authenticate())
	{
		api.GET("", ctl.GetAll)
		api.GET("/:id", ctl.GetByID)
		api.POST("", middleware.RequireAdmin(), ctl.Create)
		api.PUT("/:id", middleware.RequireAdmin(), ctl.Update)
		api.DELETE("/:id", middleware.RequireAdmin(), ctl.Delete)
	}

	server := &http.Server{
		Addr:              ":" + port,
		Handler:           r,
		ReadHeaderTimeout: 5 * time.Second,
	}
	log.Printf("Employee API listening on :%s", port)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func env(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
