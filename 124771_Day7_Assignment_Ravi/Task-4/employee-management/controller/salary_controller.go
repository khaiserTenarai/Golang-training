package controller

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"employee-management/service"
	"employee-management/view"
)

type SalaryController struct {
	service service.SalaryService
	reader  *bufio.Reader
}

func NewSalaryController(
	service service.SalaryService,
) *SalaryController {

	return &SalaryController{
		service: service,
		reader:  bufio.NewReader(os.Stdin),
	}
}

// ----------------------------------------------------
// Start
// ----------------------------------------------------

func (c *SalaryController) Start() {

	for {

		fmt.Println()
		fmt.Println("======================================")
		fmt.Println("        SALARY MANAGEMENT SYSTEM")
		fmt.Println("======================================")
		fmt.Println("1. Update Employee Salary")
		fmt.Println("2. View Salary History")
		fmt.Println("3. Exit")
		fmt.Println("======================================")

		choice :=
			c.readString(
				"Enter choice: ",
			)

		switch choice {

		case "1":

			c.updateSalary()

		case "2":

			c.viewSalaryHistory()

		case "3":

			fmt.Println()
			fmt.Println(
				"Exiting Salary Management System...",
			)

			return

		default:

			fmt.Println()
			fmt.Println(
				"Invalid choice.",
			)
		}
	}
}

// ----------------------------------------------------
// Update Salary
// ----------------------------------------------------

func (c *SalaryController) updateSalary() {

	fmt.Println()
	fmt.Println("======================================")
	fmt.Println("        UPDATE EMPLOYEE SALARY")
	fmt.Println("======================================")

	employeeIDInput :=
		c.readString(
			"Enter employee ID: ",
		)

	employeeID, err :=
		strconv.ParseInt(
			employeeIDInput,
			10,
			64,
		)

	if err != nil {

		fmt.Println(
			"Invalid employee ID.",
		)

		return
	}

	newSalaryInput :=
		c.readString(
			"Enter new salary: ",
		)

	newSalary, err :=
		strconv.ParseFloat(
			newSalaryInput,
			64,
		)

	if err != nil {

		fmt.Println(
			"Invalid salary.",
		)

		return
	}

	request :=
		view.SalaryUpdateRequest{
			EmployeeID: employeeID,
			NewSalary:  newSalary,
		}

	response, err :=
		c.service.UpdateSalary(
			context.Background(),
			request,
		)

	if err != nil {

		fmt.Println()
		fmt.Println(
			"Salary update failed:",
			err,
		)

		return
	}

	fmt.Println()
	fmt.Println(
		"Salary updated successfully.",
	)

	fmt.Println()
	fmt.Println(
		"--------------------------------------",
	)

	fmt.Printf(
		"Employee ID : %d\n",
		response.Employee.ID,
	)

	fmt.Printf(
		"Name        : %s\n",
		response.Employee.Name,
	)

	fmt.Printf(
		"Department  : %s\n",
		response.Employee.Department,
	)

	fmt.Printf(
		"Old Salary  : %.2f\n",
		response.SalaryHistory.OldSalary,
	)

	fmt.Printf(
		"New Salary  : %.2f\n",
		response.SalaryHistory.NewSalary,
	)

	fmt.Printf(
		"Changed At  : %s\n",
		response.SalaryHistory.ChangedAt.Format(
			"2006-01-02 15:04:05",
		),
	)

	fmt.Println(
		"--------------------------------------",
	)
}

// ----------------------------------------------------
// View Salary History
// ----------------------------------------------------

func (c *SalaryController) viewSalaryHistory() {

	fmt.Println()
	fmt.Println("======================================")
	fmt.Println("          SALARY HISTORY")
	fmt.Println("======================================")

	employeeIDInput :=
		c.readString(
			"Enter employee ID: ",
		)

	employeeID, err :=
		strconv.ParseInt(
			employeeIDInput,
			10,
			64,
		)

	if err != nil {

		fmt.Println(
			"Invalid employee ID.",
		)

		return
	}

	response, err :=
		c.service.GetSalaryHistory(
			context.Background(),
			employeeID,
		)

	if err != nil {

		fmt.Println()
		fmt.Println(
			"Unable to get salary history:",
			err,
		)

		return
	}

	fmt.Println()

	if len(response.History) == 0 {

		fmt.Println(
			"No salary history found.",
		)

		return
	}

	fmt.Println(
		"==========================================================================",
	)

	fmt.Printf(
		"%-6s %-15s %-15s %-25s\n",
		"ID",
		"OLD SALARY",
		"NEW SALARY",
		"CHANGED AT",
	)

	fmt.Println(
		"==========================================================================",
	)

	for _, history := range response.History {

		fmt.Printf(
			"%-6d %-15.2f %-15.2f %-25s\n",
			history.ID,
			history.OldSalary,
			history.NewSalary,
			history.ChangedAt.Format(
				"2006-01-02 15:04:05",
			),
		)
	}

	fmt.Println(
		"==========================================================================",
	)
}

// ----------------------------------------------------
// Read Input
// ----------------------------------------------------

func (c *SalaryController) readString(
	message string,
) string {

	fmt.Print(message)

	value, err :=
		c.reader.ReadString('\n')

	if err != nil {
		return ""
	}

	return strings.TrimSpace(value)
}
