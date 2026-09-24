// Command godocdemo shows a documented package in use.
package main

import (
	"errors"
	"fmt"

	"godocdemo/employee"
)

func main() {
	s := employee.NewStore()
	s.Add("Asha Rao", "Engineering", 85000)
	if _, err := s.Get(42); errors.Is(err, employee.ErrNotFound) {
		fmt.Println("ID 42:", err)
	}
	fmt.Println("Total employees:", s.Count())
}
