package main
import (
	"errors"
	"fmt"
)
func validateSalary(salary float64)error{
	if salary<=0{
		return errors.New("salary must be greater than 0")
	}
	return nil
}

func main(){
	salary :=0.0
	err :=validateSalary(salary)
	if err!=nil{
		fmt.Println("Validation Error:",err)
	}else{
		fmt.Println("Salary is valid")
	}
}