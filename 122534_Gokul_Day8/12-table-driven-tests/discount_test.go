package main

import "testing"

// "Table-driven" just means: instead of writing a separate test function
// for every case, we make a table (a slice of structs) listing each
// input alongside its expected output, then loop over the table and
// check each row. It's a very common pattern in Go tests because it
// keeps things short even when there are many cases to check.

func TestDiscountPercent(t *testing.T) {
	cases := []struct {
		name   string
		amount float64
		want   float64
	}{
		{"below 1000 gets no discount", 500, 0},
		{"exactly 1000 gets 5 percent", 1000, 5},
		{"between 1000 and 5000 gets 5 percent", 3000, 5},
		{"exactly 5000 gets 10 percent", 5000, 10},
		{"between 5000 and 10000 gets 10 percent", 7000, 10},
		{"10000 or more gets 20 percent", 10000, 20},
		{"well above 10000 still gets 20 percent", 50000, 20},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := DiscountPercent(c.amount)
			if got != c.want {
				t.Errorf("DiscountPercent(%v) = %v, want %v", c.amount, got, c.want)
			}
		})
	}
}
