package utility

import "testing"

func TestValidateName(t *testing.T) {
	if err := ValidateName("Muneera"); err != nil {
		t.Fatal(err)
	}

	if err := ValidateName(""); err == nil {
		t.Fatal("expected error")
	}
}

func TestValidateAge(t *testing.T) {
	if err := ValidateAge(25); err != nil {
		t.Fatal(err)
	}

	if err := ValidateAge(0); err == nil {
		t.Fatal("expected error")
	}
}

func TestValidateSalary(t *testing.T) {
	if err := ValidateSalary(20000); err != nil {
		t.Fatal(err)
	}

	if err := ValidateSalary(-100); err == nil {
		t.Fatal("expected error")
	}
}
