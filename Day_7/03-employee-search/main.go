package main

import (
	"fmt"
	"os"

	"employee-management/controller"
	"employee-management/database"
	"employee-management/repository"
	"employee-management/service"
	"employee-management/view"
)

func main() {

	db, err := database.Connect()

	if err != nil {
		fmt.Println("Database connection failed:", err)
		os.Exit(1)
	}

	defer db.Close()

	// Department
	departmentRepository :=
		repository.NewDepartmentRepository(db)

	departmentService :=
		service.NewDepartmentService(
			departmentRepository,
		)

	departmentView :=
		view.NewDepartmentView()

	departmentController :=
		controller.NewDepartmentController(
			departmentView,
			departmentService,
		)

	// Employee
	employeeRepository :=
		repository.NewEmployeeRepository(db)

	employeeService :=
		service.NewEmployeeService(
			employeeRepository,
		)

	employeeView :=
		view.NewEmployeeView()

	employeeController :=
		controller.NewEmployeeController(
			employeeView,
			employeeService,
		)

	// Main menu
	for {

		fmt.Println()
		fmt.Println("======================================")
		fmt.Println("       Employee Management System")
		fmt.Println("======================================")
		fmt.Println("1. Department Management")
		fmt.Println("2. Employee Management")
		fmt.Println("3. Exit")

		var choice int

		fmt.Print("Enter choice: ")
		fmt.Scan(&choice)

		switch choice {

		case 1:
			departmentController.Start()

		case 2:
			employeeController.Start()

		case 3:
			fmt.Println("Thank you..")
			return

		default:
			fmt.Println("Invalid choice.")
		}
	}
}

/*

1. go.mod
module employee-management

go 1.24

require github.com/jackc/pgx/v5 v5.7.2

2. Database
Create the database:

CREATE DATABASE employee_management_db;

Connect to it:

\c employee_management_db

Department table
CREATE TABLE departments (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE
);

Employee table
CREATE TABLE employees (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    age INT NOT NULL,
    email VARCHAR(150) NOT NULL UNIQUE,
    salary NUMERIC(12,2) NOT NULL,
    department_id INT NOT NULL,

    CONSTRAINT fk_employee_department
        FOREIGN KEY (department_id)
        REFERENCES departments(id)
);

Sample departments
INSERT INTO departments (name)
VALUES
('IT'),
('HR'),
('Finance'),
('Marketing');

Sample employees
INSERT INTO employees
(name, age, email, salary, department_id)
VALUES
('Ganesh', 21, 'ganesh@gmail.com', 60000, 1),
('Rahul', 25, 'rahul@gmail.com', 45000, 2),
('Priya', 28, 'priya@gmail.com', 75000, 1),
('Amit', 30, 'amit@gmail.com', 55000, 3),
('Sneha', 24, 'sneha@gmail.com', 50000, 2),
('Kiran', 29, 'kiran@gmail.com', 90000, 1),
('Arjun', 31, 'arjun@gmail.com', 65000, 3),
('Neha', 26, 'neha@gmail.com', 48000, 4),
('Ravi', 27, 'ravi@gmail.com', 70000, 1),
('Pooja', 23, 'pooja@gmail.com', 52000, 2);

*/
