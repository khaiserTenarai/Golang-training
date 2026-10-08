package main

import "testing"

func TestValidateEmployee(t *testing.T) {
	tests := []struct {
		name  string
		input Employee
		want  bool
	}{
		{"valid", Employee{101, "Tom", 70000}, true},
		{"invalid ID", Employee{0, "Tom", 70000}, false},
		{"empty name", Employee{101, "", 70000}, false},
		{"invalid salary", Employee{101, "Tom", -70000}, false},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			if got := validateEmployee(tt.input); got != tt.want {
				t.Errorf("got %v; want %v", got, tt.want)
			}
		})
	}
}
