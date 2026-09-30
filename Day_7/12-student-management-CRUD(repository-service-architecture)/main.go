package main
import (
	"fmt"
	"os"

	"student_app/controller"
	"student_app/database"
	"student_app/repository"
	"student_app/service"
	"student_app/view"
)

func main() {

	db, err := database.Connect()

	if err != nil {

		fmt.Println(
			"Database connection failed:",
			err,
		)

		os.Exit(1)
	}

	defer db.Close()

	// Repository Dependency

	studentRepository :=
		repository.NewStudentRepository(db)

	// Service Dependency

	studentService :=
		service.NewStudentService(
			studentRepository,
		)

	// View Dependency

	studentView :=
		view.NewStudentView()

	// Controller Dependency

	studentController :=
		controller.NewStudentController(
			studentView,
			studentService,
		)

	// Start Application

	studentController.Start()
}
