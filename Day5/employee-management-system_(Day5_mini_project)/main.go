package main

import (
	"employee-management/controller"
	"employee-management/repository"
	"employee-management/service"
	"employee-management/view"
)

func main() {

	// Create repository
	employeeRepository := repository.NewEmployeeRepository()

	// Create service
	employeeService := service.NewEmployeeService(employeeRepository)

	// Create controller
	employeeController := controller.NewEmployeeController(employeeService)

	// Create view
	employeeView := view.NewEmployeeView(employeeController)

	// Start application
	employeeView.Start()
}


/*
PS C:\Training\Go Lang\Day_5\124772_Day5(Go)_Assignment_Reddem_Ganesh_Reddy\employee-management-system_(Day5_mini_project)> go run .

========== Employee Management ==========
1. Add Employee
2. Get All Employees
3. Get Employee By ID
4. Update Employee
5. Delete Employee
6. Exit
Enter your choice: 1

------ Add Employee ------
Enter ID: 101
Enter Name: ganesh
Enter Age: 21
Enter Salary: 2900
Employee added successfully.

========== Employee Management ==========
1. Add Employee
2. Get All Employees
3. Get Employee By ID
4. Update Employee
5. Delete Employee
6. Exit
Enter your choice: 2

========== Employees ==========
ID     : 101
Name   : ganesh
Age    : 21
Salary : 2900
-------------------------------

========== Employee Management ==========
1. Add Employee
2. Get All Employees
3. Get Employee By ID
4. Update Employee
5. Delete Employee
6. Exit
Enter your choice: 3
Enter Employee ID: 101

========== Employee ==========
ID     : 101
Name   : ganesh
Age    : 21
Salary : 2900
-------------------------------

========== Employee Management ==========
1. Add Employee
2. Get All Employees
3. Get Employee By ID
4. Update Employee
5. Delete Employee
6. Exit
Enter your choice: 4

------ Update Employee ------
Enter Employee ID: 101
Enter Name: Ramesh
Enter Age: 23
Enter Salary: 89900
Employee updated successfully.

========== Employee Management ==========
1. Add Employee
2. Get All Employees
3. Get Employee By ID
4. Update Employee
5. Delete Employee
6. Exit
Enter your choice: 2

========== Employees ==========
ID     : 101
Name   : Ramesh
Age    : 23
Salary : 89900
-------------------------------

========== Employee Management ==========
1. Add Employee
2. Get All Employees
3. Get Employee By ID
4. Update Employee
5. Delete Employee
6. Exit
Enter your choice: 5
Enter Employee ID: 101
Employee deleted successfully.

========== Employee Management ==========
1. Add Employee
2. Get All Employees
3. Get Employee By ID
4. Update Employee
5. Delete Employee
6. Exit
Enter your choice: 6
Thank you!

*/