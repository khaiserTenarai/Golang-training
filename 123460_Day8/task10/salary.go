package employee

import "errors"

type Payroll struct {
	BaseSalary  float64
	Bonus       float64
	Deductions  float64
	TaxRate     float64 // e.g., 0.20 for 20%
}

// CalculateNetSalary computes the take-home pay after tax and deductions
func (p Payroll) CalculateNetSalary() (float64, error) {
	if p.BaseSalary < 0 {
		return 0, errors.New("base salary cannot be negative")
	}
	if p.Bonus < 0 {
		return 0, errors.New("bonus cannot be negative")
	}
	if p.Deductions < 0 {
		return 0, errors.New("deductions cannot be negative")
	}
	if p.TaxRate < 0 || p.TaxRate > 1 {
		return 0, errors.New("tax rate must be between 0 and 1")
	}

	gross := p.BaseSalary + p.Bonus
	taxAmount := gross * p.TaxRate
	netSalary := gross - taxAmount - p.Deductions

	return netSalary, nil
}