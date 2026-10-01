package assignment_task12

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"empty", "", true},
		{"single", "a", true},
		{"palindrome", "level", true},
		{"not", "golang", false},
		{"numbers", "1221", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome(tt.input); got != tt.want {
				t.Errorf("got %v; want %v", got, tt.want)
			}
		})
	}
}