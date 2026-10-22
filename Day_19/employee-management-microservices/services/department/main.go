
package main

import (
	"log"
	"os"

	"employee-management-microservices/database"
	"employee-management-microservices/services/department/controller"
	"employee-management-microservices/services/department/entity"
	"employee-management-microservices/services/department/repository"
	"employee-management-microservices/services/department/routes"
	"employee-management-microservices/services/department/service"
	"employee-management-microservices/shared/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load(".env")

	db, err := database.ConnectDB("DEPARTMENT")
	if err != nil {
		log.Fatal(err)
	}

	if err := db.AutoMigrate(&entity.Department{}); err != nil {
		log.Fatal("department migration failed: ", err)
	}

	repo := repository.NewDepartmentRepository(db)
	svc := service.NewDepartmentService(repo)
	ctrl := controller.NewDepartmentController(svc)

	router := gin.Default()

	// Add Correlation ID middleware.
	router.Use(middleware.CorrelationID())

	routes.RegisterDepartmentRoutes(router, ctrl)

	port := os.Getenv("DEPARTMENT_SERVICE_PORT")
	if port == "" {
		port = "8082"
	}

	log.Printf("Department Service running on port %s", port)
	log.Fatal(router.Run(":" + port))
}
