package pkg

import (
	"fmt"

	"12-project-structure-simple-student-app/internal"
)

func PrintStudent(s internal.Student) {
	fmt.Println("Name:", s.Name)
	fmt.Println("Age:", s.Age)
}
