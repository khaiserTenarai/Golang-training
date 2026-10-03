package main
import (
	"fmt"
	"strconv"

	"ems/config"
	"ems/controller"
	"ems/database"
	"ems/repository"
	"ems/service"
	"ems/view"
)

/*
	Main function.

	Layer flow:

	config
	    ↓
	database
	    ↓
	repository
	    ↓
	service
	    ↓
	view
	    ↓
	controller

	Day 9 concurrency happens
	inside the service layer.
*/
func main() {

	/*
		Load configuration.

		This reads config.env.
	*/
	cfg := config.Load()

	/*
		Convert WORKERS from string to int.

		Example:

		WORKERS=3

		"3" → 3
	*/
	workers, _ :=
		strconv.Atoi(cfg.Workers)

	/*
		Convert BUFFER from string to int.

		Example:

		BUFFER=2

		"2" → 2
	*/
	buffer, _ :=
		strconv.Atoi(cfg.Buffer)

	/*
		Connect to PostgreSQL.
	*/
	db, err := database.Connect()

	if err != nil {

		fmt.Println(
			"Database connection failed:",
			err,
		)

		return
	}

	/*
		Close database when main finishes.
	*/
	defer db.Close()

	/*
		Create repository object.
	*/
	employeeRepository :=
		repository.NewEmployeeRepository(db)

	/*
		Create service object.

		Pass:

		1. Repository
		2. Number of workers
		3. Channel buffer size
	*/
	employeeService :=
		service.NewEmployeeService(
			employeeRepository,
			workers,
			buffer,
		)

	/*
		Create view.
	*/
	employeeView :=
		view.NewEmployeeView()

	/*
		Create controller.
	*/
	employeeController :=
		controller.NewEmployeeController(
			employeeView,
			employeeService,
		)

	/*
		Start application.
	*/
	employeeController.Start()
}