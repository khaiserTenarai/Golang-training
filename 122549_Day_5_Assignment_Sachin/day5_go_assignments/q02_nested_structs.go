// 2. Create nested Address and Department structs.

package main

import "fmt"

type Address struct {
	City  string
	State string
}

type Department struct {
	Name string
	Code string
}

type Employee struct {
	ID       int
	Name     string
	Dept     Department
	HomeAddr Address
}

func main() {
	emp := Employee{
		ID:   1,
		Name: "sachin",
		Dept: Department{
			Name: "software engineering",
			Code: "ENG-11",
		},
		HomeAddr: Address{
			City:  "blr",
			State: "kar",
		},
	}

	fmt.Println("Employee:", emp.Name)
	fmt.Println("Department:", emp.Dept.Name)
	fmt.Println("City:", emp.HomeAddr.City)
}
