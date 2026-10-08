package q11

import "testing"

func TestDiscountedPrice(t *testing.T) {
	tests := []struct {
		name   string
		price  float64
		member bool
		want   float64
	}{
		{"normal", 100, false, 100},
		{"member", 100, true, 90},
		{"negative", -10, false, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DiscountedPrice(tt.price, tt.member); got != tt.want {
				t.Errorf("got %v; want %v", got, tt.want)
			}
		})
	}
}	
