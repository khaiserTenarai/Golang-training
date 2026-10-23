package main

import (
	"employee-management/config"
	"employee-management/controller"
	gqlserver "employee-management/graphql"
	"employee-management/middleware"
	"employee-management/model"
	"employee-management/repository"
	"employee-management/service"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	cfg := config.Load()
	db, err := gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{})
	if err != nil {
		log.Fatal("database connection failed:", err)
	}
	if err := db.AutoMigrate(&model.Employee{}); err != nil {
		log.Fatal("migration failed:", err)
	}
	repo := repository.NewEmployeeRepository(db)
	svc := service.NewEmployeeService(repo)
	ctrl := controller.NewEmployeeController(svc)
	gs, err := gqlserver.NewServer(svc)
	if err != nil {
		log.Fatal("GraphQL schema failed:", err)
	}
	r := gin.New()
	r.Use(gin.Recovery(), middleware.Logging())
	r.GET("/employees", ctrl.GetAll)
	r.POST("/employees", ctrl.Create)
	r.GET("/employees/:id", ctrl.GetByID)
	r.PUT("/employees/:id", ctrl.Update)
	r.DELETE("/employees/:id", ctrl.Delete)
	
	h := gqlserver.NewHTTPHandler(gs)
	r.Any("/graphql", gin.WrapH(h))
	r.GET("/playground", gin.WrapH(h))
	log.Println("REST: http://localhost:" + cfg.ServerPort + "/employees")
	log.Println("GraphQL: http://localhost:" + cfg.ServerPort + "/graphql")
	log.Println("UI: http://localhost:" + cfg.ServerPort + "/playground")
	log.Fatal(r.Run(":" + cfg.ServerPort))
}
