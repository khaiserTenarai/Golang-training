package main

import "fmt"

func main() {

	// Step 1: Created the Employee Management System program.
	fmt.Println("Employee Management System")

	// Step 2: Added basic employee management options.

	fmt.Println("1. Add Employee")
	fmt.Println("2. Search Employee")
	fmt.Println("3. Display Employees")
	fmt.Println("4. Delete Employee")

	// Step 3: Modified this file and saved the changes.
	// The changes are now in the Working Tree.

	// Step 4: Checked the Working Tree using:
	// git status


	// Step 5: Added the modified file to the Staging Area using:
	// git add employee.go

	// Step 6: Checked the Staging Area using:
	// git status

	// Step 7: The modified employee.go file is now ready to be committed.


/*

	PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\03-working-tree_vs_staging-area> git init
Initialized empty Git repository in C:/Training/Git/124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy/03-working-tree_vs_staging-area/.git/
PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\03-working-tree_vs_staging-area> git status
On branch master

No commits yet

Untracked files:
  (use "git add <file>..." to include in what will be committed)
        employee.go

nothing added to commit but untracked files present (use "git add" to track)
PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\03-working-tree_vs_staging-area> git add .
PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\03-working-tree_vs_staging-area> git status        
On branch master

No commits yet

Changes to be committed:
  (use "git rm --cached <file>..." to unstage)
        new file:   employee.go

PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\03-working-tree_vs_staging-area> git commit -m "finally i am saving employee.go changes"
[master (root-commit) f10e3ca] finally i am saving employee.go changes
 1 file changed, 31 insertions(+)
 create mode 100644 employee.go
PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\03-working-tree_vs_staging-area> 

*/
}