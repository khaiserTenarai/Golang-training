// Task 4: Type conversion string -> integer -> float (and back).
// Go never converts types automatically; every conversion is explicit.
//
// Usage: go run ./cmd/04-conversion [values...]
package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputs := []string{"42", "  75000 ", "3.99", "abc", "99999999999999999999"}
	if len(os.Args) > 1 {
		inputs = os.Args[1:]
	}

	for _, in := range inputs {
		fmt.Printf("\nInput %q\n", in)
		convert(in)
	}

	fmt.Println("\n=== Other conversions ===")
	f := 9.87
	fmt.Println("float64 9.87 -> int   :", int(f), "(truncates, doesn't round)")
	fmt.Println("float64(7) / 2        :", float64(7)/2, "   but int 7 / 2 =", 7/2)
	fmt.Println("int -> string (Itoa)  :", strconv.Itoa(2026))
	fmt.Println("float -> string       :", strconv.FormatFloat(1234.5678, 'f', 2, 64))
	fmt.Println("string(rune(65))      :", string(rune(65)), "(65 becomes the character 'A', not \"65\")")
	bv, _ := strconv.ParseBool("true")
	fmt.Println("string -> bool        :", bv)
	var big int64 = 300
	fmt.Println("int64 300 -> int8     :", int8(big), "(overflow: the value doesn't fit)")
}

// convert walks one input through string -> int -> float64.
func convert(raw string) {
	s := strings.TrimSpace(raw) // user input often has stray spaces

	// Step 1: string -> int
	n, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("  string -> int   : FAILED:", err)
		// Maybe it is a decimal: try string -> float directly.
		f, ferr := strconv.ParseFloat(s, 64)
		if ferr != nil {
			fmt.Println("  string -> float : FAILED:", ferr)
			return
		}
		fmt.Printf("  string -> float : %.2f\n", f)
		if f > math.MaxInt64 || f < math.MinInt64 {
			fmt.Println("  float  -> int   : REFUSED: too large for int (would give a garbage value)")
			return
		}
		fmt.Printf("  float  -> int   : %d (decimal part dropped)\n", int(f))
		return
	}
	fmt.Printf("  string -> int   : %d  (%T)\n", n, n)

	// Step 2: int -> float64
	f := float64(n)
	fmt.Printf("  int    -> float : %.2f  (%T)\n", f, f)
	fmt.Printf("  with a 10%% hike : %.2f\n", f*1.10)
}
