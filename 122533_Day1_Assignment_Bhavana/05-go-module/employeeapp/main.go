// 5. Go Module
//
// Create a module named employeeapp.
// Create a separate package and call it from main.go.

package main

import (
	"employeeapp/greeting"
	"fmt"
)

func main() {
	fmt.Println(greeting.Hello("Team"))
}
