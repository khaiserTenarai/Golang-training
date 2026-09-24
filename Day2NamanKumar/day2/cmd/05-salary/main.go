// Task 5: Employee salary calculator using arithmetic operators.
//
// Usage: go run ./cmd/05-salary [basic] [overtimeHours] [bonus%]
//
//	go run ./cmd/05-salary 50000 10 5
package main

import (
	"fmt"
	"os"
	"strconv"

	"day2/internal/salary"
)

func main() {
	basic, overtime, bonus := 50000.0, 10, 5.0
	var err error
	if len(os.Args) > 1 {
		if basic, err = strconv.ParseFloat(os.Args[1], 64); err != nil || basic <= 0 {
			exit("basic salary must be a positive number")
		}
	}
	if len(os.Args) > 2 {
		if overtime, err = strconv.Atoi(os.Args[2]); err != nil || overtime < 0 {
			exit("overtime hours must be a whole number >= 0")
		}
	}
	if len(os.Args) > 3 {
		if bonus, err = strconv.ParseFloat(os.Args[3], 64); err != nil || bonus < 0 {
			exit("bonus % must be a number >= 0")
		}
	}

	s := salary.Calculate(basic, overtime, bonus)

	fmt.Println("========== MONTHLY PAY SLIP ==========")
	line("Basic salary", s.Basic, "")
	line("HRA (40% of basic)", s.HRA, "basic * 0.40")
	line("DA  (10% of basic)", s.DA, "basic * 0.10")
	line(fmt.Sprintf("Overtime (%d hrs)", overtime), s.Overtime, "hours * perHour * 1.5")
	line(fmt.Sprintf("Bonus (%.0f%%)", bonus), s.Bonus, "basic * bonus / 100")
	fmt.Println("--------------------------------------")
	line("GROSS", s.Gross, "sum of the above  (+)")
	line("PF (12% of basic)", -s.PF, "basic * 0.12")
	line("Income tax (10%)", -s.Tax, "(gross - 25,000) * 0.10")
	line("Professional tax", -s.ProfTax, "")
	fmt.Println("--------------------------------------")
	line("TOTAL DEDUCTIONS", -s.Deductions, "")
	line("NET PAY", s.Net, "gross - deductions  (-)")
	fmt.Println("======================================")
	line("Annual take-home", s.Annual, "net * 12  (*)")
	line("Per working day", s.PerDay, "basic / 22  (/)")
	line("Per hour", s.PerHour, "perDay / 8  (/)")

	notes, rem := salary.SplitNotes(int(s.Net))
	fmt.Printf("\nModulus (%%): net pay = %d x ₹500 notes + ₹%d remainder\n", notes, rem)

	fmt.Println("\n=== Assignment operators ===")
	x := 1000.0
	x += 500 // x = x + 500
	fmt.Println("x += 500  ->", x)
	x -= 200
	fmt.Println("x -= 200  ->", x)
	x *= 2
	fmt.Println("x *= 2    ->", x)
	x /= 4
	fmt.Println("x /= 4    ->", x)
	n := 17
	n %= 5
	fmt.Println("17 %= 5   ->", n)
	n++
	fmt.Println("n++       ->", n)

	fmt.Println("\nNote: the rates are simple example values, not real tax rules.")
}

func line(label string, v float64, formula string) {
	fmt.Printf("%-20s ₹%12.2f   %s\n", label, v, formula)
}

func exit(msg string) {
	fmt.Fprintln(os.Stderr, "Error:", msg)
	fmt.Fprintln(os.Stderr, "Usage: go run ./cmd/05-salary [basic] [overtimeHours] [bonus%]")
	os.Exit(1)
}
