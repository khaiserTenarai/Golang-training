package utility
import (
	"errors"
	"strings"
	"time"

	"attendance_leave/model"
)

var (
	InvalidIDError        = errors.New("invalid ID")
	InvalidNameError      = errors.New("invalid employee name")
	InvalidEmailError     = errors.New("invalid email")
	InvalidReasonError    = errors.New("invalid leave reason")
	InvalidDateError      = errors.New("invalid leave dates")
)

func ValidateID(id int) error {

	if id <= 0 {
		return InvalidIDError
	}

	return nil
}

func ValidateEmployee(
	employee model.Employee,
) error {

	if strings.TrimSpace(employee.Name) == "" {
		return InvalidNameError
	}

	return nil
}

func ValidateEmail(email string) error {

	email = strings.TrimSpace(email)

	if len(email) <= 3 || len(email) >= 100 {
		return InvalidEmailError
	}

	if !strings.Contains(email, "@") ||
		!strings.Contains(email, ".") {

		return InvalidEmailError
	}

	return nil
}

func ValidateLeave(
	leave model.Leave,
) error {

	if leave.EmployeeID <= 0 {
		return InvalidIDError
	}

	if leave.FromDate.IsZero() ||
		leave.ToDate.IsZero() {

		return InvalidDateError
	}

	if leave.ToDate.Before(leave.FromDate) {
		return InvalidDateError
	}

	if strings.TrimSpace(leave.Reason) == "" {
		return InvalidReasonError
	}

	return nil
}

func ParseDate(value string) (time.Time, error) {

	return time.Parse(
		"2006-01-02",
		value,
	)
}
