package main

import (
	"fmt"
	"log"
	"net/http"
	"task12_repository_service_architecture/config"
	"task12_repository_service_architecture/controller"
	"task12_repository_service_architecture/repository"
	"task12_repository_service_architecture/routes"
	"task12_repository_service_architecture/service"
)

func main() {
	// Database layer
	db := config.ConnectDB()
	defer db.Close()
	config.CreateTable(db)

	// Dependency Injection: Controller -> Service -> Repository -> PostgreSQL
	repo := repository.NewPostgresEmployeeRepository(db) // Repository (interface)
	svc := service.NewEmployeeService(repo)               // Service (interface, depends on repo interface)
	ctrl := controller.NewEmployeeController(svc)          // Controller (depends on service interface)
	router := routes.SetupRoutes(ctrl)

	fmt.Println("Architecture: Controller -> Service -> Repository -> PostgreSQL")
	fmt.Println("Server running on http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}
