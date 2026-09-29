package main

import (
	"time"

	"app/controller"
	"app/model"
	"app/repository"
	"app/service"
)

func main() {
	repo := repository.NewMemoryEmployeeRepository()

	svc := service.NewEmployeeService(repo)

	ctrl := controller.NewEmployeeController(svc)

	emp := &model.Employee{
	    Person: model.Person{
	        FirstName: "Indu",
	        LastName:  "Mathi",
	    },
	    ID:       101,
	    Email:    "indu.mathi@example.com",
	    JobTitle: "Developer",
	    Salary:   100000.0,
	    HireDate: time.Now(),
	    IsActive: true,
	}

	ctrl.Create(emp)
	ctrl.Show(101)

	ctrl.ProcessRaise(101, 10.0)
	ctrl.Show(101)

	ctrl.Deactivate(101)
	ctrl.Show(101)
}