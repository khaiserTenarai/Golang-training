package coverage

func CalculateSalary(salary float64, bonus float64) float64 {
	if salary < 0 {
		return 0
	}

	return salary + bonus
}
