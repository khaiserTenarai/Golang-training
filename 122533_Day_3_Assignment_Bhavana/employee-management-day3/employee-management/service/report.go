package service

import "fmt"

func (s *EmployeeService) GenerateSummary() {
	if len(s.employees) == 0 {
		fmt.Println("No employees to summarize.")
		return
	}

	var totalSalary float64
	for _, e := range s.employees {
		totalSalary += e.Salary
	}
	average := totalSalary / float64(len(s.employees))

	fmt.Println("Employee count :", len(s.employees))
	fmt.Printf("Average salary : %.2f\n", average)
}
