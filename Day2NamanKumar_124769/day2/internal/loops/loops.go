// Package loops solves five classic problems with for loops.
package loops

// Factorial returns n! (n * (n-1) * ... * 1). It returns 0 and false
// for negative n or when the result would overflow uint64 (n > 20).
func Factorial(n int) (uint64, bool) {
	if n < 0 || n > 20 {
		return 0, false
	}
	result := uint64(1)
	for i := 2; i <= n; i++ {
		result *= uint64(i)
	}
	return result, true
}

// Fibonacci returns the first n Fibonacci numbers: 0, 1, 1, 2, 3, 5...
func Fibonacci(n int) []uint64 {
	if n <= 0 {
		return nil
	}
	seq := make([]uint64, n)
	a, b := uint64(0), uint64(1)
	for i := 0; i < n; i++ {
		seq[i] = a
		a, b = b, a+b
	}
	return seq
}

// IsPrime reports whether n is prime. It only tests divisors up to sqrt(n).
func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// PrimesUpTo returns all primes from 2 to limit.
func PrimesUpTo(limit int) []int {
	var out []int
	for n := 2; n <= limit; n++ {
		if IsPrime(n) {
			out = append(out, n)
		}
	}
	return out
}

// Reverse returns the digits of n in reverse order: 12345 -> 54321.
// The sign is kept: -123 -> -321.
func Reverse(n int) int {
	sign := 1
	if n < 0 {
		sign, n = -1, -n
	}
	rev := 0
	for n > 0 {
		rev = rev*10 + n%10 // take the last digit
		n /= 10             // drop the last digit
	}
	return rev * sign
}

// IsPalindrome reports whether a non-negative number reads the same
// forwards and backwards, e.g. 12321.
func IsPalindrome(n int) bool {
	return n >= 0 && n == Reverse(n)
}

// IsPalindromeString reports whether text reads the same both ways,
// ignoring English letter case and spaces. It compares runes (so Indian
// scripts work) using a two-pointer for loop.
func IsPalindromeString(s string) bool {
	var r []rune
	for _, c := range s {
		if c >= 'A' && c <= 'Z' {
			c += 'a' - 'A'
		}
		if c != ' ' {
			r = append(r, c)
		}
	}
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		if r[i] != r[j] {
			return false
		}
	}
	return true
}
