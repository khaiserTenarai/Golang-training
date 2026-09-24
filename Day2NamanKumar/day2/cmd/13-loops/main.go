// Task 13: Five problems solved with for loops.
//
// Usage: go run ./cmd/13-loops [number]   (default 12321)
package main

import (
	"fmt"
	"os"
	"strconv"

	"day2/internal/loops"
)

func main() {
	n := 12321
	if len(os.Args) > 1 {
		v, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error: please give a whole number")
			os.Exit(1)
		}
		n = v
	}

	fmt.Println("=== 1. Factorial (classic 3-part for loop) ===")
	for _, k := range []int{0, 5, 10, 20, 21} {
		if f, ok := loops.Factorial(k); ok {
			fmt.Printf("%2d! = %d\n", k, f)
		} else {
			fmt.Printf("%2d! = too large for uint64\n", k)
		}
	}

	fmt.Println("\n=== 2. Fibonacci (first 15 numbers) ===")
	fmt.Println(loops.Fibonacci(15))

	fmt.Println("\n=== 3. Prime numbers ===")
	for _, k := range []int{1, 2, 17, 21, 97, n} {
		fmt.Printf("%-6d prime? %v\n", k, loops.IsPrime(k))
	}
	fmt.Println("Primes up to 50:", loops.PrimesUpTo(50))

	fmt.Println("\n=== 4. Reverse number (while-style for loop) ===")
	for _, k := range []int{12345, 1200, -987, n} {
		fmt.Printf("%-6d -> %d\n", k, loops.Reverse(k))
	}

	fmt.Println("\n=== 5. Palindrome ===")
	for _, k := range []int{121, 12321, 123, 1001, n} {
		fmt.Printf("%-6d palindrome? %v\n", k, loops.IsPalindrome(k))
	}
	for _, s := range []string{"Madam", "nurses run", "golang", "ಕನಕ"} {
		fmt.Printf("%-11q palindrome? %v\n", s, loops.IsPalindromeString(s))
	}

	fmt.Println("\n=== All forms of the Go for loop ===")
	fmt.Print("classic       : ")
	for i := 1; i <= 5; i++ {
		fmt.Print(i, " ")
	}
	fmt.Print("\nwhile-style   : ")
	x := 1
	for x < 50 {
		fmt.Print(x, " ")
		x *= 3
	}
	fmt.Print("\ninfinite+break: ")
	count := 0
	for {
		count++
		if count%2 == 0 {
			continue // skip even numbers
		}
		if count > 9 {
			break
		}
		fmt.Print(count, " ")
	}
	fmt.Print("\nrange int     : ")
	for i := range 5 { // Go 1.22+
		fmt.Print(i, " ")
	}
	fmt.Println("\nlabeled break : first pair (i,j) with i*j == 12:")
outer:
	for i := 1; i <= 6; i++ {
		for j := 1; j <= 6; j++ {
			if i*j == 12 {
				fmt.Printf("                found i=%d j=%d\n", i, j)
				break outer
			}
		}
	}
}
