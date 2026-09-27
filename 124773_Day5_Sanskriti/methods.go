package main
import "fmt"

type Emp struct{
	ID int
	Name string
	Email string
	Age int
	Salary float64
}

func (e Emp) Display(){
	fmt.Println("Emp Name: ",e.Name)
	fmt.Println("Age",e.Age)
}

func main(){
	emp := Emp{
		ID: 1,
		Name: "Sans",
		Email: "sans@gmail.com",
		Age: 22,
		Salary: 33000.0,
	}

	emp.Display()
}
