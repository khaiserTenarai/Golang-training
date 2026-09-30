package controller
import (
	"fmt"

	"student_app/service"
	"student_app/view"
)

type StudentController interface {
	Start()
	Process(choice int)
}

type StudentControllerImpl struct {
	view    view.StudentView
	service service.StudentService
}

func NewStudentController(
	view view.StudentView,
	service service.StudentService,
) StudentController {

	return &StudentControllerImpl{
		view:    view,
		service: service,
	}
}

func (c *StudentControllerImpl) Start() {

	for {

		choice := c.view.ShowMenu()

		if choice == 6 {

			fmt.Println("Thank you..")

			return
		}

		c.Process(choice)
	}
}

func (c *StudentControllerImpl) Process(
	choice int,
) {

	switch choice {

	case 1:

		student := c.view.ReadStudent()

		err := c.service.Save(student)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Student saved successfully.")

	case 2:

		id := c.view.ReadID()

		student, err :=
			c.service.FindByID(id)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		c.view.DisplayStudent(student)

	case 3:

		students, err :=
			c.service.FindAll()

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		c.view.DisplayStudents(students)

	case 4:

		student :=
			c.view.ReadStudentForUpdate()

		err :=
			c.service.Update(student)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Student updated successfully.")

	case 5:

		id := c.view.ReadID()

		err := c.service.Delete(id)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Student deleted successfully.")

	default:

		fmt.Println("Invalid choice.")
	}
}
