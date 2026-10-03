package controller
/*
	EmployeeController defines
	controller operations.
*/
type EmployeeController interface {

	// Starts the application.
	Start()

	// Processes menu choice.
	Process(choice int)
}