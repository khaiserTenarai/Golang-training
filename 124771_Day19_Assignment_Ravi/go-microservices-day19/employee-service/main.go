package main

import (
	"day19/employee-service/client"
	"day19/employee-service/config"
	"day19/employee-service/database"
	"day19/employee-service/handler"
	"day19/employee-service/health"
	"day19/employee-service/middleware"
	"day19/employee-service/model"
	"day19/employee-service/repository"
	"day19/employee-service/resilience"
	"day19/employee-service/service"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	db, err := database.Open(cfg)
	if err != nil {
		log.Fatal(err)
	}
	// Task 3: AutoMigrate prepares the Employee table in PostgreSQL.
	if err := db.AutoMigrate(&model.Employee{}); err != nil {
		log.Fatal(err)
	}
	repo := repository.NewEmployeeRepository(db)
	// Tasks 10–13: configured timeout, retry count, circuit breaker and correlation propagation.
	breaker := resilience.NewCircuitBreaker(cfg.CircuitFailureThreshold, time.Duration(cfg.CircuitOpenSeconds)*time.Second)
	departmentClient := client.NewDepartmentClient(cfg.DepartmentServiceURL, time.Duration(cfg.HTTPTimeoutSeconds)*time.Second, cfg.RetryCount, breaker)
	employeeService := service.NewEmployeeService(repo, departmentClient)
	employeeHandler := handler.NewEmployeeHandler(employeeService)
	r := gin.New()
	r.Use(gin.Recovery(), middleware.CorrelationID())
	health.Register(r, db)
	employeeHandler.Register(r)
	log.Printf("Employee Service listening on :%s; Department Service URL=%s", cfg.Port, cfg.DepartmentServiceURL)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
