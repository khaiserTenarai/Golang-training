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

type EmployeeController struct {
	service service.EmployeeService
	reader  *bufio.Reader
}

// Constructor
func NewEmployeeController(
	service service.EmployeeService,
) *EmployeeController {

	return &EmployeeController{
		service: service,
		reader:  bufio.NewReader(os.Stdin),
	}
}

// --------------------------------------------
// Start Application
// --------------------------------------------

func (c *EmployeeController) Start() {

	for {

		fmt.Println()
		fmt.Println("======================================")
		fmt.Println("     EMPLOYEE MANAGEMENT SYSTEM")
		fmt.Println("======================================")
		fmt.Println("1. Search Employees")
		fmt.Println("2. Exit")
		fmt.Println("======================================")

		choice :=
			c.readString(
				"Enter choice: ",
			)

		switch choice {

		case "1":

			c.SearchEmployees()

		case "2":

			fmt.Println()
			fmt.Println(
				"Exiting Employee Management System...",
			)

			return

		default:

			fmt.Println()
			fmt.Println(
				"Invalid choice. Please try again.",
			)
		}
	}
}

// --------------------------------------------
// Search Employees
// --------------------------------------------

func (c *EmployeeController) SearchEmployees() {

	fmt.Println()
	fmt.Println("======================================")
	fmt.Println("          SEARCH EMPLOYEES")
	fmt.Println("======================================")

	// --------------------------------------------
	// Employee Name
	// --------------------------------------------

	name :=
		c.readString(
			"Enter employee name: ",
		)

	// --------------------------------------------
	// Department
	// --------------------------------------------

	department :=
		c.readString(
			"Enter department: ",
		)

	// --------------------------------------------
	// Minimum Salary
	// --------------------------------------------

	salaryMinInput :=
		c.readString(
			"Enter minimum salary (Enter to skip): ",
		)

	var salaryMin *float64

	if salaryMinInput != "" {

		value, err :=
			strconv.ParseFloat(
				salaryMinInput,
				64,
			)

		if err != nil {

			fmt.Println()
			fmt.Println(
				"Invalid minimum salary.",
			)

			return
		}

		salaryMin = &value
	}

	// --------------------------------------------
	// Maximum Salary
	// --------------------------------------------

	salaryMaxInput :=
		c.readString(
			"Enter maximum salary (Enter to skip): ",
		)

	var salaryMax *float64

	if salaryMaxInput != "" {

		value, err :=
			strconv.ParseFloat(
				salaryMaxInput,
				64,
			)

		if err != nil {

			fmt.Println()
			fmt.Println(
				"Invalid maximum salary.",
			)

			return
		}

		salaryMax = &value
	}

	// --------------------------------------------
	// Sorting
	// --------------------------------------------

	fmt.Println()
	fmt.Println("Sort By")
	fmt.Println("1. ID")
	fmt.Println("2. Name")
	fmt.Println("3. Department")
	fmt.Println("4. Salary")

	sortChoice :=
		c.readString(
			"Enter choice: ",
		)

	sortBy, valid :=
		getSortField(sortChoice)

	if !valid {

		fmt.Println()
		fmt.Println(
			"Invalid sort field.",
		)

		return
	}

	// --------------------------------------------
	// Sort Order
	// --------------------------------------------

	fmt.Println()
	fmt.Println("Sort Order")
	fmt.Println("1. Ascending")
	fmt.Println("2. Descending")

	orderChoice :=
		c.readString(
			"Enter choice: ",
		)

	sortOrder, valid :=
		getSortOrder(orderChoice)

	if !valid {

		fmt.Println()
		fmt.Println(
			"Invalid sort order.",
		)

		return
	}

	// --------------------------------------------
	// Page
	// --------------------------------------------

	pageInput :=
		c.readString(
			"Enter page number: ",
		)

	page, err :=
		strconv.Atoi(pageInput)

	if err != nil {

		fmt.Println()
		fmt.Println(
			"Invalid page number.",
		)

		return
	}

	// --------------------------------------------
	// Page Size
	// --------------------------------------------

	pageSizeInput :=
		c.readString(
			"Enter page size: ",
		)

	pageSize, err :=
		strconv.Atoi(pageSizeInput)

	if err != nil {

		fmt.Println()
		fmt.Println(
			"Invalid page size.",
		)

		return
	}

	// --------------------------------------------
	// Create Search Request
	// --------------------------------------------

	request :=
		view.EmployeeSearchRequest{

			Name: name,

			Department: department,

			SalaryMin: salaryMin,

			SalaryMax: salaryMax,

			Page: page,

			PageSize: pageSize,

			SortBy: sortBy,

			SortOrder: sortOrder,
		}

	// --------------------------------------------
	// Call Service
	// --------------------------------------------

	response, err :=
		c.service.SearchEmployees(
			context.Background(),
			request,
		)

	if err != nil {

		fmt.Println()
		fmt.Println(
			"Search failed:",
			err,
		)

		return
	}

	// --------------------------------------------
	// Display Result
	// --------------------------------------------

	c.displayEmployees(response)
}

// --------------------------------------------
// Read Console Input
// --------------------------------------------

func (c *EmployeeController) readString(
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

// --------------------------------------------
// Sort Field
// --------------------------------------------

func getSortField(
	choice string,
) (string, bool) {

	switch choice {

	case "1":
		return "id", true

	case "2":
		return "name", true

	case "3":
		return "department", true

	case "4":
		return "salary", true

	default:
		return "", false
	}
}

// --------------------------------------------
// Sort Order
// --------------------------------------------

func getSortOrder(
	choice string,
) (string, bool) {

	switch choice {

	case "1":
		return "asc", true

	case "2":
		return "desc", true

	default:
		return "", false
	}
}

// --------------------------------------------
// Display Employees
// --------------------------------------------

func (c *EmployeeController) displayEmployees(
	response *view.EmployeeSearchResponse,
) {

	fmt.Println()

	fmt.Println(
		"==============================================================",
	)

	fmt.Printf(
		"%-6s %-22s %-15s %-12s\n",
		"ID",
		"NAME",
		"DEPARTMENT",
		"SALARY",
	)

	fmt.Println(
		"==============================================================",
	)

	if len(response.Data) == 0 {

		fmt.Println(
			"No employees found.",
		)

	} else {

		for _, employee := range response.Data {

			fmt.Printf(
				"%-6d %-22s %-15s %.2f\n",
				employee.ID,
				employee.Name,
				employee.Department,
				employee.Salary,
			)
		}
	}

	fmt.Println(
		"==============================================================",
	)

	fmt.Printf(
		"Page       : %d\n",
		response.Page,
	)

	fmt.Printf(
		"Page Size  : %d\n",
		response.PageSize,
	)

	fmt.Printf(
		"Total      : %d\n",
		response.Total,
	)

	fmt.Printf(
		"Total Pages: %d\n",
		response.TotalPages,
	)

	fmt.Println(
		"==============================================================",
	)
}
