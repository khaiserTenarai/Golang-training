package main

func CalculateSalary(hourlyRate float64, hoursWorked float64) float64 {
	if hoursWorked <= 0 || hourlyRate <= 0 {
		return 0
	}
	return hourlyRate * hoursWorked
}