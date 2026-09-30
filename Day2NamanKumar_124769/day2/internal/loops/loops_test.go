package loops

import (
	"reflect"
	"testing"
)

func TestFactorial(t *testing.T) {
	cases := map[int]uint64{0: 1, 1: 1, 5: 120, 10: 3628800, 20: 2432902008176640000}
	for n, want := range cases {
		if got, ok := Factorial(n); !ok || got != want {
			t.Errorf("Factorial(%d) = %d, want %d", n, got, want)
		}
	}
	if _, ok := Factorial(-1); ok {
		t.Error("negative should fail")
	}
	if _, ok := Factorial(21); ok {
		t.Error("21 should overflow")
	}
}

func TestFibonacci(t *testing.T) {
	want := []uint64{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}
	if got := Fibonacci(10); !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v", got)
	}
	if Fibonacci(0) != nil {
		t.Fatal("want nil")
	}
}

func TestIsPrime(t *testing.T) {
	for _, p := range []int{2, 3, 5, 7, 13, 97, 7919} {
		if !IsPrime(p) {
			t.Errorf("%d should be prime", p)
		}
	}
	for _, n := range []int{-7, 0, 1, 4, 9, 15, 100} {
		if IsPrime(n) {
			t.Errorf("%d should not be prime", n)
		}
	}
	if got := PrimesUpTo(20); !reflect.DeepEqual(got, []int{2, 3, 5, 7, 11, 13, 17, 19}) {
		t.Errorf("PrimesUpTo(20) = %v", got)
	}
}

func TestReverse(t *testing.T) {
	cases := map[int]int{12345: 54321, 100: 1, 7: 7, 0: 0, -123: -321}
	for in, want := range cases {
		if got := Reverse(in); got != want {
			t.Errorf("Reverse(%d) = %d, want %d", in, got, want)
		}
	}
}

func TestPalindrome(t *testing.T) {
	for _, n := range []int{0, 7, 121, 12321, 1001} {
		if !IsPalindrome(n) {
			t.Errorf("%d should be palindrome", n)
		}
	}
	for _, n := range []int{10, 123, -121} {
		if IsPalindrome(n) {
			t.Errorf("%d should not be palindrome", n)
		}
	}
	if !IsPalindromeString("Madam") || !IsPalindromeString("nurses run") || IsPalindromeString("golang") {
		t.Error("string palindrome wrong")
	}
	if !IsPalindromeString("ಕನಕ") {
		t.Error("unicode palindrome wrong")
	}
}
