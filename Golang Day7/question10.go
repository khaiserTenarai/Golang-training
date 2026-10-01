package main

import "errors"

type EmployeeSalary struct {
	BaseSalary float64
	Bonus      float64
	Deductions float64
}

func CalculateNetSalary(e EmployeeSalary) (float64, error) {
	if e.BaseSalary < 0 {
		return 0, errors.New("base salary cannot be negative")
	}
	if e.Bonus < 0 {
		return 0, errors.New("bonus cannot be negative")
	}
	if e.Deductions < 0 {
		return 0, errors.New("deductions cannot be negative")
	}

	netSalary := (e.BaseSalary + e.Bonus) - e.Deductions
	if netSalary < 0 {
		return 0, errors.New("deductions exceed total earnings")
	}

	return netSalary, nil
}
