package q10

import "testing"

func TestCalculateBonus(t *testing.T) {
	got, err := CalculateBonus(500000)
	if err != nil {
		t.Fatal(err)
	}
	if got != 50000 {
		t.Fatalf("got %v; want 50000", got)
	}
}

func TestCalculateBonusNegativeSalary(t *testing.T) {
	_, err := CalculateBonus(-10000)
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestUpdateSalary(t *testing.T) {
	emp := Employee{ID: 101, Name: "Sasi", Salary: 500000}
	err := UpdateSalary(&emp, 50000)
	if err != nil {
		t.Fatal(err)
	}
	if emp.Salary != 550000 {
		t.Fatalf("got %v; want 550000", emp.Salary)
	}
}