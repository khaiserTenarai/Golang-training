package main
import (
	"errors"
	"fmt"
)

var ErrEmployeeNotFound = errors.New("employee not found")
func findEmployee(id int) error {

	if id != 1 {
		return fmt.Errorf(
			"ID %d: %w",
			id,
			ErrEmployeeNotFound,
		)
	}

	return nil
}

func main() {
	err := findEmployee(10)

	if err != nil {

		fmt.Println("Error:", err)

		if errors.Is(err, ErrEmployeeNotFound) {
			fmt.Println("Employee not found")
		}
	}
}