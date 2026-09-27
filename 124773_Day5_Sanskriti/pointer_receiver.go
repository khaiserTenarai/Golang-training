package main
import "fmt"

type Emp struct{
	ID int
	Name string
	Email string
	Age int
	Salary float64
}

func (e *Emp) ChangeName(name string){
	e.Name = name
}

func main(){
	emp := Emp{
		ID: 1,
		Name: "Sans",
		Email: "sans@gmail.com",
		Age: 22,
		Salary: 33000.0,
	}

	fmt.Println("Before changing name", emp.Name)
	emp.ChangeName("Amit")
	fmt.Println("After changing name", emp.Name)
}