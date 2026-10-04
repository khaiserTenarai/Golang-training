package main

import "testing"

func TestValidateAge(t *testing.T) {
	tests := []struct {
		name string
		age  int
		want bool
	}{
		{"Valid age", 25, true},
		{"Another valid age", 40, true},
		{"Zero age", 0, false},
		{"Negative age", -5, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := ValidateAge(test.age)

			got := err == nil

			if got != test.want {
				t.Errorf(
					"Expected %v but got %v",
					test.want,
					got,
				)
			}
		})
	}
}