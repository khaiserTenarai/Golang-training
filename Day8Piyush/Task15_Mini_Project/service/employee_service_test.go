// Task 15 Mini Project: Comprehensive tests for Employee Service.
// Covers: unit tests, table-driven tests, benchmarks, coverage, race safety.
//
// Run all tests:     go test -v ./service/...
// Race detection:    go test -race -v ./service/...
// Coverage:          go test -cover ./service/...
// Coverage report:   go test -coverprofile=coverage.out ./service/... && go tool cover -html=coverage.out
// Benchmarks:        go test -bench=. -benchmem ./service/...

package service

import (
	"sync"
	"testing"

	"employee_service/models"
)

// ============================================================
// Validation Tests (table-driven)
// ============================================================

func TestValidateEmployee(t *testing.T) {
	tests := []struct {
		name      string
		emp       models.Employee
		wantError bool
	}{
		{"valid", models.Employee{Name: "Piyush", Email: "p@test.com", Age: 25, Department: "IT", Salary: 50000}, false},
		{"empty name", models.Employee{Name: "", Email: "p@test.com", Age: 25, Department: "IT", Salary: 50000}, true},
		{"short name", models.Employee{Name: "A", Email: "p@test.com", Age: 25, Department: "IT", Salary: 50000}, true},
		{"empty email", models.Employee{Name: "Piyush", Email: "", Age: 25, Department: "IT", Salary: 50000}, true},
		{"bad email", models.Employee{Name: "Piyush", Email: "nope", Age: 25, Department: "IT", Salary: 50000}, true},
		{"underage", models.Employee{Name: "Piyush", Email: "p@t.com", Age: 16, Department: "IT", Salary: 50000}, true},
		{"overage", models.Employee{Name: "Piyush", Email: "p@t.com", Age: 70, Department: "IT", Salary: 50000}, true},
		{"age 18 ok", models.Employee{Name: "Piyush", Email: "p@t.com", Age: 18, Department: "IT", Salary: 50000}, false},
		{"age 65 ok", models.Employee{Name: "Piyush", Email: "p@t.com", Age: 65, Department: "IT", Salary: 50000}, false},
		{"neg salary", models.Employee{Name: "Piyush", Email: "p@t.com", Age: 25, Department: "IT", Salary: -100}, true},
		{"zero salary", models.Employee{Name: "Piyush", Email: "p@t.com", Age: 25, Department: "IT", Salary: 0}, false},
		{"no dept", models.Employee{Name: "Piyush", Email: "p@t.com", Age: 25, Department: "", Salary: 50000}, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidateEmployee(tc.emp)
			if tc.wantError && err == nil {
				t.Error("expected error, got nil")
			}
			if !tc.wantError && err != nil {
				t.Errorf("expected no error, got: %v", err)
			}
		})
	}
}

// ============================================================
// CRUD Tests
// ============================================================

func validEmp(email string) models.Employee {
	return models.Employee{Name: "Piyush", Email: email, Age: 25, Department: "IT", Salary: 50000}
}

func TestAdd(t *testing.T) {
	svc := NewEmployeeService()
	emp, err := svc.Add(validEmp("a@b.com"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if emp.ID != 1 {
		t.Errorf("expected ID 1, got %d", emp.ID)
	}
	if svc.Count() != 1 {
		t.Errorf("expected count 1, got %d", svc.Count())
	}
}

func TestAdd_InvalidEmployee(t *testing.T) {
	svc := NewEmployeeService()
	_, err := svc.Add(models.Employee{Name: "", Email: "a@b.com", Age: 25, Department: "IT", Salary: 50000})
	if err == nil {
		t.Error("expected error for invalid employee")
	}
}

func TestAdd_DuplicateEmail(t *testing.T) {
	svc := NewEmployeeService()
	svc.Add(validEmp("dup@test.com"))
	_, err := svc.Add(validEmp("dup@test.com"))
	if err == nil {
		t.Error("expected error for duplicate email")
	}
}

func TestGet(t *testing.T) {
	svc := NewEmployeeService()
	svc.Add(validEmp("g@t.com"))
	emp, err := svc.Get(1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if emp.Name != "Piyush" {
		t.Errorf("expected Piyush, got %s", emp.Name)
	}
}

func TestGet_NotFound(t *testing.T) {
	svc := NewEmployeeService()
	_, err := svc.Get(999)
	if err == nil {
		t.Error("expected not found error")
	}
}

func TestGetAll(t *testing.T) {
	svc := NewEmployeeService()
	svc.Add(validEmp("a1@t.com"))
	svc.Add(validEmp("a2@t.com"))
	all := svc.GetAll()
	if len(all) != 2 {
		t.Errorf("expected 2 employees, got %d", len(all))
	}
}

func TestUpdate(t *testing.T) {
	svc := NewEmployeeService()
	svc.Add(validEmp("u@t.com"))
	updated := validEmp("u@t.com")
	updated.Name = "Updated"
	emp, err := svc.Update(1, updated)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if emp.Name != "Updated" {
		t.Errorf("expected Updated, got %s", emp.Name)
	}
}

func TestUpdate_NotFound(t *testing.T) {
	svc := NewEmployeeService()
	_, err := svc.Update(999, validEmp("x@t.com"))
	if err == nil {
		t.Error("expected not found error")
	}
}

func TestDelete(t *testing.T) {
	svc := NewEmployeeService()
	svc.Add(validEmp("d@t.com"))
	err := svc.Delete(1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if svc.Count() != 0 {
		t.Errorf("expected 0 after delete, got %d", svc.Count())
	}
}

func TestDelete_NotFound(t *testing.T) {
	svc := NewEmployeeService()
	err := svc.Delete(999)
	if err == nil {
		t.Error("expected not found error")
	}
}

// ============================================================
// Salary Tests (table-driven)
// ============================================================

func TestCalculateNetSalary(t *testing.T) {
	tests := []struct {
		name    string
		gross   float64
		tax     float64
		want    float64
		wantErr bool
	}{
		{"normal", 100000, 20, 80000, false},
		{"zero tax", 50000, 0, 50000, false},
		{"full tax", 50000, 100, 0, false},
		{"neg gross", -1, 10, 0, true},
		{"neg tax", 50000, -1, 0, true},
		{"tax over 100", 50000, 101, 0, true},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := CalculateNetSalary(tc.gross, tc.tax)
			if tc.wantErr {
				if err == nil {
					t.Error("expected error")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tc.want {
				t.Errorf("expected %.2f, got %.2f", tc.want, got)
			}
		})
	}
}

func TestCalculateAnnualSalary(t *testing.T) {
	if CalculateAnnualSalary(50000) != 600000 {
		t.Error("expected 600000")
	}
}

func TestCalculateBonus(t *testing.T) {
	tests := []struct {
		name    string
		salary  float64
		rating  int
		want    float64
		wantErr bool
	}{
		{"rating 1", 100000, 1, 0, false},
		{"rating 3", 100000, 3, 10000, false},
		{"rating 5", 100000, 5, 20000, false},
		{"bad rating 0", 100000, 0, 0, true},
		{"bad rating 6", 100000, 6, 0, true},
		{"neg salary", -1, 3, 0, true},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := CalculateBonus(tc.salary, tc.rating)
			if tc.wantErr {
				if err == nil {
					t.Error("expected error")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tc.want {
				t.Errorf("expected %.2f, got %.2f", tc.want, got)
			}
		})
	}
}

func TestGetTaxBracket(t *testing.T) {
	tests := []struct {
		salary float64
		want   string
	}{
		{200000, "No Tax"},
		{400000, "5%"},
		{800000, "20%"},
		{1500000, "30%"},
	}
	for _, tc := range tests {
		got := GetTaxBracket(tc.salary)
		if got != tc.want {
			t.Errorf("salary %.0f: expected %s, got %s", tc.salary, tc.want, got)
		}
	}
}

// ============================================================
// Race Safety Test (run with: go test -race)
// ============================================================

func TestConcurrentAccess(t *testing.T) {
	svc := NewEmployeeService()
	var wg sync.WaitGroup

	// Concurrent adds
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			email := "user" + string(rune('a'+n%26)) + "@test.com"
			svc.Add(models.Employee{
				Name: "User", Email: email, Age: 25,
				Department: "IT", Salary: 50000,
			})
		}(i)
	}

	// Concurrent reads
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			svc.GetAll()
			svc.Count()
		}()
	}

	wg.Wait()
	t.Logf("Final count after concurrent ops: %d", svc.Count())
}

// ============================================================
// Benchmarks
// ============================================================

func BenchmarkValidateEmployee(b *testing.B) {
	emp := models.Employee{Name: "Piyush", Email: "p@test.com", Age: 25, Department: "IT", Salary: 50000}
	for i := 0; i < b.N; i++ {
		ValidateEmployee(emp)
	}
}

func BenchmarkAdd(b *testing.B) {
	svc := NewEmployeeService()
	for i := 0; i < b.N; i++ {
		svc.Add(models.Employee{
			Name: "User", Email: "unique" + string(rune(i%10000)) + "@t.com",
			Age: 25, Department: "IT", Salary: 50000,
		})
	}
}

func BenchmarkCalculateNetSalary(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateNetSalary(100000, 20)
	}
}

func BenchmarkCalculateBonus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateBonus(100000, 4)
	}
}
