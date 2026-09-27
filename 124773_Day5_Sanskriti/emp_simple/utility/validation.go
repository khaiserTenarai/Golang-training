package utility


import(
	"errors"
	//"strings"
)

func ValidateAge(age int) error{
	if age<= 0 || age >60{
		return errors.New("Invalid Age")

	}

	return nil
}

func ValidateEntry(value string) error{
	if value == ""{
		return errors.New("Value should not be empty")
	}
	return nil
}

func ValidateSalary(salary float64) error{
	if salary <=1 {
		return errors.New("Invalid Salary")
	}
	return nil
}