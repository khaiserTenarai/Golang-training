package controller

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"example.com/employee-management/models"
	"example.com/employee-management/repository"
	"example.com/employee-management/service"
)

type AppController struct {
	empService  service.EmployeeService
	deptService service.DepartmentService
	reader      *bufio.Reader
}

func NewAppController(
	empService service.EmployeeService,
	deptService service.DepartmentService,
) *AppController {
	return &AppController{
		empService:  empService,
		deptService: deptService,
		reader:      bufio.NewReader(os.Stdin),
	}
}

func (c *AppController) Start() {
	for {
		c.showMenu()
		choice := c.readInt("Enter your choice: ")

		switch choice {
		// Employee Management
		case 1:
			c.addEmployee()
		case 2:
			c.getEmployee()
		case 3:
			c.getAllEmployees()
		case 4:
			c.updateEmployee()
		case 5:
			c.deleteEmployee()

		// Department Management & Mapping
		case 6:
			c.addDepartment()
		case 7:
			c.getAllDepartments()
		case 8:
			c.assignDepartmentToEmployee()
		case 9:
			c.getEmployeeWithDepartment()

		case 10:
			fmt.Println("\nThank you for using Employee & Department Management System.")
			return

		default:
			fmt.Println("Invalid choice. Please select 1 to 10.")
		}

		c.pause()
	}
}

func (c *AppController) showMenu() {
	fmt.Println()
	fmt.Println("==============================================")
	fmt.Println("   EMPLOYEE & DEPARTMENT MANAGEMENT SYSTEM")
	fmt.Println("==============================================")
	fmt.Println(" 1. Add Employee")
	fmt.Println(" 2. Find Employee")
	fmt.Println(" 3. Find All Employees")
	fmt.Println(" 4. Update Employee")
	fmt.Println(" 5. Delete Employee")
	fmt.Println("----------------------------------------------")
	fmt.Println(" 6. Add Department")
	fmt.Println(" 7. Find All Departments")
	fmt.Println(" 8. Assign Department to Employee")
	fmt.Println(" 9. View Employee with Department Details")
	fmt.Println("----------------------------------------------")
	fmt.Println("10. Exit")
	fmt.Println("==============================================")
}

// ==================================================
// EMPLOYEE CONTROLLER METHODS
// ==================================================

func (c *AppController) addEmployee() {
	fmt.Println("\n---------- ADD EMPLOYEE ----------")

	employee := models.Employee{}
	employee.Name = c.readString("Enter Name: ")
	employee.Email = c.readString("Enter Email: ")
	employee.Age = c.readInt("Enter Age: ")
	employee.Salary = c.readFloat("Enter Salary: ")
	employee.City = c.readString("Enter City: ")
	employee.State = c.readString("Enter State: ")
	employee.Pincode = c.readString("Enter Pincode: ")

	err := c.empService.AddEmployee(employee)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Employee added successfully.")
}

func (c *AppController) getEmployee() {
	fmt.Println("\n---------- FIND EMPLOYEE ----------")

	id := c.readInt64("Enter Employee ID: ")
	employee, err := c.empService.GetEmployee(id)

	if errors.Is(err, repository.ErrNotFound) {
		fmt.Println("Employee not found.")
		return
	}
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("\nID: %d | Name: %s | Email: %s | Age: %d | Salary: %.2f | City: %s | State: %s\n",
		employee.ID, employee.Name, employee.Email, employee.Age, employee.Salary, employee.City, employee.State)
}

func (c *AppController) getAllEmployees() {
	fmt.Println("\n---------- ALL EMPLOYEES ----------")

	employees, err := c.empService.GetAllEmployees()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(employees) == 0 {
		fmt.Println("No employees found.")
		return
	}

	fmt.Printf("\n%-5s %-20s %-28s %-5s %-12s %-18s\n", "ID", "NAME", "EMAIL", "AGE", "SALARY", "CITY")
	fmt.Println(strings.Repeat("-", 95))

	for _, emp := range employees {
		fmt.Printf("%-5d %-20s %-28s %-5d %-12.2f %-18s\n",
			emp.ID, emp.Name, emp.Email, emp.Age, emp.Salary, emp.City)
	}
}

func (c *AppController) updateEmployee() {
	fmt.Println("\n---------- UPDATE EMPLOYEE ----------")

	id := c.readInt64("Enter Employee ID: ")
	employee, err := c.empService.GetEmployee(id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Enter new values:")
	employee.Name = c.readString("Enter Name: ")
	employee.Email = c.readString("Enter Email: ")
	employee.Age = c.readInt("Enter Age: ")
	employee.Salary = c.readFloat("Enter Salary: ")
	employee.City = c.readString("Enter City: ")
	employee.State = c.readString("Enter State: ")
	employee.Pincode = c.readString("Enter Pincode: ")

	err = c.empService.UpdateEmployee(employee)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Employee updated successfully.")
}

func (c *AppController) deleteEmployee() {
	fmt.Println("\n---------- DELETE EMPLOYEE ----------")

	id := c.readInt64("Enter Employee ID: ")
	confirmation := c.readString("Are you sure you want to delete? (y/n): ")

	if strings.ToLower(confirmation) == "y" || strings.ToLower(confirmation) == "yes" {
		err := c.empService.DeleteEmployee(id)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Employee deleted successfully.")
	} else {
		fmt.Println("Delete operation cancelled.")
	}
}

// ==================================================
// DEPARTMENT & MAPPING CONTROLLER METHODS
// ==================================================

func (c *AppController) addDepartment() {
	fmt.Println("\n---------- ADD DEPARTMENT ----------")

	dept := models.Department{}
	dept.Name = c.readString("Enter Department Name: ")
	dept.Code = c.readString("Enter Department Code (e.g. ENG, HR): ")

	err := c.deptService.AddDepartment(dept)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Department added successfully.")
}

func (c *AppController) getAllDepartments() {
	fmt.Println("\n---------- ALL DEPARTMENTS ----------")

	depts, err := c.deptService.GetAllDepartments()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(depts) == 0 {
		fmt.Println("No departments found.")
		return
	}

	fmt.Printf("\n%-5s %-25s %-10s\n", "ID", "NAME", "CODE")
	fmt.Println(strings.Repeat("-", 45))

	for _, d := range depts {
		fmt.Printf("%-5d %-25s %-10s\n", d.ID, d.Name, d.Code)
	}
}

func (c *AppController) assignDepartmentToEmployee() {
	fmt.Println("\n---------- ASSIGN DEPARTMENT ----------")

	empID := c.readInt64("Enter Employee ID: ")
	deptID := c.readInt("Enter Department ID: ")

	err := c.empService.AssignDepartment(empID, deptID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Department assigned to employee successfully.")
}

func (c *AppController) getEmployeeWithDepartment() {
	fmt.Println("\n---------- VIEW EMPLOYEE WITH DEPARTMENT ----------")

	id := c.readInt64("Enter Employee ID: ")
	empWithDept, err := c.empService.GetEmployeeWithDepartment(id)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("\nID         : %d\n", empWithDept.ID)
	fmt.Printf("Name       : %s\n", empWithDept.Name)
	fmt.Printf("Email      : %s\n", empWithDept.Email)
	fmt.Printf("Department : %s (%s)\n", empWithDept.DepartmentName, empWithDept.DepartmentCode)
}

// ==================================================
// INPUT HELPERS
// ==================================================

func (c *AppController) readString(message string) string {
	for {
		fmt.Print(message)
		value, err := c.reader.ReadString('\n')
		if err != nil {
			continue
		}
		value = strings.TrimSpace(value)
		if value != "" {
			return value
		}
		fmt.Println("Value cannot be empty.")
	}
}

func (c *AppController) readInt(message string) int {
	for {
		value := c.readString(message)
		number, err := strconv.Atoi(value)
		if err == nil {
			return number
		}
		fmt.Println("Please enter a valid integer.")
	}
}

func (c *AppController) readInt64(message string) int64 {
	for {
		value := c.readString(message)
		number, err := strconv.ParseInt(value, 10, 64)
		if err == nil {
			return number
		}
		fmt.Println("Please enter a valid number.")
	}
}

func (c *AppController) readFloat(message string) float64 {
	for {
		value := c.readString(message)
		number, err := strconv.ParseFloat(value, 64)
		if err == nil {
			return number
		}
		fmt.Println("Please enter a valid decimal number.")
	}
}

func (c *AppController) pause() {
	fmt.Println()
	fmt.Print("Press ENTER to continue...")
	_, _ = c.reader.ReadString('\n')
}