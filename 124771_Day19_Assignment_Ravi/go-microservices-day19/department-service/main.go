package main

import (
	"day19/department-service/config"
	"day19/department-service/database"
	"day19/department-service/handler"
	"day19/department-service/health"
	"day19/department-service/middleware"
	"day19/department-service/model"
	"day19/department-service/repository"
	"day19/department-service/service"
	"github.com/gin-gonic/gin"
	"log"
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
	// AutoMigrate creates/updates this service's table for this learning project.
	if err := db.AutoMigrate(&model.Department{}); err != nil {
		log.Fatal(err)
	}
	repo := repository.NewDepartmentRepository(db)
	svc := service.NewDepartmentService(repo)
	h := handler.NewDepartmentHandler(svc)
	r := gin.New()
	r.Use(gin.Recovery(), middleware.CorrelationID())
	health.Register(r, db)
	h.Register(r)
	log.Printf("Department Service listening on :%s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
