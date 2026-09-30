package employee

import "errors"

var ErrInvalidPercent = errors.New("percent must be between 0 and 100")

// NetSalary returns base + bonus% - tax%, where tax applies to the gross amount.
func NetSalary(base, bonusPct, taxPct float64) (float64, error) {
	if base <= 0 {
		return 0, ErrInvalidSalary
	}
	if bonusPct < 0 || bonusPct > 100 || taxPct < 0 || taxPct > 100 {
		return 0, ErrInvalidPercent
	}
	gross := base + base*bonusPct/100
	return gross - gross*taxPct/100, nil
}