package main

import (
	"log"
	"os"

	"employee-management-microservices/database"
	"employee-management-microservices/services/employee/controller"
	"employee-management-microservices/services/employee/entity"
	"employee-management-microservices/services/employee/repository"
	"employee-management-microservices/services/employee/routes"
	"employee-management-microservices/services/employee/service"
	"employee-management-microservices/shared/httpclient"
	"employee-management-microservices/shared/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load(".env")

	db, err := database.ConnectDB("EMPLOYEE")
	if err != nil {
		log.Fatal(err)
	}

	if err := db.AutoMigrate(&entity.Employee{}); err != nil {
		log.Fatal("employee migration failed: ", err)
	}

	repo := repository.NewEmployeeRepository(db)
	departmentClient := httpclient.NewDepartmentClient()
	svc := service.NewEmployeeService(repo, departmentClient)
	ctrl := controller.NewEmployeeController(svc)

	router := gin.Default()
	router.Use(middleware.CorrelationID())
	routes.RegisterEmployeeRoutes(router, ctrl)

	port := os.Getenv("EMPLOYEE_SERVICE_PORT")
	if port == "" {
		port = "8081"
	}

	log.Printf("Employee Service running on port %s", port)
	log.Fatal(router.Run(":" + port))
}
