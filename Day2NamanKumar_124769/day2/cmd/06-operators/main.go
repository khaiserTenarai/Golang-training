// Task 6: Comparison and logical operators, used for real HR decisions.
package main

import "fmt"

type Employee struct {
	Name       string
	Experience int     // years
	Rating     float64 // out of 5
	Salary     float64
	IsManager  bool
	OnNotice   bool
}

func main() {
	fmt.Println("=== Comparison operators ===")
	a, b := 10, 20
	fmt.Printf("%d == %d : %v\n", a, b, a == b)
	fmt.Printf("%d != %d : %v\n", a, b, a != b)
	fmt.Printf("%d <  %d : %v\n", a, b, a < b)
	fmt.Printf("%d <= %d : %v\n", a, b, a <= b)
	fmt.Printf("%d >  %d : %v\n", a, b, a > b)
	fmt.Printf("%d >= %d : %v\n", a, b, a >= b)
	fmt.Println(`"apple" < "banana" :`, "apple" < "banana", "(strings compare alphabetically)")

	fmt.Println("\n=== Logical operators ===")
	t, f := true, false
	fmt.Println("true && false :", t && f, " (AND: both must be true)")
	fmt.Println("true || false :", t || f, "  (OR: at least one true)")
	fmt.Println("!true         :", !t, " (NOT: flips the value)")

	fmt.Println("\n=== Truth table ===")
	fmt.Println("A      B      A&&B   A||B")
	for _, x := range []bool{true, false} {
		for _, y := range []bool{true, false} {
			fmt.Printf("%-6v %-6v %-6v %-6v\n", x, y, x && y, x || y)
		}
	}

	employees := []Employee{
		{"Anita Rao", 5, 4.5, 90000, false, false},
		{"Rahul Mehta", 2, 4.8, 45000, false, false},
		{"Priya Nair", 8, 3.5, 150000, true, false},
		{"Arjun Singh", 4, 4.2, 70000, false, true},
	}

	fmt.Println("\n=== HR rules ===")
	fmt.Println("Bonus   : (experience >= 3 && rating >= 4.0 || isManager) && !onNotice")
	fmt.Println("Promote : experience >= 4 && rating >= 4.5")
	fmt.Println("Review  : rating < 3.8 || salary > 120000")
	fmt.Println()
	fmt.Printf("%-12s %-4s %-6s %-7s %-7s %-6s\n", "Name", "Exp", "Rating", "Bonus", "Promote", "Review")
	for _, e := range employees {
		bonus := (e.Experience >= 3 && e.Rating >= 4.0 || e.IsManager) && !e.OnNotice
		promote := e.Experience >= 4 && e.Rating >= 4.5
		review := e.Rating < 3.8 || e.Salary > 120000
		fmt.Printf("%-12s %-4d %-6.1f %-7v %-7v %-6v\n", e.Name, e.Experience, e.Rating, bonus, promote, review)
	}

	fmt.Println("\n=== Short-circuit evaluation ===")
	fmt.Println("false && check() ->", false && check("AND"), "(check never ran)")
	fmt.Println("true  || check() ->", true || check("OR"), "(check never ran)")
	fmt.Println("true  && check() ->", true && check("AND"))

	// Real use: guard against a nil pointer before using it
	var e *Employee
	if e != nil && e.Salary > 50000 {
		fmt.Println("high earner")
	} else {
		fmt.Println("Safe: e is nil, so e.Salary was never read (no panic)")
	}
}

func check(op string) bool {
	fmt.Print("[check() ran for ", op, "] ")
	return true
}
