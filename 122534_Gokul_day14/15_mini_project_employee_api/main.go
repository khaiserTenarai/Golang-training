/*
Question 15: Day 14 Mini Project - Build a Gin Employee CRUD REST API.

Run with: go run .
Endpoints:
    POST   /api/v1/employees
    GET    /api/v1/employees            (optional ?department=Tech)
    GET    /api/v1/employees/:id
    PUT    /api/v1/employees/:id
    PATCH  /api/v1/employees/:id
    DELETE /api/v1/employees/:id
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"

	"employeegin/controller"
	"employeegin/dao"
	"employeegin/service"
)

func readPort() string {
	fmt.Print("Enter the port (press Enter for 8080): ")
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = strings.TrimSpace(text)
	if text == "" {
		return "8080"
	}
	return text
}

func main() {
	employeeDAO := dao.NewMemoryEmployeeDAO()
	employeeService := service.NewEmployeeService(employeeDAO)
	employeeController := controller.NewEmployeeController(employeeService)

	router := gin.Default()
	controller.RegisterRoutes(router, employeeController)

	if err := router.Run(":" + readPort()); err != nil {
		log.Fatal("server stopped:", err)
	}
}
