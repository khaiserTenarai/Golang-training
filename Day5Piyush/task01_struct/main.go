package main
import "fmt"

type Employee struct{
	Id int
	Name string
	Email string
	Salary float64
	Age int
	Phone int
	Department string
	City string
}

func main(){
	employee := Employee{
		Id : 101,
		Name : "Piyush",
		Email : "piyush@gmail.com",
		Salary : 25000.5,
		Age : 21,
		Phone: 7973701413,
		Department : "IT",
		City : "Bangalore",
	}

	fmt.Println("Id: ", employee.Id)
	fmt.Println("Name: ", employee.Name)
	fmt.Println("Email: ", employee.Email)
	fmt.Println("Salary: ", employee.Salary)
	fmt.Println("Age: ", employee.Age)
	fmt.Println("Phone: ", employee.Phone)
	fmt.Println("Department: ", employee.Department)
	fmt.Println("City: ", employee.City)
}

