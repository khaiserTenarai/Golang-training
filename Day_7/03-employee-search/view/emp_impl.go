package view
import (
	"fmt"

	"employee-management/model"
)

type EmployeeViewImpl struct {
}

func NewEmployeeView() EmployeeView {
	return &EmployeeViewImpl{}
}

func (v *EmployeeViewImpl) ShowMenu() int {

	fmt.Println()
	fmt.Println("========== Employee Management ==========")
	fmt.Println("1. Save Employee")
	fmt.Println("2. Find Employee")
	fmt.Println("3. Find All Employees")
	fmt.Println("4. Update Employee")
	fmt.Println("5. Delete Employee")
	fmt.Println("6. Search Employees")
	fmt.Println("7. Exit")

	var choice int

	fmt.Print("Enter choice: ")
	fmt.Scan(&choice)

	return choice
}

func (v *EmployeeViewImpl) ReadEmployee() model.Employee {

	var employee model.Employee

	fmt.Println()
	fmt.Println("---------- Enter Employee ----------")

	

	fmt.Print("Enter Name: ")
	fmt.Scan(&employee.Name)

	fmt.Print("Enter Age: ")
	fmt.Scan(&employee.Age)

	fmt.Print("Enter Email: ")
	fmt.Scan(&employee.Email)

	fmt.Print("Enter Salary: ")
	fmt.Scan(&employee.Salary)

	fmt.Print("Enter Department ID: ")
	fmt.Scan(&employee.Department.ID)

	return employee
}

func (v *EmployeeViewImpl) ReadID() int {

	var id int

	fmt.Print("Enter ID: ")
	fmt.Scan(&id)

	return id
}

func (v *EmployeeViewImpl) DisplayEmployee(
	employee model.Employee,
) {

	fmt.Println()
	fmt.Println("---------- Employee ----------")

	fmt.Println("ID:", employee.ID)
	fmt.Println("Name:", employee.Name)
	fmt.Println("Age:", employee.Age)
	fmt.Println("Email:", employee.Email)
	fmt.Println("Salary:", employee.Salary)
	fmt.Println("Department ID:", employee.Department.ID)
	fmt.Println("Department:", employee.Department.Name)
}

func (v *EmployeeViewImpl) DisplayEmployees(
	employees []model.Employee,
) {

	fmt.Println()
	fmt.Println("---------- Employees ----------")

	if len(employees) == 0 {
		fmt.Println("No employees found.")
		return
	}

	for _, employee := range employees {

		fmt.Printf(
			"ID: %d | Name: %s | Age: %d | Email: %s | Salary: %.2f | Department: %s\n",
			employee.ID,
			employee.Name,
			employee.Age,
			employee.Email,
			employee.Salary,
			employee.Department.Name,
		)
	}
}

func (v *EmployeeViewImpl) ReadSearch() model.EmployeeSearch {

	var search model.EmployeeSearch

	fmt.Println()
	fmt.Println("---------- Employee Search ----------")

	fmt.Print("Enter Name (- for all): ")
	fmt.Scan(&search.Name)

	if search.Name == "-" {
		search.Name = ""
	}

	fmt.Print("Enter Department ID (0 for all): ")
	fmt.Scan(&search.DepartmentID)

	fmt.Print("Enter Minimum Salary (0 for all): ")
	fmt.Scan(&search.MinSalary)

	fmt.Print("Enter Maximum Salary (0 for all): ")
	fmt.Scan(&search.MaxSalary)

	fmt.Print("Enter Page Number: ")
	fmt.Scan(&search.Page)

	fmt.Print("Enter Page Size: ")
	fmt.Scan(&search.Size)

	fmt.Println()
	fmt.Println("Sort By")
	fmt.Println("1. ID")
	fmt.Println("2. Name")
	fmt.Println("3. Department")
	fmt.Println("4. Salary")

	var choice int

	fmt.Print("Enter choice: ")
	fmt.Scan(&choice)

	switch choice {

	case 2:
		search.SortBy = "name"

	case 3:
		search.SortBy = "department"

	case 4:
		search.SortBy = "salary"

	default:
		search.SortBy = "id"
	}

	fmt.Println()
	fmt.Println("Sort Order")
	fmt.Println("1. Ascending")
	fmt.Println("2. Descending")

	fmt.Print("Enter choice: ")
	fmt.Scan(&choice)

	if choice == 2 {
		search.SortOrder = "desc"
	} else {
		search.SortOrder = "asc"
	}

	return search
}

func (v *EmployeeViewImpl) DisplaySearchResult(
	employees []model.Employee,
	total int,
	search model.EmployeeSearch,
) {

	fmt.Println()
	fmt.Println("========== Search Result ==========")

	if len(employees) == 0 {
		fmt.Println("No employees found.")
		return
	}

	for _, employee := range employees {

		fmt.Printf(
			"ID: %d | Name: %s | Age: %d | Email: %s | Salary: %.2f | Department: %s\n",
			employee.ID,
			employee.Name,
			employee.Age,
			employee.Email,
			employee.Salary,
			employee.Department.Name,
		)
	}

	totalPages := (total + search.Size - 1) / search.Size

	fmt.Println()
	fmt.Println("Total Employees:", total)
	fmt.Println("Current Page:", search.Page)
	fmt.Println("Page Size:", search.Size)
	fmt.Println("Total Pages:", totalPages)
}