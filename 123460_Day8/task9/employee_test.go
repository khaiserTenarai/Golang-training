package employee

import (
	"testing"
)

func TestValidateEmployee(t *testing.T) {
	tests := []struct {
		name    string
		emp     Employee
		wantErr bool
	}{
		{
			name: "Valid Employee",
			emp: Employee{
				ID:     1,
				Name:   "John Doe",
				Age:    30,
				Email:  "john.doe@example.com",
				Salary: 75000.00,
			},
			wantErr: false,
		},
		{
			name: "Invalid ID (Zero or Negative)",
			emp: Employee{
				ID:     0,
				Name:   "John Doe",
				Age:    30,
				Email:  "john.doe@example.com",
				Salary: 75000.00,
			},
			wantErr: true,
		},
		{
			name: "Empty Name",
			emp: Employee{
				ID:     1,
				Name:   "   ",
				Age:    30,
				Email:  "john.doe@example.com",
				Salary: 75000.00,
			},
			wantErr: true,
		},
		{
			name: "Underage (Less than 18)",
			emp: Employee{
				ID:     1,
				Name:   "Jane Doe",
				Age:    17,
				Email:  "jane.doe@example.com",
				Salary: 50000.00,
			},
			wantErr: true,
		},
		{
			name: "Overage (Greater than 65)",
			emp: Employee{
				ID:     1,
				Name:   "Senior Citizen",
				Age:    66,
				Email:  "senior@example.com",
				Salary: 90000.00,
			},
			wantErr: true,
		},
		{
			name: "Invalid Email Format",
			emp: Employee{
				ID:     1,
				Name:   "John Doe",
				Age:    25,
				Email:  "john.doeexample.com",
				Salary: 60000.00,
			},
			wantErr: true,
		},
		{
			name: "Non-positive Salary",
			emp: Employee{
				ID:     1,
				Name:   "John Doe",
				Age:    28,
				Email:  "john.doe@example.com",
				Salary: 0,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.emp.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}