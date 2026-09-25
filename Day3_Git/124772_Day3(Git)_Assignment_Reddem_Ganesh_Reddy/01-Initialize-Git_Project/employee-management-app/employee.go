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

Test my go project :
---------------------

C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddycd 01-Initialize-Git_Project

C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\01-Initialize-Git_Project>cd employee-management-app




C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\01-Initialize-Git_Project\employee-management-app>go mod init 


employee-management-app
go: creating new go.mod: module employee-management-app
go: to add module requirements and sums:
        go mod tidy



		C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\01-Initialize-Git_Project\employee-management-app>go run .
Employee Management System
Employee added successfully
Employee search completed
Displaying employees
Employee deleted successfully


Then ,I am Initializing Git:
------------------------------

C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\01-Initialize-Git_Project\employee-management-app>git init

Initialized empty Git repository in C:/Training/Git/124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy/01-Initialize-Git_Project/employee-management-app/.git/


C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\01-Initialize-Git_Project\employee-management-app>git status
On branch master

No commits yet

Untracked files:
  (use "git add <file>..." to include in what will be committed)
        employee.go
        go.mod
        main.go

nothing added to commit but untracked files present (use "git add" to track)


I want to change branch name :
---------------------------

C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\01-Initialize-Git_Project\employee-management-app>git branch -M main


C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\01-Initialize-Git_Project\employee-management-app>git status
On branch main

No commits yet

Untracked files:
  (use "git add <file>..." to include in what will be committed)
        employee.go
        go.mod
        main.go

nothing added to commit but untracked files present (use "git add" to track)


commit changes :
---------------
C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\01-Initialize-Git_Project\employee-management-app>git commit -m "Day-3 Task-1"
[main (root-commit) b8b94a0] Day-3 Task-1
 3 files changed, 109 insertions(+)
 create mode 100644 employee.go
 create mode 100644 go.mod
 create mode 100644 main.go
 

*/