package controller
import (
	"fmt"

	"salary-management/service"
	"salary-management/view"
)

type EmployeeControllerImpl struct {
	view          view.EmployeeView
	service       service.EmployeeService
	salaryService service.SalaryService
	salaryView    view.SalaryView
}

func NewEmployeeController(
	view view.EmployeeView,
	service service.EmployeeService,
	salaryService service.SalaryService,
	salaryView view.SalaryView,
) EmployeeControllerInterface {

	return &EmployeeControllerImpl{
		view:          view,
		service:       service,
		salaryService: salaryService,
		salaryView:    salaryView,
	}
}

func (c *EmployeeControllerImpl) Start() {

	for {

		choice := c.view.ShowMenu()

		if choice == 8 {
			fmt.Println("Thank you..")
			return
		}

		c.Process(choice)
	}
}

func (c *EmployeeControllerImpl) Process(
	choice int,
) {

	switch choice {

	case 1:

		employee := c.view.ReadEmployee()

		err := c.service.Save(employee)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println(
			"Employee saved successfully.",
		)

	case 2:

		id := c.view.ReadID()

		employee, err :=
			c.service.FindByID(id)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		c.view.DisplayEmployee(employee)

	case 3:

		employees, err :=
			c.service.FindAll()

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		c.view.DisplayEmployees(employees)

	case 4:

		employee :=
			c.view.ReadEmployeeForUpdate()

		err := c.service.Update(employee)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println(
			"Employee updated successfully.",
		)

	case 5:

		employeeID, salary :=
			c.salaryView.ReadSalaryUpdate()

		err := c.salaryService.UpdateSalary(
			employeeID,
			salary,
		)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println(
			"Salary updated successfully.",
		)

	case 6:

		employeeID :=
			c.salaryView.ReadEmployeeID()

		history, err :=
			c.salaryService.FindHistory(
				employeeID,
			)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		c.salaryView.DisplayHistory(history)

	case 7:

		id := c.view.ReadID()

		err := c.service.Delete(id)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println(
			"Employee deleted successfully.",
		)

	default:

		fmt.Println("Invalid choice.")
	}
}
