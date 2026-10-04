package main

import (
	"errors"
	"fmt"
	"math"
	"sync"
	"testing"
)

func validEmployee() Employee {
	return Employee{Name: "Asha", Email: "asha@corp.com", Age: 30, Dept: "Eng", Salary: 60000}
}

func TestValidate(t *testing.T) {
	tests := []struct {
		name    string
		mutate  func(*Employee)
		wantErr bool
	}{
		{"valid employee", func(e *Employee) {}, false},
		{"empty name", func(e *Employee) { e.Name = "  " }, true},
		{"email missing @", func(e *Employee) { e.Email = "ashacorp.com" }, true},
		{"email missing dot", func(e *Employee) { e.Email = "asha@corp" }, true},
		{"age too young", func(e *Employee) { e.Age = 17 }, true},
		{"age too old", func(e *Employee) { e.Age = 66 }, true},
		{"age lower boundary", func(e *Employee) { e.Age = 18 }, false},
		{"age upper boundary", func(e *Employee) { e.Age = 65 }, false},
		{"invalid dept", func(e *Employee) { e.Dept = "Space" }, true},
		{"zero salary", func(e *Employee) { e.Salary = 0 }, true},
		{"negative salary", func(e *Employee) { e.Salary = -5 }, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := validEmployee()
			tt.mutate(&e)
			err := e.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCalculateTax(t *testing.T) {
	tests := []struct {
		name   string
		annual float64
		want   float64
	}{
		{"below slab", 200000, 0},
		{"exactly 250000", 250000, 0},
		{"second slab", 300000, 2500},
		{"exactly 500000", 500000, 12500},
		{"top slab", 600000, 32500},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateTax(tt.annual); math.Abs(got-tt.want) > 0.01 {
				t.Errorf("CalculateTax(%v) = %v, want %v", tt.annual, got, tt.want)
			}
		})
	}
}

func TestCalculateSalary(t *testing.T) {
	tests := []struct {
		name        string
		base, bonus float64
		wantGross   float64
		wantTax     float64
		wantNet     float64
		wantErr     bool
	}{
		{"with bonus", 400000, 10, 440000, 9500, 430500, false},
		{"no bonus", 200000, 0, 200000, 0, 200000, false},
		{"zero base", 0, 10, 0, 0, 0, true},
		{"negative bonus", 100000, -1, 0, 0, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g, tax, n, err := CalculateSalary(tt.base, tt.bonus)
			if (err != nil) != tt.wantErr {
				t.Fatalf("err = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if math.Abs(g-tt.wantGross) > 0.01 || math.Abs(tax-tt.wantTax) > 0.01 || math.Abs(n-tt.wantNet) > 0.01 {
				t.Errorf("got (%v, %v, %v), want (%v, %v, %v)", g, tax, n, tt.wantGross, tt.wantTax, tt.wantNet)
			}
		})
	}
}

func TestAverageEmpty(t *testing.T) {
	if got := average(nil); got != 0 {
		t.Errorf("average(nil) = %v, want 0", got)
	}
}

func TestParseSalary(t *testing.T) {
	if _, err := parseSalary("abc"); err == nil {
		t.Error("expected error for abc")
	}
	if v, err := parseSalary("5000"); err != nil || v != 5000 {
		t.Errorf("got %v, %v", v, err)
	}
}

func TestGroupByDept(t *testing.T) {
	got := groupByDept([]rec{{"A", 1, "Eng"}, {"B", 2, "Eng"}, {"C", 3, "Ops"}})
	if got["Eng"] != 2 || got["Ops"] != 1 {
		t.Errorf("got %v", got)
	}
}

func TestServiceAddAndGet(t *testing.T) {
	s := NewService()
	e, err := s.Add(validEmployee())
	if err != nil {
		t.Fatalf("Add failed: %v", err)
	}
	if e.ID != 1 {
		t.Errorf("ID = %d, want 1", e.ID)
	}
	got, err := s.Get(e.ID)
	if err != nil || got.Name != "Asha" {
		t.Errorf("Get = %+v, %v", got, err)
	}
}

func TestServiceDuplicateEmail(t *testing.T) {
	s := NewService()
	if _, err := s.Add(validEmployee()); err != nil {
		t.Fatal(err)
	}
	_, err := s.Add(validEmployee())
	if !errors.Is(err, ErrDuplicateEmail) {
		t.Errorf("err = %v, want ErrDuplicateEmail", err)
	}
}

func TestServiceInvalidEmployeeRejected(t *testing.T) {
	s := NewService()
	e := validEmployee()
	e.Age = 5
	if _, err := s.Add(e); err == nil {
		t.Error("expected validation error")
	}
	if s.Count() != 0 {
		t.Errorf("Count = %d, want 0", s.Count())
	}
}

func TestServiceGetDeleteNotFound(t *testing.T) {
	s := NewService()
	if _, err := s.Get(99); !errors.Is(err, ErrNotFound) {
		t.Errorf("Get err = %v", err)
	}
	if err := s.Delete(99); !errors.Is(err, ErrNotFound) {
		t.Errorf("Delete err = %v", err)
	}
}

func TestServiceDelete(t *testing.T) {
	s := NewService()
	e, _ := s.Add(validEmployee())
	if err := s.Delete(e.ID); err != nil {
		t.Fatal(err)
	}
	if s.Count() != 0 {
		t.Error("expected empty service")
	}
}

func TestServiceConcurrent(t *testing.T) {
	s := NewService()
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			e := validEmployee()
			e.Email = fmt.Sprintf("user%d@corp.com", i)
			_, _ = s.Add(e)
			s.Count()
			_, _ = s.Get(1)
		}(i)
	}
	wg.Wait()
	if s.Count() != 100 {
		t.Errorf("Count = %d, want 100", s.Count())
	}
}

func BenchmarkValidate(b *testing.B) {
	e := validEmployee()
	for i := 0; i < b.N; i++ {
		_ = e.Validate()
	}
}

func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(600000)
	}
}

func BenchmarkServiceAdd(b *testing.B) {
	s := NewService()
	for i := 0; i < b.N; i++ {
		e := validEmployee()
		e.Email = fmt.Sprintf("u%d@corp.com", i)
		_, _ = s.Add(e)
	}
}
