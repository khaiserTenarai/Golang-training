package main

import (
	"log"

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
)

func main() {
	cfg := config.Load()

	db, err := gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{})
	if err != nil {
		log.Fatal("database connection failed:", err)
	}
	// Department first: employees.department_id references it.
	if err := db.AutoMigrate(&model.Department{}, &model.Employee{}); err != nil {
		log.Fatal("migration failed:", err)
	}
	if err := repository.SeedSampleData(db); err != nil {
		log.Fatal("seeding failed:", err)
	}

	// repository -> service -> (REST controller | GraphQL resolver)
	empRepo := repository.NewEmployeeRepository(db)
	deptRepo := repository.NewDepartmentRepository(db)
	empSvc := service.NewEmployeeService(empRepo, deptRepo)
	deptSvc := service.NewDepartmentService(deptRepo, empRepo)

	empCtrl := controller.NewEmployeeController(empSvc)
	deptCtrl := controller.NewDepartmentController(deptSvc)

	gs, err := gqlserver.NewServer(empSvc, deptSvc)
	if err != nil {
		log.Fatal("GraphQL schema failed:", err)
	}

	r := gin.New()
	r.Use(gin.Recovery(), middleware.Logging())

	// REST
	r.GET("/employees", empCtrl.GetAll)
	r.POST("/employees", empCtrl.Create)
	r.GET("/employees/:id", empCtrl.GetByID)
	r.PUT("/employees/:id", empCtrl.Update)
	r.DELETE("/employees/:id", empCtrl.Delete)
	r.GET("/departments", deptCtrl.GetAll)

	// GraphQL
	h := gqlserver.NewHTTPHandler(gs)
	r.Any("/graphql", gin.WrapH(h))
	r.GET("/playground", gin.WrapH(h))

	log.Println("REST:    http://localhost:" + cfg.ServerPort + "/employees")
	log.Println("GraphQL: http://localhost:" + cfg.ServerPort + "/graphql")
	log.Println("UI:      http://localhost:" + cfg.ServerPort + "/playground")
	log.Fatal(r.Run(":" + cfg.ServerPort))
}
