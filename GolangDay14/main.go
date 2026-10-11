package main

import (
	"log"
	"net/http"
	"os"

	"employee_management/config"
	"employee_management/controller"
	"employee_management/middleware"
	"employee_management/repository"
	"employee_management/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	config.ConnectDatabase()

	r := gin.Default()
	r.Use(middleware.LoggerMiddleware())

	os.MkdirAll("public", 0755)
	r.Static("/static", "./public")

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Page not found - Custom 404"})
	})
	r.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed - Custom 405"})
	})

	repo := repository.EmployeeRepository{}
	svc := service.EmployeeService{Repo: repo}
	ctrl := controller.EmployeeController{Service: svc}

	v1 := r.Group("/api/v1")
	{
		employees := v1.Group("/employees")
		{
			employees.GET("", ctrl.GetEmployees)
			employees.GET("/:id", ctrl.GetEmployeeByID)
			employees.POST("", ctrl.CreateEmployee)
			employees.PUT("/:id", ctrl.UpdateEmployeePUT)
			employees.PATCH("/:id", ctrl.UpdateEmployeePATCH)
			employees.DELETE("/:id", ctrl.DeleteEmployee)
			employees.POST("/form", ctrl.SubmitForm)
		}
	}

	r.Run(":8080")
}