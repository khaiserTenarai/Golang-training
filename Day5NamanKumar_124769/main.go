package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	repo := NewInMemoryEmployeeRepository()
	svc := NewEmployeeService(repo)
	ctrl := NewEmployeeController(svc)

	e1 := Employee{
		Person: Person{
			FirstName: "Asha",
			LastName:  "Rao",
			Email:     "asha.rao@example.com",
		},
		Salary:   85000,
		HireDate: "2022-03-14",
		Address: Address{
			Street:  "12 MG Road",
			City:    "Bengaluru",
			State:   "Karnataka",
			ZipCode: "560001",
		},
		Department: Department{
			ID:    1,
			Name:  "Engineering",
			Floor: 4,
		},
	}

	e2 := Employee{
		Person: Person{
			FirstName: "Ravi",
			LastName:  "Kumar",
			Email:     "ravi.kumar@example.com",
		},
		Salary:   65000,
		HireDate: "2023-07-01",
		Address: Address{
			Street:  "45 Residency Road",
			City:    "Chennai",
			State:   "Tamil Nadu",
			ZipCode: "600002",
		},
		Department: Department{
			ID:    2,
			Name:  "Sales",
			Floor: 2,
		},
	}

	ctrl.CreateEmployee(e1)
	ctrl.CreateEmployee(e2)

	fmt.Println("\n--- all employees ---")
	ctrl.ListEmployees()

	fmt.Println("\n--- json marshal ---")
	ctrl.ShowEmployeeJSON(1)

	fmt.Println("\n--- json unmarshal ---")
	jsonData := []byte(`{
		"id": 3,
		"person": {"first_name": "Meera", "last_name": "Nair", "email": "meera.nair@example.com"},
		"salary": 72000,
		"hire_date": "2024-01-10",
		"address": {"street": "9 Park Street", "city": "Kolkata", "state": "West Bengal", "zip_code": "700016"},
		"department": {"id": 3, "name": "Marketing", "floor": 1},
		"active": true
	}`)
	var e3 Employee
	if err := json.Unmarshal(jsonData, &e3); err != nil {
		fmt.Println("unmarshal error:", err)
	} else {
		fmt.Println("unmarshaled:", e3)
		fmt.Println("annual salary:", e3.AnnualSalary())
	}

	fmt.Println("\n--- pointer receiver: give raise ---")
	ctrl.GiveRaise(1, 5000)

	fmt.Println("\n--- value receiver: full name and annual salary ---")
	emp, _ := svc.GetEmployee(2)
	fmt.Println("full name:", emp.FullName())
	fmt.Println("annual salary:", emp.AnnualSalary())

	fmt.Println("\n--- delete employee ---")
	ctrl.DeleteEmployee(2)

	fmt.Println("\n--- final list ---")
	ctrl.ListEmployees()
}
