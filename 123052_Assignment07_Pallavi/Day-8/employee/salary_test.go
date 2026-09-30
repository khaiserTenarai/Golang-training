package employee

import (
	"errors"
	"math"
	"testing"
)

func TestNetSalary(t *testing.T) {
	tests := []struct {
		name             string
		base, bonus, tax float64
		want             float64
		wantErr          error
	}{
		{"no bonus no tax", 1000, 0, 0, 1000, nil},
		{"bonus only", 1000, 10, 0, 1100, nil},
		{"bonus and tax", 1000, 10, 20, 880, nil},
		{"zero base", 0, 10, 10, 0, ErrInvalidSalary},
		{"bonus too high", 1000, 101, 0, 0, ErrInvalidPercent},
		{"negative tax", 1000, 0, -1, 0, ErrInvalidPercent},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := NetSalary(tc.base, tc.bonus, tc.tax)
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("error = %v; want %v", err, tc.wantErr)
			}
			if math.Abs(got-tc.want) > 0.001 {
				t.Errorf("NetSalary = %.2f; want %.2f", got, tc.want)
			}
		})
	}
}