package utility

import (
	"testing"

	"ems/model"
)

func TestValidateEmployee(t *testing.T) {

	employee := model.Employee{
		Name: "Ganesh",
		Age:  25,
	}

	err := ValidateEmployee(employee)

	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
}
