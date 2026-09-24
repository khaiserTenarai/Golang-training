// Command packagedesign wires the employee and utils packages together.
package main

import (
	"fmt"
	"log"

	"packagedesign/employee"
)

func main() {
	e, err := employee.New(101, "asha rao", 1250000)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Before raise:", e)

	e.GiveRaise(10)
	fmt.Println("After 10% raise:", e)
}
