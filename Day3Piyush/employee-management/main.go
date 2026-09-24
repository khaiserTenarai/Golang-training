package main

func main() {
	employees := []Employee{
		{
			ID:         1,
			Name:       "Piyush",
			Salary:     50000,
			Department: "IT",
		},
		{
			ID:         2,
			Name:       "Rahul",
			Salary:     60000,
			Department: "HR",
		},
	}

	for _, employee := range employees {
		employee.Display()
	}

	GenerateReport(employees)
}