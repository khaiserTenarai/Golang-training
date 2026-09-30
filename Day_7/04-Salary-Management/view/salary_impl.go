package view
import (
	"fmt"

	"salary-management/model"
)

type SalaryViewImpl struct {
}

func NewSalaryView() SalaryView {
	return &SalaryViewImpl{}
}

func (v *SalaryViewImpl) ReadSalaryUpdate() (
	int,
	float64,
) {

	var employeeID int
	var salary float64

	fmt.Println()
	fmt.Println("---------- Update Salary ----------")

	fmt.Print("Enter Employee ID: ")
	fmt.Scan(&employeeID)

	fmt.Print("Enter New Salary: ")
	fmt.Scan(&salary)

	return employeeID, salary
}

func (v *SalaryViewImpl) ReadEmployeeID() int {

	var employeeID int

	fmt.Print("Enter Employee ID: ")
	fmt.Scan(&employeeID)

	return employeeID
}

func (v *SalaryViewImpl) DisplayHistory(
	history []model.SalaryHistory,
) {

	if len(history) == 0 {
		fmt.Println("No salary history found.")
		return
	}

	fmt.Println()
	fmt.Println("---------- Salary History ----------")

	for _, item := range history {

		fmt.Println(
			"History ID:",
			item.ID,
			"| Employee ID:",
			item.EmployeeID,
			"| Old Salary:",
			item.OldSalary,
			"| New Salary:",
			item.NewSalary,
			"| Changed At:",
			item.ChangedAt.Format(
				"2006-01-02 15:04:05",
			),
		)
	}
}
