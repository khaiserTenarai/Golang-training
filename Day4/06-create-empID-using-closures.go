package main


import "fmt"


//closure
func EmployeeIDGenerator()(func() (id int)){

	id := 100
	return func() int{
		id++
		return id
	}

}
func main(){

	nextEmpId := EmployeeIDGenerator()
	
	
	fmt.Println(nextEmpId())
	fmt.Println(nextEmpId())
	fmt.Println(nextEmpId())
	fmt.Println(nextEmpId())
	
}