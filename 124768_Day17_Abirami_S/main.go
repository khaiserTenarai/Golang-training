package main

import (
	"log"

	"day17-employee-platform/config"
	gqlserver "day17-employee-platform/graphql"
	"day17-employee-platform/handler"
	"day17-employee-platform/middleware"
	"day17-employee-platform/model"
	"day17-employee-platform/repository"
	"day17-employee-platform/service"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg := config.Load()

	db, err := gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{})
	if err != nil {
		log.Fatal("database connection failed: ", err)
	}

	if err := db.AutoMigrate(
		&model.Department{},
		&model.Employee{},
	); err != nil {
		log.Fatal("database migration failed: ", err)
	}

	employeeRepo := repository.NewEmployeeRepository(db)
	departmentRepo := repository.NewDepartmentRepository(db)

	employeeService := service.NewEmployeeService(employeeRepo)
	departmentService := service.NewDepartmentService(departmentRepo)

	authHandler := handler.NewAuthHandler(cfg)
	employeeHandler := handler.NewEmployeeHandler(employeeService)

	graphQLServer, err := gqlserver.NewServer(employeeService, departmentService)
	if err != nil {
		log.Fatal("GraphQL schema failed: ", err)
	}

	r := gin.Default()

	r.POST("/login", authHandler.Login)

	r.GET("/employees", employeeHandler.List)
	r.GET("/employees/:id", employeeHandler.Get)

	protected := r.Group("/employees")
	protected.Use(middleware.RequireAuth(cfg.JWTSecret))
	protected.Use(middleware.RequireAdmin())
	{
		protected.POST("", employeeHandler.Create)
		protected.PUT("/:id", employeeHandler.Update)
		protected.DELETE("/:id", employeeHandler.Delete)
	}

	graphQLHandler := gqlserver.WithOptionalRole(graphQLServer.HTTPHandler(), cfg.JWTSecret)
	r.Any("/graphql", gin.WrapH(graphQLHandler))

	log.Println("REST: http://localhost:" + cfg.ServerPort + "/employees")
	log.Println("Login: http://localhost:" + cfg.ServerPort + "/login")
	log.Println("GraphQL: http://localhost:" + cfg.ServerPort + "/graphql")

	log.Fatal(r.Run(":" + cfg.ServerPort))
}
