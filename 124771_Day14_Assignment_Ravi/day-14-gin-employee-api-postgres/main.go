package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"day-14-gin-employee-api/config"
	"day-14-gin-employee-api/controller"
	"day-14-gin-employee-api/repository"
	"day-14-gin-employee-api/routes"
	"day-14-gin-employee-api/service"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	cfg := config.Load()

	dsnURL := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(cfg.DBUser, cfg.DBPassword),
		Host:   fmt.Sprintf("%s:%s", cfg.DBHost, cfg.DBPort),
		Path:   cfg.DBName,
	}
	query := dsnURL.Query()
	query.Set("sslmode", cfg.DBSSLMode)
	dsnURL.RawQuery = query.Encode()
	dsn := dsnURL.String()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Printf("database pool error: %v", err)
		return
	}
	defer db.Close()

	if err := db.Ping(ctx); err != nil {
		log.Printf("database connection error: %v", err)
		return
	}

	if err := repository.CreateSchema(ctx, db); err != nil {
		log.Printf("create table error: %v", err)
		return
	}

	repo := repository.NewPostgresEmployeeRepository(db)
	svc := service.NewEmployeeService(repo)
	ctl := controller.NewEmployeeController(svc)

	//task-1  Gin application.
	// Gin's built-in logger and recovery are used; no custom middleware package.
	r := gin.Default()
	routes.RegisterRoutes(r, ctl)

	r.Static("/static", "./static")
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "route not found"})
	})
	r.HandleMethodNotAllowed = true
	r.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "method not allowed"})
	})

	server := &http.Server{
		Addr:              ":" + cfg.ServerPort,
		Handler:           r,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("Gin Employee Management API started on http://localhost:%s", cfg.ServerPort)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Printf("server error: %v", err)
	}
}
