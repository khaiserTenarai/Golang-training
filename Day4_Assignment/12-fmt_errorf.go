package main
import (
	"errors"
	"fmt"
)

var ErrEmployeeNotFound = errors.New("employee not found")

func findEmployee(id int) error {

	if id != 1 {

		return fmt.Errorf(
			"employee ID %d: %w",
			id,
			ErrEmployeeNotFound,
		)
	}

	return nil
}

func main() {

	err := findEmployee(10)

	if err != nil {
		fmt.Println(err)
	}
}