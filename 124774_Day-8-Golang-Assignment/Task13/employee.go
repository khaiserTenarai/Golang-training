package benchmark

type Employee struct {
	Name   string
	Salary float64
}

func CalculateAnnualSalary(employee Employee) float64 {
	return employee.Salary * 12
}

func CalculateBonus(employee Employee) float64 {
	return employee.Salary * 0.10
}
