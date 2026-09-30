package employee

import (
	"errors"
	"testing"
)

func TestEmployeeValidate(t *testing.T) {
	valid := Employee{ID: 1, Name: "Asha", Email: "asha@x.com", Age: 30, BaseSalary: 50000}

	tests := []struct {
		name   string
		modify func(e *Employee)
		want   error
	}{
		{"valid", func(e *Employee) {}, nil},
		{"zero id", func(e *Employee) { e.ID = 0 }, ErrInvalidID},
		{"blank name", func(e *Employee) { e.Name = "  " }, ErrEmptyName},
		{"bad email", func(e *Employee) { e.Email = "asha.x.com" }, ErrInvalidEmail},
		{"too young", func(e *Employee) { e.Age = 17 }, ErrInvalidAge},
		{"too old", func(e *Employee) { e.Age = 66 }, ErrInvalidAge},
		{"zero salary", func(e *Employee) { e.BaseSalary = 0 }, ErrInvalidSalary},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			e := valid
			tc.modify(&e)
			if got := e.Validate(); !errors.Is(got, tc.want) {
				t.Errorf("Validate() = %v; want %v", got, tc.want)
			}
		})
	}
}