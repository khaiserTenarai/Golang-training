package main
import "fmt"

type Address struct{
	City string
	State string
	Pincode int
}

type Department struct{
	Id int
	Name string
	Manager string
}
type Employee struct {
	Id int
	Name string
	Age int
	Salary float64
	Address Address
	Department Department
}

func main(){
	employee := Employee{
		Id: 101,
		Name: "Piyush",
		Age: 21,
		Salary: 25000,
		Address : Address{
			City : "Bangalore",
			State : "Karnataka",
			Pincode  : 560037,
		},

		Department : Department{
			Id : 10,
			Name: "IT",
			Manager: "Khaiser",
		},
	}

	fmt.Println("Employee:", employee.Name)
	fmt.Println("City:", employee.Address.City)
	fmt.Println("State:", employee.Address.State)
	fmt.Println("Department:", employee.Department.Name)
	fmt.Println("Manager:", employee.Department.Manager)
}