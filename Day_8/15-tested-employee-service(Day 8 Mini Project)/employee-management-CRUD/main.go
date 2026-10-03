package main

import (
	"fmt"
	"os"

	"ems/controller"
	"ems/database"
	"ems/repository"
	"ems/service"
	"ems/view"
)

func main() {

	db, err := database.Connect()

	if err != nil {
		fmt.Println("Database connection failed:", err)
		os.Exit(1)
	}

	defer db.Close()

	employeeRepository := repository.NewEmployeeRepository(db)

	employeeService := service.NewEmployeeService(
		employeeRepository,
	)

	employeeView := view.NewEmployeeView()

	employeeController := controller.NewEmployeeController(
		employeeView,
		employeeService,
	)

	employeeController.Start()
}
