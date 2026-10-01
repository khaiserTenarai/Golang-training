package controller

import (
    "bufio"
    "errors"
    "fmt"
    "os"
    "strconv"
    "strings"

    "example.com/employee-management/model"
    "example.com/employee-management/repository"
    "example.com/employee-management/service"
)

type EmployeeController struct {
    service service.EmployeeService
    reader  *bufio.Reader
}

func NewEmployeeController(
    service service.EmployeeService,
) *EmployeeController {

    return &EmployeeController{
        service: service,
        reader:  bufio.NewReader(os.Stdin),
    }
}

func (c *EmployeeController) Start() {
    for {
        c.showMenu()
        choice := c.readInt("Enter your choice: ")
        switch choice {
        case 1:
            c.addEmployee()
        case 2:
            c.getEmployee()
        case 3:
            c.getAllEmployees()
        case 4:
            c.updateEmployee()
        case 5:
            c.updateSalary()
        case 6:
            c.getSalaryHistory()
        case 7:
            c.deleteEmployee()
        case 8:
            fmt.Println("\nThank you for using Employee Management System.")
            return
        default:
            fmt.Println("Invalid choice. Please select 1 to 8.")
        }
        c.pause()
    }
}

func (c *EmployeeController) showMenu() {
    fmt.Println()
    fmt.Println("==============================================")
    fmt.Println("       EMPLOYEE MANAGEMENT SYSTEM")
    fmt.Println("==============================================")
    fmt.Println("1. Add Employee")
    fmt.Println("2. Find Employee")
    fmt.Println("3. Find All Employees")
    fmt.Println("4. Update Employee")
    fmt.Println("5. Update Salary")
    fmt.Println("6. View Salary History")
    fmt.Println("7. Delete Employee")
    fmt.Println("8. Exit")
    fmt.Println("==============================================")
}

func (c *EmployeeController) addEmployee() {

    fmt.Println()
    fmt.Println("---------- ADD EMPLOYEE ----------")

    employee := model.Employee{}

    employee.Name = c.readString("Enter Name: ")
    employee.Email = c.readString("Enter Email: ")
    employee.Age = c.readInt("Enter Age: ")
    employee.Salary = c.readFloat("Enter Salary: ")

    employee.Address.City =
        c.readString("Enter City: ")

    employee.Address.State =
        c.readString("Enter State: ")

    employee.Address.Pincode =
        c.readString("Enter Pincode: ")

    err := c.service.AddEmployee(employee)

    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println()
    fmt.Println("Employee added successfully.")
}

func (c *EmployeeController) getEmployee() {

    fmt.Println()
    fmt.Println("---------- FIND EMPLOYEE ----------")

    id := c.readInt64("Enter Employee ID: ")

    employee, err := c.service.GetEmployee(id)

    if errors.Is(err, repository.ErrNotFound) {
        fmt.Println("Employee not found.")
        return
    }

    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    employee.Display()
}

func (c *EmployeeController) getAllEmployees() {

    fmt.Println()
    fmt.Println("---------- ALL EMPLOYEES ----------")

    employees, err := c.service.GetAllEmployees()

    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    if len(employees) == 0 {
        fmt.Println("No employees found.")
        return
    }

    fmt.Println()
    fmt.Printf(
        "%-5s %-20s %-28s %-5s %-12s %-18s\n",
        "ID",
        "NAME",
        "EMAIL",
        "AGE",
        "SALARY",
        "CITY",
    )

    fmt.Println(
        strings.Repeat("-", 95),
    )

    for _, employee := range employees {

        fmt.Printf(
            "%-5d %-20s %-28s %-5d %-12.2f %-18s\n",
            employee.ID,
            employee.Name,
            employee.Email,
            employee.Age,
            employee.Salary,
            employee.Address.City,
        )
    }
}

func (c *EmployeeController) updateEmployee() {

    fmt.Println()
    fmt.Println("---------- UPDATE EMPLOYEE ----------")

    id := c.readInt64("Enter Employee ID: ")

    employee, err := c.service.GetEmployee(id)

    if errors.Is(err, repository.ErrNotFound) {
        fmt.Println("Employee not found.")
        return
    }

    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println()
    fmt.Println("Current employee:")
    employee.Display()

    fmt.Println()
    fmt.Println("Enter new values.")

    employee.Name = c.readString("Enter Name: ")
    employee.Email = c.readString("Enter Email: ")
    employee.Age = c.readInt("Enter Age: ")
    employee.Salary = c.readFloat("Enter Salary: ")

    employee.Address.City =
        c.readString("Enter City: ")

    employee.Address.State =
        c.readString("Enter State: ")

    employee.Address.Pincode =
        c.readString("Enter Pincode: ")

    err = c.service.UpdateEmployee(employee)

    if errors.Is(err, repository.ErrNotFound) {
        fmt.Println("Employee not found.")
        return
    }

    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Employee updated successfully.")
}

func (c *EmployeeController) deleteEmployee() {

    fmt.Println()
    fmt.Println("---------- DELETE EMPLOYEE ----------")

    id := c.readInt64("Enter Employee ID: ")

    employee, err := c.service.GetEmployee(id)

    if errors.Is(err, repository.ErrNotFound) {
        fmt.Println("Employee not found.")
        return
    }

    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    employee.Display()

    confirmation := c.readString(
        "Are you sure you want to delete? (y/n): ",
    )

    switch strings.ToLower(confirmation) {

    case "y", "yes":

        err = c.service.DeleteEmployee(id)

        if err != nil {
            fmt.Println("Error:", err)
            return
        }

        fmt.Println("Employee deleted successfully.")

    case "n", "no":
        fmt.Println("Delete operation cancelled.")

    default:
        fmt.Println("Invalid response. Delete operation cancelled.")
    }
}

func (c *EmployeeController) updateSalary() {
    fmt.Println("\n---------- UPDATE SALARY ----------")
    id := c.readInt64("Enter Employee ID: ")
    newSalary := c.readFloat("Enter New Salary: ")

    err := c.service.UpdateSalary(id, newSalary)
    if errors.Is(err, repository.ErrNotFound) {
        fmt.Println("Employee not found.")
        return
    }
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Salary updated successfully.")
}

func (c *EmployeeController) getSalaryHistory() {
    fmt.Println("\n---------- SALARY HISTORY ----------")
    id := c.readInt64("Enter Employee ID: ")

    history, err := c.service.GetSalaryHistory(id)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    if len(history) == 0 {
        fmt.Println("No salary history found for this employee.")
        return
    }

    fmt.Printf("%-5s %-12s %-12s %-25s\n", "ID", "OLD SALARY", "NEW SALARY", "CHANGED AT")
    fmt.Println(strings.Repeat("-", 60))
    for _, h := range history {
        fmt.Printf(
            "%-5d %-12.2f %-12.2f %-25s\n",
            h.ID,
            h.OldSalary,
            h.NewSalary,
            h.ChangedAt.Format("2006-01-02 15:04:05"),
        )
    }
}

func (c *EmployeeController) readString(
    message string,
) string {

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

func (c *EmployeeController) readInt(
    message string,
) int {

    for {

        value := c.readString(message)

        number, err := strconv.Atoi(value)

        if err == nil {
            return number
        }

        fmt.Println("Please enter a valid integer.")
    }
}

func (c *EmployeeController) readInt64(
    message string,
) int64 {

    for {

        value := c.readString(message)

        number, err := strconv.ParseInt(value, 10, 64)

        if err == nil {
            return number
        }

        fmt.Println("Please enter a valid number.")
    }
}

func (c *EmployeeController) readFloat(
    message string,
) float64 {

    for {

        value := c.readString(message)

        number, err := strconv.ParseFloat(value, 64)

        if err == nil {
            return number
        }

        fmt.Println("Please enter a valid decimal number.")
    }
}

func (c *EmployeeController) pause() {

    fmt.Println()
    fmt.Print("Press ENTER to continue...")

    _, _ = c.reader.ReadString('\n')
}
