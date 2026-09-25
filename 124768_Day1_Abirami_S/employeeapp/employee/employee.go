package employee

import (
	"fmt"

	"github.com/google/uuid"
)

// task5
// Intro Message
func DisplayMessage() {
	fmt.Println("Hi I am an employee")

	// task6
	// Use of External Package
	id := uuid.New()
	fmt.Println("Employee ID: ", id)
}
