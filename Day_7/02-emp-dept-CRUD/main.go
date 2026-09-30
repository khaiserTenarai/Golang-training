package main
import (
	"fmt"
	"os"

	"department-management/controller"
	"department-management/database"
	"department-management/repository"
	"department-management/service"
	"department-management/view"
)

func main() {

	db, err := database.Connect()

	if err != nil {
		fmt.Println("Database connection failed:", err)
		os.Exit(1)
	}

	defer db.Close()

	// Department
	departmentRepository := repository.NewDepartmentRepository(db)

	departmentService := service.NewDepartmentService(
		departmentRepository,
	)

	departmentView := view.NewDepartmentView()

	departmentController := controller.NewDepartmentController(
		departmentView,
		departmentService,
	)

	// Employee
	employeeRepository := repository.NewEmployeeRepository(db)

	employeeService := service.NewEmployeeService(
		employeeRepository,
	)

	employeeView := view.NewEmployeeView()

	employeeController := controller.NewEmployeeController(
		employeeView,
		employeeService,
	)

	for {

		fmt.Println("\n========== Department Management System ==========")
		fmt.Println("1. Department Management")
		fmt.Println("2. Employee Management")
		fmt.Println("3. Exit")

		var choice int

		fmt.Print("Enter choice: ")
		fmt.Scan(&choice)

		switch choice {

		case 1:
			departmentController.Start()

		case 2:
			employeeController.Start()

		case 3:
			fmt.Println("Thank you..")
			return

		default:
			fmt.Println("Invalid choice.")
		}
	}
}
