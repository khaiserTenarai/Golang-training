package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zip_code"`
	Country string `json:"country"`
}

// Nested Department struct
type Department struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
}

// Employee struct with 8+ fields
type Employee struct {
	ID           int        `json:"id"`
	FirstName    string     `json:"first_name"`
	LastName     string     `json:"last_name"`
	Email        string     `json:"email"`
	Age          int        `json:"age"`
	Salary       float64    `json:"salary"`
	IsActive     bool       `json:"is_active"`
	Phone        string     `json:"phone"`
	Address      Address    `json:"address"`       // Nested struct
	Department   Department `json:"department"`    // Nested struct
	JoiningYear  int        `json:"joining_year"`
}

// Value receiver
// It receives a copy of Employee.
func (e Employee) GetFullName() string {
	return e.FirstName + " " + e.LastName
}

// Value receiver
func (e Employee) GetAnnualSalary() float64 {
	return e.Salary * 12
}

// Pointer receiver
// It can modify the original Employee.
func (e *Employee) GiveRaise(percent float64) {
	e.Salary += e.Salary * percent / 100
}

// Pointer receiver
func (e *Employee) Deactivate() {
	e.IsActive = false
}

// EmployeeRepository interface
type EmployeeRepository interface {
	Add(employee Employee)
	GetByID(id int) (Employee, bool)
	GetAll() []Employee
}

// In-memory repository implementation
type InMemoryEmployeeRepository struct {
	employees []Employee
}

// Implement Add
func (r *InMemoryEmployeeRepository) Add(employee Employee) {
	r.employees = append(r.employees, employee)
}

// Implement GetByID
func (r *InMemoryEmployeeRepository) GetByID(id int) (Employee, bool) {
	for _, employee := range r.employees {
		if employee.ID == id {
			return employee, true
		}
	}

	return Employee{}, false
}

// Implement GetAll
func (r *InMemoryEmployeeRepository) GetAll() []Employee {
	return r.employees
}

// Composition instead of inheritance
type EmployeeService struct {
	Repository EmployeeRepository
}

func (s EmployeeService) FindEmployee(id int) {
	employee, found := s.Repository.GetByID(id)

	if !found {
		fmt.Println("Employee not found")
		return
	}

	fmt.Println("Employee:", employee.GetFullName())
	fmt.Println("Department:", employee.Department.Name)
}

func main() {

	// Create an Employee
	employee := Employee{
		ID:          101,
		FirstName:   "Rahul",
		LastName:    "Sharma",
		Email:       "rahul@example.com",
		Age:         30,
		Salary:      75000,
		IsActive:    true,
		Phone:       "9876543210",
		JoiningYear: 2022,

		Address: Address{
			Street:  "MG Road",
			City:    "Mumbai",
			State:   "Maharashtra",
			ZipCode: "400001",
			Country: "India",
		},

		Department: Department{
			ID:       10,
			Name:     "Engineering",
			Location: "Mumbai",
		},
	}

	// -------------------------------
	// Value receiver demonstration
	// -------------------------------

	fmt.Println("Full Name:", employee.GetFullName())
	fmt.Println("Annual Salary:", employee.GetAnnualSalary())

	// -------------------------------
	// Pointer receiver demonstration
	// -------------------------------

	fmt.Println("Salary before raise:", employee.Salary)

	employee.GiveRaise(10)

	fmt.Println("Salary after raise:", employee.Salary)

	employee.Deactivate()

	fmt.Println("Is Active:", employee.IsActive)

	// -------------------------------
	// Marshal Employee -> JSON
	// -------------------------------

	jsonData, err := json.MarshalIndent(employee, "", "    ")
	if err != nil {
		fmt.Println("Marshal error:", err)
		return
	}

	fmt.Println("\nEmployee JSON:")
	fmt.Println(string(jsonData))

	// -------------------------------
	// Unmarshal JSON -> Employee
	// -------------------------------

	var employeeFromJSON Employee

	err = json.Unmarshal(jsonData, &employeeFromJSON)
	if err != nil {
		fmt.Println("Unmarshal error:", err)
		return
	}

	fmt.Println("\nEmployee after Unmarshal:")
	fmt.Printf("%+v\n", employeeFromJSON)

	// -------------------------------
	// Repository
	// -------------------------------

	repository := &InMemoryEmployeeRepository{}

	repository.Add(employee)
	repository.Add(employeeFromJSON)

	// -------------------------------
	// Composition
	// -------------------------------

	service := EmployeeService{
		Repository: repository,
	}

	service.FindEmployee(101)

	fmt.Println("\nAll Employees:")

	for _, emp := range repository.GetAll() {
		fmt.Println(emp.GetFullName())
	}
}
