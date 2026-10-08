package main

import (
	"sync"
	"testing"
)

func TestValidateEmployee(t *testing.T) {
	tests := []struct {
		name     string
		employee Employee
		want     bool
	}{
		{
			name:     "valid employee",
			employee: Employee{ID: 1, Name: "Abirami", Salary: 30000},
			want:     true,
		},
		{
			name:     "invalid ID",
			employee: Employee{ID: 0, Name: "Abirami", Salary: 30000},
			want:     false,
		},
		{
			name:     "empty name",
			employee: Employee{ID: 1, Name: "", Salary: 30000},
			want:     false,
		},
		{
			name:     "negative salary",
			employee: Employee{ID: 1, Name: "Abirami", Salary: -1000},
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ValidateEmployee(tt.employee)

			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmployeeRace(t *testing.T) {
	employee := Employee{
		ID:     1,
		Name:   "Abirami",
		Salary: 30000,
	}

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()

		for i := 0; i < 1000; i++ {
			employee.Salary = float64(i)
		}
	}()

	go func() {
		defer wg.Done()

		for i := 0; i < 1000; i++ {
			ValidateEmployee(employee)
		}
	}()

	wg.Wait()
}
