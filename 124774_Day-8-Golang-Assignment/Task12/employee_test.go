package table_driven

import "testing"

func TestIsValidAge(t *testing.T) {

	tests := []struct {
		name string
		age  int
		want bool
	}{
		{"valid", 25, true},
		{"zero", 0, false},
		{"negative", -5, false},
		{"maximum", 60, true},
		{"above maximum", 65, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := IsValidAge(tt.age)

			if got != tt.want {
				t.Errorf("got %v; want %v", got, tt.want)
			}
		})
	}
}
