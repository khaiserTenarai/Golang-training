package salary

import (
	"math"
	"testing"
)

func near(a, b float64) bool { return math.Abs(a-b) < 0.01 }

func TestCalculate(t *testing.T) {
	s := Calculate(50000, 0, 0)
	// HRA 20000, DA 5000, gross 75000, PF 6000, tax (75000-25000)*10% = 5000, PT 200
	if !near(s.Gross, 75000) || !near(s.PF, 6000) || !near(s.Tax, 5000) || !near(s.Net, 63800) {
		t.Fatalf("unexpected slip: %+v", s)
	}
}

func TestOvertimeAndBonus(t *testing.T) {
	s := Calculate(44000, 10, 5)
	// per hour = 44000/22/8 = 250; OT = 10*250*1.5 = 3750; bonus = 2200
	if !near(s.Overtime, 3750) || !near(s.Bonus, 2200) {
		t.Fatalf("OT %v bonus %v", s.Overtime, s.Bonus)
	}
}

func TestNoTaxBelowExemption(t *testing.T) {
	if s := Calculate(10000, 0, 0); s.Tax != 0 {
		t.Fatalf("expected no tax, got %v", s.Tax)
	}
}

func TestSplitNotes(t *testing.T) {
	if n, r := SplitNotes(63800); n != 127 || r != 300 {
		t.Fatalf("got %d, %d", n, r)
	}
}
