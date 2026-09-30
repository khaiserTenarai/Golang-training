package main
import (
	"fmt"
	"log"

	"salary-management/config"
	"salary-management/controller"
	"salary-management/database"
	"salary-management/repository"
	"salary-management/service"
	"salary-management/view"
)

func main() {

	// Load configuration
	cfg := config.Load()

	// Connect to database
	db, err := database.Connect(cfg)

	if err != nil {
		log.Fatal(
			"Database connection failed:",
			err,
		)
	}

	defer db.Close()

	fmt.Println()
	fmt.Println("==========================================")
	fmt.Println(cfg.AppName)
	fmt.Println("==========================================")

	// -----------------------------
	// Employee Repository
	// -----------------------------

	employeeRepository :=
		repository.NewEmployeeRepository(db)

	// -----------------------------
	// Employee Service
	// -----------------------------

	employeeService :=
		service.NewEmployeeService(
			employeeRepository,
		)

	// -----------------------------
	// Salary Repository
	// -----------------------------

	salaryRepository :=
		repository.NewSalaryRepository(db)

	// -----------------------------
	// Salary Service
	// -----------------------------

	salaryService :=
		service.NewSalaryService(
			salaryRepository,
		)

	// -----------------------------
	// Views
	// -----------------------------

	employeeView :=
		view.NewEmployeeView()

	salaryView :=
		view.NewSalaryView()

	// -----------------------------
	// Controller
	// -----------------------------

	employeeController :=
		controller.NewEmployeeController(
			employeeView,
			employeeService,
			salaryService,
			salaryView,
		)

	// Start application
	employeeController.Start()
}
