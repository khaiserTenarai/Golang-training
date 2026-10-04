// Tasks 9, 10, 11, 12: Unit tests, salary tests, coverage, table-driven tests
// Run tests:       go test -v
// Run coverage:    go test -cover
// Coverage report: go test -coverprofile=coverage.out && go tool cover -html=coverage.out
// Task 11 (coverage) is demonstrated by running: go test -cover -v

package employee

import (
	"testing"
)

// ============================================================
// Task 9: Unit Tests for Employee Validation
// ============================================================

func TestValidateEmployee_ValidEmployee(t *testing.T) {
	emp := Employee{
		Name:       "Piyush",
		Email:      "piyush@example.com",
		Age:        25,
		Department: "Engineering",
		Salary:     50000,
	}
	err := ValidateEmployee(emp)
	if err != nil {
		t.Errorf("expected no error for valid employee, got: %v", err)
	}
}

func TestValidateEmployee_EmptyName(t *testing.T) {
	emp := Employee{Name: "", Email: "a@b.com", Age: 25, Department: "IT", Salary: 1000}
	err := ValidateEmployee(emp)
	if err == nil {
		t.Error("expected error for empty name")
	}
}

func TestValidateEmployee_ShortName(t *testing.T) {
	emp := Employee{Name: "A", Email: "a@b.com", Age: 25, Department: "IT", Salary: 1000}
	err := ValidateEmployee(emp)
	if err == nil {
		t.Error("expected error for short name")
	}
}

func TestValidateEmployee_EmptyEmail(t *testing.T) {
	emp := Employee{Name: "Piyush", Email: "", Age: 25, Department: "IT", Salary: 1000}
	err := ValidateEmployee(emp)
	if err == nil {
		t.Error("expected error for empty email")
	}
}

func TestValidateEmployee_InvalidEmail(t *testing.T) {
	emp := Employee{Name: "Piyush", Email: "invalid-email", Age: 25, Department: "IT", Salary: 1000}
	err := ValidateEmployee(emp)
	if err == nil {
		t.Error("expected error for invalid email")
	}
}

func TestValidateEmployee_AgeTooYoung(t *testing.T) {
	emp := Employee{Name: "Piyush", Email: "p@b.com", Age: 16, Department: "IT", Salary: 1000}
	err := ValidateEmployee(emp)
	if err == nil {
		t.Error("expected error for underage")
	}
}

func TestValidateEmployee_AgeTooOld(t *testing.T) {
	emp := Employee{Name: "Piyush", Email: "p@b.com", Age: 70, Department: "IT", Salary: 1000}
	err := ValidateEmployee(emp)
	if err == nil {
		t.Error("expected error for overage")
	}
}

func TestValidateEmployee_NegativeSalary(t *testing.T) {
	emp := Employee{Name: "Piyush", Email: "p@b.com", Age: 25, Department: "IT", Salary: -5000}
	err := ValidateEmployee(emp)
	if err == nil {
		t.Error("expected error for negative salary")
	}
}

// ============================================================
// Task 10: Salary Calculation Tests
// ============================================================

func TestCalculateNetSalary(t *testing.T) {
	net, err := CalculateNetSalary(100000, 20)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if net != 80000 {
		t.Errorf("expected 80000, got %.2f", net)
	}
}

func TestCalculateNetSalary_NegativeGross(t *testing.T) {
	_, err := CalculateNetSalary(-50000, 20)
	if err == nil {
		t.Error("expected error for negative gross salary")
	}
}

func TestCalculateNetSalary_InvalidTax(t *testing.T) {
	_, err := CalculateNetSalary(50000, 110)
	if err == nil {
		t.Error("expected error for tax > 100")
	}
}

func TestCalculateAnnualSalary(t *testing.T) {
	annual := CalculateAnnualSalary(50000)
	if annual != 600000 {
		t.Errorf("expected 600000, got %.2f", annual)
	}
}

func TestCalculateBonus(t *testing.T) {
	bonus, err := CalculateBonus(100000, 5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if bonus != 20000 {
		t.Errorf("expected 20000, got %.2f", bonus)
	}
}

func TestCalculateBonus_InvalidRating(t *testing.T) {
	_, err := CalculateBonus(100000, 6)
	if err == nil {
		t.Error("expected error for invalid rating")
	}
}

func TestCalculateSalaryHike(t *testing.T) {
	newSalary, err := CalculateSalaryHike(50000, 10)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if newSalary != 55000 {
		t.Errorf("expected 55000, got %.2f", newSalary)
	}
}

func TestGetTaxBracket(t *testing.T) {
	if GetTaxBracket(200000) != "No Tax" {
		t.Error("expected 'No Tax' for 200000")
	}
	if GetTaxBracket(400000) != "5%" {
		t.Error("expected '5%' for 400000")
	}
	if GetTaxBracket(800000) != "20%" {
		t.Error("expected '20%' for 800000")
	}
	if GetTaxBracket(1500000) != "30%" {
		t.Error("expected '30%' for 1500000")
	}
}

// ============================================================
// Task 12: Table-Driven Tests
// ============================================================

func TestValidateEmployee_TableDriven(t *testing.T) {
	tests := []struct {
		name      string
		employee  Employee
		wantError bool
	}{
		{
			name:      "valid employee",
			employee:  Employee{Name: "Piyush", Email: "p@test.com", Age: 25, Department: "IT", Salary: 50000},
			wantError: false,
		},
		{
			name:      "empty name",
			employee:  Employee{Name: "", Email: "p@test.com", Age: 25, Department: "IT", Salary: 50000},
			wantError: true,
		},
		{
			name:      "name too short",
			employee:  Employee{Name: "A", Email: "p@test.com", Age: 25, Department: "IT", Salary: 50000},
			wantError: true,
		},
		{
			name:      "empty email",
			employee:  Employee{Name: "Piyush", Email: "", Age: 25, Department: "IT", Salary: 50000},
			wantError: true,
		},
		{
			name:      "invalid email no @",
			employee:  Employee{Name: "Piyush", Email: "notanemail", Age: 25, Department: "IT", Salary: 50000},
			wantError: true,
		},
		{
			name:      "age below 18",
			employee:  Employee{Name: "Piyush", Email: "p@test.com", Age: 15, Department: "IT", Salary: 50000},
			wantError: true,
		},
		{
			name:      "age above 65",
			employee:  Employee{Name: "Piyush", Email: "p@test.com", Age: 70, Department: "IT", Salary: 50000},
			wantError: true,
		},
		{
			name:      "boundary age 18",
			employee:  Employee{Name: "Piyush", Email: "p@test.com", Age: 18, Department: "IT", Salary: 50000},
			wantError: false,
		},
		{
			name:      "boundary age 65",
			employee:  Employee{Name: "Piyush", Email: "p@test.com", Age: 65, Department: "IT", Salary: 50000},
			wantError: false,
		},
		{
			name:      "negative salary",
			employee:  Employee{Name: "Piyush", Email: "p@test.com", Age: 25, Department: "IT", Salary: -100},
			wantError: true,
		},
		{
			name:      "zero salary is valid",
			employee:  Employee{Name: "Piyush", Email: "p@test.com", Age: 25, Department: "IT", Salary: 0},
			wantError: false,
		},
		{
			name:      "empty department",
			employee:  Employee{Name: "Piyush", Email: "p@test.com", Age: 25, Department: "", Salary: 50000},
			wantError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidateEmployee(tc.employee)
			if tc.wantError && err == nil {
				t.Errorf("expected error but got nil")
			}
			if !tc.wantError && err != nil {
				t.Errorf("expected no error but got: %v", err)
			}
		})
	}
}

func TestCalculateNetSalary_TableDriven(t *testing.T) {
	tests := []struct {
		name      string
		gross     float64
		tax       float64
		expected  float64
		wantError bool
	}{
		{"normal tax", 100000, 10, 90000, false},
		{"zero tax", 100000, 0, 100000, false},
		{"full tax", 100000, 100, 0, false},
		{"negative salary", -50000, 10, 0, true},
		{"negative tax", 50000, -5, 0, true},
		{"tax over 100", 50000, 150, 0, true},
		{"zero salary zero tax", 0, 0, 0, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := CalculateNetSalary(tc.gross, tc.tax)
			if tc.wantError {
				if err == nil {
					t.Errorf("expected error but got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result != tc.expected {
				t.Errorf("expected %.2f, got %.2f", tc.expected, result)
			}
		})
	}
}

func TestCalculateBonus_TableDriven(t *testing.T) {
	tests := []struct {
		name      string
		salary    float64
		rating    int
		expected  float64
		wantError bool
	}{
		{"rating 1 - no bonus", 100000, 1, 0, false},
		{"rating 2 - 5%", 100000, 2, 5000, false},
		{"rating 3 - 10%", 100000, 3, 10000, false},
		{"rating 4 - 15%", 100000, 4, 15000, false},
		{"rating 5 - 20%", 100000, 5, 20000, false},
		{"rating 0 - invalid", 100000, 0, 0, true},
		{"rating 6 - invalid", 100000, 6, 0, true},
		{"negative salary", -5000, 3, 0, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := CalculateBonus(tc.salary, tc.rating)
			if tc.wantError {
				if err == nil {
					t.Errorf("expected error but got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result != tc.expected {
				t.Errorf("expected %.2f, got %.2f", tc.expected, result)
			}
		})
	}
}

// ============================================================
// Employee Store Tests
// ============================================================

func TestEmployeeStore_Add(t *testing.T) {
	store := NewEmployeeStore()
	emp := Employee{Name: "Piyush", Email: "p@test.com", Age: 25, Department: "IT", Salary: 50000}
	created, err := store.Add(emp)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if created.ID != 1 {
		t.Errorf("expected ID 1, got %d", created.ID)
	}
	if store.Count() != 1 {
		t.Errorf("expected count 1, got %d", store.Count())
	}
}

func TestEmployeeStore_Get(t *testing.T) {
	store := NewEmployeeStore()
	emp := Employee{Name: "Piyush", Email: "p@test.com", Age: 25, Department: "IT", Salary: 50000}
	store.Add(emp)

	found, err := store.Get(1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if found.Name != "Piyush" {
		t.Errorf("expected Piyush, got %s", found.Name)
	}
}

func TestEmployeeStore_GetNotFound(t *testing.T) {
	store := NewEmployeeStore()
	_, err := store.Get(999)
	if err == nil {
		t.Error("expected error for non-existent employee")
	}
}

func TestEmployeeStore_Delete(t *testing.T) {
	store := NewEmployeeStore()
	emp := Employee{Name: "Piyush", Email: "p@test.com", Age: 25, Department: "IT", Salary: 50000}
	store.Add(emp)

	err := store.Delete(1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if store.Count() != 0 {
		t.Errorf("expected count 0 after delete, got %d", store.Count())
	}
}
