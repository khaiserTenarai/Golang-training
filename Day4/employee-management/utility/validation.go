package utility

func IsValidName(name string) bool {
	return name != ""
}

func IsValidAge(age int) bool {
	return age >= 18 && age <= 60
}

func IsValidSalary(salary float64) bool {
	return salary > 0
}

func IsValidPosition(position string) bool {
	return position != ""
}
