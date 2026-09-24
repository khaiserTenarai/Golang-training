// Command employeeapp demonstrates calling a local package from main.
package main

import (
	"fmt"

	"employeeapp/employee" // <module name>/<package folder>
)

func main() {
	e := employee.New(1, "Asha Rao", 85000)
	fmt.Println(e.Describe())
}
