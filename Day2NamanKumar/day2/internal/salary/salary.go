// Package salary calculates an employee's pay slip using basic
// arithmetic operators (+ - * / %).
//
// The percentages are simple illustrative values for learning,
// not official tax or payroll rules.
package salary

// Rates used in the calculation (percent of basic pay).
const (
	HRARate       = 0.40 // House Rent Allowance: 40% of basic
	DARate        = 0.10 // Dearness Allowance: 10% of basic
	PFRate        = 0.12 // Provident Fund: 12% of basic
	TaxRate       = 0.10 // flat illustrative tax on gross above the exemption
	TaxExemption  = 25000.0
	ProfessionTax = 200.0
	WorkingDays   = 22
	OvertimeRate  = 1.5 // overtime pays 1.5x the hourly rate
	HoursPerDay   = 8
)

// Slip holds every figure of a monthly pay slip.
type Slip struct {
	Basic, HRA, DA, Overtime, Bonus, Gross float64
	PF, Tax, ProfTax, Deductions, Net      float64
	Annual, PerDay, PerHour                float64
}

// Calculate builds a monthly pay slip from the basic salary, overtime
// hours worked and a bonus percentage (for example 5 for 5%).
func Calculate(basic float64, overtimeHours int, bonusPercent float64) Slip {
	s := Slip{Basic: basic}

	// Addition and multiplication
	s.HRA = basic * HRARate
	s.DA = basic * DARate

	// Division
	s.PerDay = basic / WorkingDays
	s.PerHour = s.PerDay / HoursPerDay

	s.Overtime = float64(overtimeHours) * s.PerHour * OvertimeRate
	s.Bonus = basic * bonusPercent / 100
	s.Gross = s.Basic + s.HRA + s.DA + s.Overtime + s.Bonus

	// Subtraction
	s.PF = basic * PFRate
	taxable := s.Gross - TaxExemption
	if taxable > 0 {
		s.Tax = taxable * TaxRate
	}
	s.ProfTax = ProfessionTax
	s.Deductions = s.PF + s.Tax + s.ProfTax
	s.Net = s.Gross - s.Deductions

	s.Annual = s.Net * 12
	return s
}

// SplitNotes shows the modulus operator: how many ₹500 notes and
// what remainder make up an amount.
func SplitNotes(amount int) (notes500, remainder int) {
	return amount / 500, amount % 500
}
