package main

import "fmt"

func main() {
	fmt.Println("Employee Management System")
//This is for sample purpose

	addEmployee()
	searchEmployee()
	displayEmployees()
	deleteEmployee()
}

/*

PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\08-feature-employee-branch> go mod init employee-management-app
go: creating new go.mod: module employee-management-app
go: to add module requirements and sums:
        go mod tidy
PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\08-feature-employee-branch> git init
Initialized empty Git repository in C:/Training/Git/124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy/08-feature-employee-branch/.git/
PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\08-feature-employee-branch> git add .
warning: in the working copy of 'go.mod', LF will be replaced by CRLF the next time Git touches it
PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\08-feature-employee-branch> git commit -m "version 1"
[master (root-commit) 344da4c] version 1
 3 files changed, 36 insertions(+)
 create mode 100644 employee-management-app/employee.go
 create mode 100644 employee-management-app/main.go
 create mode 100644 go.mod
PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\08-feature-employee-branch> git status
On branch master
nothing to commit, working tree clean


PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\08-feature-employee-branch> git branch -M main
PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\08-feature-employee-branch> git branch
* main



PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\08-feature-employee-branch> git branch feature/employee
PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\08-feature-employee-branch> git branch
  feature/employee
* main
PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\08-feature-employee-branch> git switch feature/employee
M       employee-management-app/main.go
Switched to branch 'feature/employee'
PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\08-feature-employee-branch> git branch
* feature/employee
  main
*/