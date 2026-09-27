// 11. Day 5 Mini Project: Refactor Employee Management into main, controller, service with interface, Dao or repository with interface, model, methods.

package main

import (
	"employee-app/controller"
	"employee-app/repository"
	"employee-app/service"
)

func main() {
	repo := repository.NewMemoryEmployeeRepository()
	svc := service.NewEmployeeService(repo)
	ctrl := controller.NewEmployeeController(svc)

	ctrl.Register(101, "Alice Johnson", "Engineering", 85000)
	ctrl.Register(102, "Bob Smith", "Marketing", 62000)

	ctrl.Show(101)
	ctrl.ListAll()
}
