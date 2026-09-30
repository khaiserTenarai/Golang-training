package main

import "testing"

func TestNormalizePagination(t *testing.T) {
	cases := []struct {
		inputPage, inputSize   int
		expectedPage, expected int
	}{
		{0, 0, 1, 10},
		{-5, -5, 1, 10},
		{2, 20, 2, 20},
		{1, 500, 1, 100},
	}

	for _, c := range cases {
		page, size := normalizePagination(c.inputPage, c.inputSize)
		if page != c.expectedPage || size != c.expected {
			t.Errorf("normalizePagination(%d, %d) = (%d, %d), want (%d, %d)",
				c.inputPage, c.inputSize, page, size, c.expectedPage, c.expected)
		}
	}
}
