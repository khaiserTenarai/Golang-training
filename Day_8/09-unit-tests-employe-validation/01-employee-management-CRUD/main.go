package main

import (
	"fmt"
	"os"

	"ems/controller"
	"ems/database"
	"ems/repository"
	"ems/service"
	"ems/view"
)

func main() {

	db, err := database.Connect()

	if err != nil {
		fmt.Println("Database connection failed:", err)
		os.Exit(1)
	}

	defer db.Close()

	employeeRepository := repository.NewEmployeeRepository(db)

	employeeService := service.NewEmployeeService(
		employeeRepository,
	)

	employeeView := view.NewEmployeeView()

	employeeController := controller.NewEmployeeController(
		employeeView,
		employeeService,
	)

	employeeController.Start()
}

/*Execution and Output :
-----------------------------

PS C:\Training\Go Lang\Day_8\124772_Day_8(Go)_Assignment_Reddem_Ganesh_Reddy\09-unit-tests-employe-validation\01-employee-management-CRUD> go fmt ./...
PS C:\Training\Go Lang\Day_8\124772_Day_8(Go)_Assignment_Reddem_Ganesh_Reddy\09-unit-tests-employe-validation\01-employee-management-CRUD> go test ./...
?       ems     [no test files]
?       ems/config      [no test files]
?       ems/controller  [no test files]
?       ems/database    [no test files]
?       ems/model       [no test files]
?       ems/repository  [no test files]
ok      ems/service     (cached)
ok      ems/utility     (cached)
?       ems/view        [no test files]
PS C:\Training\Go Lang\Day_8\124772_Day_8(Go)_Assignment_Reddem_Ganesh_Reddy\09-unit-tests-employe-validation\01-employee-management-CRUD> go test -cover ./...
        ems             coverage: 0.0% of statements
        ems/config              coverage: 0.0% of statements
        ems/controller          coverage: 0.0% of statements
        ems/database            coverage: 0.0% of statements
?       ems/model       [no test files]
        ems/repository          coverage: 0.0% of statements
ok      ems/service     (cached)        coverage: 38.5% of statements
ok      ems/utility     (cached)        coverage: 18.8% of statements
        ems/view                coverage: 0.0% of statements

		*/