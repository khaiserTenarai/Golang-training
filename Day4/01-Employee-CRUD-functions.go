package main

import "fmt"

type Employee struct {
		ID int
		Name string
		Salary float64
	}


	var employees []Employee


	//create Employee function
	func CreateEmployee(){
		var emp Employee

		fmt.Println("Enter ID : ")
		fmt.Scan(&emp.ID)

		fmt.Println("Enter Name : ")
		fmt.Scan(&emp.Name)

		fmt.Println("Enter Salary :  ")
		fmt.Scan(&emp.Salary)

		employees = append(employees, emp)

		fmt.Println("Employee created..")
	}


	// read employees 
	func ReadEmployee(){

		if len(employees) == 0{
			fmt.Println("No employees found !..")
			return 
		}
		for _, emp := range employees{
			fmt.Println(emp)
		}
	}

	// Update employee
	func UpdateEmployee(){

		var id int 
		fmt.Println("Enter ID to update : ")
		fmt.Scan(&id)


		for ind := range employees{
			if employees[ind].ID == id{

				fmt.Println("Enter a new name: ")
				fmt.Scan(&employees[ind].Name)

				fmt.Println("Enter new Salary : ")
				fmt.Scan(&employees[ind].Salary)

				fmt.Println("Employee data updated..")
				return
			}
		}
		fmt.Println("Employee not found..")
	}

	//Delete employee 
	func DeleteEmployee(){

		var id int 

		fmt.Println("Enter ID to delete : ")
		fmt.Scan(&id)

		for ind, emp := range employees{
			if emp.ID  == id{
				employees = append(employees[:ind], employees[ind+1:]...)
				fmt.Println("Employee data deleted..")
				return
			}
		}
		fmt.Println("Employee data not found ..")
	}
func main(){

	//continue until user exits

	for {
		fmt.Println("============ Employee CRUD =============")

		fmt.Println("1.CREATE")
		fmt.Println("2.READ")
		fmt.Println("3.UPDATE")
		fmt.Println("4.DELETE")
		fmt.Println("5.exit")

		var choice int 

		fmt.Println("Enter Choice : ")
		fmt.Scan(&choice)


		switch choice{
		case 1: CreateEmployee()
		case 2:ReadEmployee()
		case 3:UpdateEmployee()
		case 4:DeleteEmployee()
		case 5: fmt.Println("Program ended..")
		return
		default :
		fmt.Println("Invalid choice ..Try again..")
		}
	}



	
}