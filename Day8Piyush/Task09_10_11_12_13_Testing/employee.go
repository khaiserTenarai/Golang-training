// Tasks 9-13: Employee & Salary logic used by test files.
// This file contains the code being tested — no main() needed.
// Run tests: go test -v ./...
// Run benchmarks: go test -bench=. -benchmem
// Run coverage: go test -cover
// Run coverage report: go test -coverprofile=coverage.out && go tool cover -html=coverage.out

package employee

import (
	"errors"
	"fmt"
)

// ============================================================
// Employee model and validation (Task 9)
// ============================================================

type Employee struct {
	ID         int
	Name       string
	Email      string
	Age        int
	Department string
	Salary     float64
}

func ValidateEmployee(emp Employee) error {
	if emp.Name == "" {
		return errors.New("name is required")
	}
	if len(emp.Name) < 2 {
		return errors.New("name must be at least 2 characters")
	}
	if emp.Email == "" {
		return errors.New("email is required")
	}
	if !isValidEmail(emp.Email) {
		return errors.New("invalid email format")
	}
	if emp.Age < 18 {
		return errors.New("age must be at least 18")
	}
	if emp.Age > 65 {
		return errors.New("age must be at most 65")
	}
	if emp.Department == "" {
		return errors.New("department is required")
	}
	if emp.Salary < 0 {
		return errors.New("salary cannot be negative")
	}
	return nil
}

func isValidEmail(email string) bool {
	hasAt := false
	hasDot := false
	atIndex := -1
	for i, ch := range email {
		if ch == '@' {
			if hasAt {
				return false // multiple @
			}
			hasAt = true
			atIndex = i
		}
		if ch == '.' && i > atIndex && atIndex > 0 {
			hasDot = true
		}
	}
	return hasAt && hasDot && atIndex > 0
}

// ============================================================
// Salary calculations (Task 10)
// ============================================================

func CalculateNetSalary(grossSalary float64, taxPercent float64) (float64, error) {
	if grossSalary < 0 {
		return 0, errors.New("gross salary cannot be negative")
	}
	if taxPercent < 0 || taxPercent > 100 {
		return 0, errors.New("tax percent must be between 0 and 100")
	}
	tax := grossSalary * (taxPercent / 100)
	return grossSalary - tax, nil
}

func CalculateAnnualSalary(monthlySalary float64) float64 {
	return monthlySalary * 12
}

func CalculateBonus(salary float64, performanceRating int) (float64, error) {
	if salary < 0 {
		return 0, errors.New("salary cannot be negative")
	}
	if performanceRating < 1 || performanceRating > 5 {
		return 0, errors.New("rating must be between 1 and 5")
	}
	bonusPercent := map[int]float64{
		1: 0.0,
		2: 5.0,
		3: 10.0,
		4: 15.0,
		5: 20.0,
	}
	return salary * (bonusPercent[performanceRating] / 100), nil
}

func CalculateSalaryHike(currentSalary float64, hikePercent float64) (float64, error) {
	if currentSalary < 0 {
		return 0, errors.New("salary cannot be negative")
	}
	if hikePercent < 0 {
		return 0, errors.New("hike percent cannot be negative")
	}
	hike := currentSalary * (hikePercent / 100)
	return currentSalary + hike, nil
}

func GetTaxBracket(salary float64) string {
	switch {
	case salary <= 250000:
		return "No Tax"
	case salary <= 500000:
		return "5%"
	case salary <= 1000000:
		return "20%"
	default:
		return "30%"
	}
}

// ============================================================
// Employee operations
// ============================================================

type EmployeeStore struct {
	employees map[int]Employee
	nextID    int
}

func NewEmployeeStore() *EmployeeStore {
	return &EmployeeStore{
		employees: make(map[int]Employee),
		nextID:    1,
	}
}

func (s *EmployeeStore) Add(emp Employee) (Employee, error) {
	if err := ValidateEmployee(emp); err != nil {
		return Employee{}, err
	}
	emp.ID = s.nextID
	s.nextID++
	s.employees[emp.ID] = emp
	return emp, nil
}

func (s *EmployeeStore) Get(id int) (Employee, error) {
	emp, exists := s.employees[id]
	if !exists {
		return Employee{}, fmt.Errorf("employee with ID %d not found", id)
	}
	return emp, nil
}

func (s *EmployeeStore) GetAll() []Employee {
	result := make([]Employee, 0, len(s.employees))
	for _, emp := range s.employees {
		result = append(result, emp)
	}
	return result
}

func (s *EmployeeStore) Delete(id int) error {
	if _, exists := s.employees[id]; !exists {
		return fmt.Errorf("employee with ID %d not found", id)
	}
	delete(s.employees, id)
	return nil
}

func (s *EmployeeStore) Count() int {
	return len(s.employees)
}
