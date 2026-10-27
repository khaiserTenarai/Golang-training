package main

import "testing"

func TestCheckAge(t *testing.T) {
	// Table-driven test: several inputs checked in one test.
	tests := []struct {
		name    string
		age     int
		wantErr bool
	}{
		{name: "under age", age: 16, wantErr: true},
		{name: "exactly 18", age: 18, wantErr: false},
		{name: "adult", age: 25, wantErr: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := CheckAge(test.age)
			if (err != nil) != test.wantErr {
				t.Errorf("CheckAge(%d) error = %v; wantErr %v", test.age, err, test.wantErr)
			}
		})
	}
}
