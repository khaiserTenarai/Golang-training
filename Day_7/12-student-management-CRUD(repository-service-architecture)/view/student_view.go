package view
import (
	"fmt"

	"student_app/model"
)

type StudentView interface {
	ShowMenu() int
	ReadStudent() model.Student
	ReadStudentForUpdate() model.Student
	ReadID() int
	DisplayStudent(student model.Student)
	DisplayStudents(students []model.Student)
}

type StudentViewImpl struct {
}

func NewStudentView() StudentView {
	return &StudentViewImpl{}
}

func (v *StudentViewImpl) ShowMenu() int {

	fmt.Println("\n========== Student Management ==========")
	fmt.Println("1. Save Student")
	fmt.Println("2. Find Student")
	fmt.Println("3. Find All Students")
	fmt.Println("4. Update Student")
	fmt.Println("5. Delete Student")
	fmt.Println("6. Exit")

	var choice int

	fmt.Print("Enter choice: ")
	fmt.Scan(&choice)

	return choice
}

func (v *StudentViewImpl) ReadStudent() model.Student {

	var student model.Student

	fmt.Println("\n---------- Add Student ----------")

	fmt.Print("Enter Name: ")
	fmt.Scan(&student.Name)

	fmt.Print("Enter Age: ")
	fmt.Scan(&student.Age)

	fmt.Print("Enter Email: ")
	fmt.Scan(&student.Email)

	return student
}

func (v *StudentViewImpl) ReadStudentForUpdate() model.Student {

	var student model.Student

	fmt.Println("\n---------- Update Student ----------")

	fmt.Print("Enter ID: ")
	fmt.Scan(&student.ID)

	fmt.Print("Enter Name: ")
	fmt.Scan(&student.Name)

	fmt.Print("Enter Age: ")
	fmt.Scan(&student.Age)

	fmt.Print("Enter Email: ")
	fmt.Scan(&student.Email)

	return student
}

func (v *StudentViewImpl) ReadID() int {

	var id int

	fmt.Print("Enter Student ID: ")
	fmt.Scan(&id)

	return id
}

func (v *StudentViewImpl) DisplayStudent(
	student model.Student,
) {

	fmt.Println("\n---------- Student ----------")

	fmt.Println("ID    :", student.ID)
	fmt.Println("Name  :", student.Name)
	fmt.Println("Age   :", student.Age)
	fmt.Println("Email :", student.Email)
}

func (v *StudentViewImpl) DisplayStudents(
	students []model.Student,
) {

	if len(students) == 0 {
		fmt.Println("No students found.")
		return
	}

	fmt.Println("\n---------- Students ----------")

	for _, student := range students {

		fmt.Println(
			"ID:",
			student.ID,
			"| Name:",
			student.Name,
			"| Age:",
			student.Age,
			"| Email:",
			student.Email,
		)
	}
}
