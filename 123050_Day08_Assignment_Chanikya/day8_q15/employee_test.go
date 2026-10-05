package main

import "testing"

func TestValidateEmployee(t *testing.T) {
	tests := []struct {
		name     string
		employee Employee
		wantErr  bool
	}{
		{
			name:     "Valid employee",
			employee: Employee{Name: "Chanikya", Salary: 50000},
			wantErr:  false,
		},
		{
			name:     "Missing name",
			employee: Employee{Name: "", Salary: 50000},
			wantErr:  true,
		},
		{
			name:     "Invalid salary",
			employee: Employee{Name: "Chanikya", Salary: 0},
			wantErr:  true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := ValidateEmployee(test.employee)

			if (err != nil) != test.wantErr {
				t.Errorf("expected error: %v, got: %v", test.wantErr, err)
			}
		})
	}
}
