package main

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
			name:    "valid employee",
			emp:     Employee{Name: "Alice", Age: 30, Role: "Developer"},
			wantErr: false,
		},
		{
			name:    "empty name",
			emp:     Employee{Name: "", Age: 30, Role: "Developer"},
			wantErr: true,
		},
		{
			name:    "underage employee",
			emp:     Employee{Name: "Bob", Age: 17, Role: "Intern"},
			wantErr: true,
		},
		{
			name:    "overage employee",
			emp:     Employee{Name: "Charlie", Age: 66, Role: "Manager"},
			wantErr: true,
		},
		{
			name:    "empty role",
			emp:     Employee{Name: "David", Age: 25, Role: ""},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateEmployee(tt.emp)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateEmployee() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}