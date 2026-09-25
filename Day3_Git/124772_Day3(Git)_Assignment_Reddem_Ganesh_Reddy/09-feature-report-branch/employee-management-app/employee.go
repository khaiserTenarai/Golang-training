package main

import "fmt"

func addEmployee() {
	fmt.Println("Employee added successfully")
}

func searchEmployee() {
	fmt.Println("Employee search completed")
}

func displayEmployees() {
	fmt.Println("Displaying employees")
}

func deleteEmployee() {
	fmt.Println("Employee deleted successfully")
}

/*
PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\09-feature-report-branch> git branch
* master
PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\09-feature-report-branch> git branch -M main
PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\09-feature-report-branch> git branch
* main
PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\09-feature-report-branch> git branch feature/report
PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\09-feature-report-branch> git branch
  feature/report
* main
PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\09-feature-report-branch> git switch feature/report
Switched to branch 'feature/report'
PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\09-feature-report-branch> git branch
* feature/report
  main
PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\09-feature-report-branch> git branch feature/employee
PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\09-feature-report-branch> git branch
  feature/employee
* feature/report
  main
PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\09-feature-report-branch> 
*/
