package mathutil

import "testing"

func TestDivide(t *testing.T) {
	// 1. Define the "table" of test cases
	tests := []struct {
		name      string  // Name of the test case
		a         float64 // Input 1
		b         float64 // Input 2
		want      float64 // Expected result
		expectErr bool    // Do we expect an error?
	}{
		// 2. Populate the test cases
		{
			name:      "Simple division",
			a:         10.0,
			b:         2.0,
			want:      5.0,
			expectErr: false,
		},
		{
			name:      "Division with negative result",
			a:         -15.0,
			b:         3.0,
			want:      -5.0,
			expectErr: false,
		},
		{
			name:      "Divide by zero",
			a:         10.0,
			b:         0.0,
			want:      0.0,
			expectErr: true,
		},
		{
			name:      "Fractional result",
			a:         5.0,
			b:         2.0,
			want:      2.5,
			expectErr: false,
		},
	}

	// 3. Loop through the table
	for _, tt := range tests {
		// t.Run creates a sub-test using the 'name' field
		t.Run(tt.name, func(t *testing.T) {
			got, err := Divide(tt.a, tt.b)

			// Check if the error status matches our expectation
			if (err != nil) != tt.expectErr {
				t.Errorf("Divide() error = %v, expectErr %v", err, tt.expectErr)
				return
			}

			// Check if the result matches our expectation
			if !tt.expectErr && got != tt.want {
				t.Errorf("Divide() = %v, want %v", got, tt.want)
			}
		})
	}
}