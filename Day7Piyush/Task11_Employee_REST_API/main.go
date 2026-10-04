package main

import (
	"fmt"
	"log"
	"net/http"
	"task11_employee_rest_api/config"
	"task11_employee_rest_api/controller"
	"task11_employee_rest_api/repository"
	"task11_employee_rest_api/routes"
)

func main() {
	db := config.ConnectDB()
	defer db.Close()

	config.CreateTable(db)

	repo := repository.NewEmployeeRepository(db)
	ctrl := controller.NewEmployeeController(repo)
	router := routes.SetupRoutes(ctrl)

	fmt.Println("Server running on http://localhost:8080")
	fmt.Println("Endpoints:")
	fmt.Println("  POST   /api/employees")
	fmt.Println("  GET    /api/employees")
	fmt.Println("  GET    /api/employees/{id}")
	fmt.Println("  PUT    /api/employees/{id}")
	fmt.Println("  DELETE /api/employees/{id}")
	log.Fatal(http.ListenAndServe(":8080", router))
}
