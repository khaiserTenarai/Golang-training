package main

import "testing"

func TestCalculate(t *testing.T) {
	tests := []struct {
		name    string
		a, b    float64
		op      string
		want    float64
		wantErr bool
	}{
		{"add", 2, 3, "+", 5, false},
		{"sub", 5, 3, "-", 2, false},
		{"mul", 4, 3, "*", 12, false},
		{"mul-x", 4, 3, "x", 12, false},
		{"div", 10, 4, "/", 2.5, false},
		{"div-zero", 1, 0, "/", 0, true},
		{"bad-op", 1, 1, "%", 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Calculate(tt.a, tt.op, tt.b)
			if (err != nil) != tt.wantErr {
				t.Fatalf("err = %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
