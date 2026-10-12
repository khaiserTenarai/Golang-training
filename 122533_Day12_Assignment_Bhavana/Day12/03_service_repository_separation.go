// Day 12 - Task 3: Separate Service and Repository
//
// The repository only knows how to store and find employees. The
// service decides what's allowed (name required, salary above 0, no
// repeated email) and asks the repository to do the storing. The
// service only sees the EmployeeRepository interface, so it doesn't
// care how the data is actually kept.
//
// This one runs in the terminal instead of over HTTP, to keep the
// focus on the two layers.
package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Employee struct {
	Name   string
	Email  string
	Salary float64
}

// ---------- repository ----------

type EmployeeRepository interface {
	Save(e Employee)
	FindAll() []Employee
	FindByEmail(email string) (Employee, bool)
}

type memoryRepository struct {
	employees []Employee
}

func (m *memoryRepository) Save(e Employee) {
	m.employees = append(m.employees, e)
}

func (m *memoryRepository) FindAll() []Employee {
	return m.employees
}

func (m *memoryRepository) FindByEmail(email string) (Employee, bool) {
	for _, e := range m.employees {
		if strings.EqualFold(e.Email, email) {
			return e, true
		}
	}
	return Employee{}, false
}

// ---------- service ----------

type EmployeeService struct {
	repo EmployeeRepository
}

func (s *EmployeeService) Register(name, email string, salary float64) error {
	if strings.TrimSpace(name) == "" {
		return errors.New("name can't be empty")
	}
	if salary <= 0 {
		return errors.New("salary must be greater than 0")
	}
	if _, exists := s.repo.FindByEmail(email); exists {
		return errors.New("that email is already registered")
	}

	s.repo.Save(Employee{Name: name, Email: email, Salary: salary})
	return nil
}

func (s *EmployeeService) List() []Employee {
	return s.repo.FindAll()
}

// ---------- main ----------

func ask(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	service := &EmployeeService{repo: &memoryRepository{}}

	for {
		fmt.Println("\n1. Add employee")
		fmt.Println("2. List employees")
		fmt.Println("3. Exit")

		switch ask(reader, "Choose an option: ") {
		case "1":
			name := ask(reader, "Name: ")
			email := ask(reader, "Email: ")
			salary, _ := strconv.ParseFloat(ask(reader, "Salary: "), 64)

			if err := service.Register(name, email, salary); err != nil {
				fmt.Println("Couldn't add the employee:", err)
			} else {
				fmt.Println("Employee added.")
			}
		case "2":
			employees := service.List()
			if len(employees) == 0 {
				fmt.Println("Nobody has been added yet.")
			}
			for _, e := range employees {
				fmt.Printf("%s <%s> - %.2f\n", e.Name, e.Email, e.Salary)
			}
		case "3":
			fmt.Println("Bye!")
			return
		default:
			fmt.Println("Please pick 1, 2 or 3.")
		}
	}
}
