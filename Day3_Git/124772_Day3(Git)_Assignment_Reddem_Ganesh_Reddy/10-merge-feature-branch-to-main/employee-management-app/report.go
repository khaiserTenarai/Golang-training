package main

import "fmt"

func generateReport() {
	fmt.Println("Employee report generated")
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

PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\09-feature-report-branch> git branch 
feature/employee

PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\09-feature-report-branch> git branch
  feature/employee
* feature/report
  main

PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\09-feature-report-branch> git add .

PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\09-feature-report-branch> git commit -m "added output"

[feature/report e0552d9] added output
 1 file changed, 22 insertions(+)

PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\09-feature-report-branch> go mod init employee-management-ap
go: creating new go.mod: module employee-management-ap
go: to add module requirements and sums:
        go mod tidy

		PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\10-merge-feature-branch-to-main> git branch -M main
PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\10-merge-feature-branch-to-main> git branch
* main
PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\10-merge-feature-branch-to-main> git branch feature/employee
PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\10-merge-feature-branch-to-main> git branch feature/report
PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\10-merge-feature-branch-to-main> git branch
  feature/employee
  feature/report
* main
PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\10-merge-feature-branch-to-main> git status
On branch main
nothing to commit, working tree clean
PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\10-merge-feature-branch-to-main> git branch
  feature/employee
  feature/report
* main
PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\10-merge-feature-branch-to-main> git switch -c feature/employee
fatal: a branch named 'feature/employee' already exists
PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\10-merge-feature-branch-to-main> git switch feature/employee
Switched to branch 'feature/employee'
PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\10-merge-feature-branch-to-main> git add employee.go
fatal: pathspec 'employee.go' did not match any files
PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\10-merge-feature-branch-to-main> git commit -m "added employee.go"
On branch feature/employee
Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
        modified:   employee-management-app/employee.go

no changes added to commit (use "git add" and/or "git commit -a")
PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\10-merge-feature-branch-to-main> git init
Reinitialized existing Git repository in C:/Training/Git/124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy/10-merge-feature-branch-to-main/.git/
PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\10-merge-feature-branch-to-main> git add .
PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\10-merge-feature-branch-to-main> git commit -m "added employee"
[feature/employee 2218f88] added employee
 1 file changed, 1 insertion(+), 14 deletions(-)
PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\10-merge-feature-branch-to-main> git switch main
Switched to branch 'main'
PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\10-merge-feature-branch-to-main> git switch feature/employee
Switched to branch 'feature/employee'
PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\10-merge-feature-branch-to-main> git switch main            
Switched to branch 'main'
PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\10-merge-feature-branch-to-main> git status
On branch main
Changes not staged for commit:
  (use "git add/rm <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
        deleted:    employee-management-app/employee.go

no changes added to commit (use "git add" and/or "git commit -a")
PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\10-merge-feature-branch-to-main> git commit -m "deleted employee.go from main branch"
On branch main
Changes not staged for commit:
  (use "git add/rm <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
        deleted:    employee-management-app/employee.go

no changes added to commit (use "git add" and/or "git commit -a")
PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\10-merge-feature-branch-to-main> git add .
PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\10-merge-feature-branch-to-main> git commit -m "deleted employee.go from main branch"
[main dfe25a9] deleted employee.go from main branch
 1 file changed, 20 deletions(-)
 delete mode 100644 employee-management-app/employee.go
PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\10-merge-feature-branch-to-main> git switch feature/report
Switched to branch 'feature/report'
PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\10-merge-feature-branch-to-main> git init 
Reinitialized existing Git repository in C:/Training/Git/124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy/10-merge-feature-branch-to-main/.git/
PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\10-merge-feature-branch-to-main> git add report.go
fatal: pathspec 'report.go' did not match any files
PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\10-merge-feature-branch-to-main> git branch     
  feature/employee
* feature/report
  main
PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\10-merge-feature-branch-to-main> git add report.go
fatal: pathspec 'report.go' did not match any files
PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\10-merge-feature-branch-to-main> cd .\employee-management-ap 
PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\10-merge-feature-branch-to-main\employee-management-app> git branch
  feature/employee
* feature/report
  main
PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\10-merge-feature-branch-to-main\employee-management-app> git add report.go
PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\10-merge-feature-branch-to-main\employee-management-app> git commit -m "added report"
[feature/report 3ea766e] added report
 1 file changed, 7 insertions(+)
 create mode 100644 employee-management-app/report.go
PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\10-merge-feature-branch-to-main\employee-management-app> git switch main
Switched to branch 'main'
PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\10-merge-feature-branch-to-main\employee-management-app> git merge feature/employee
CONFLICT (modify/delete): employee-management-app/employee.go deleted in HEAD and modified in feature/employee.  Version feature/employee of employee-management-app/employee.go left in tree.
Automatic merge failed; fix conflicts and then commit the result.
PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\10-merge-feature-branch-to-main\employee-management-app> git merge feature/report
error: Merging is not possible because you have unmerged files.
hint: Fix them up in the work tree, and then use 'git add/rm <file>'
hint: as appropriate to mark resolution and make a commit.
fatal: Exiting because of an unresolved conflict.
PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\10-merge-feature-branch-to-main\employee-management-app> git rm employee.go
rm 'employee-management-app/employee.go'
PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\10-merge-feature-branch-to-main\employee-management-app> git branch
  feature/employee
  feature/report
* main
PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\10-merge-feature-branch-to-main\employee-management-app> git merge feature/report
fatal: You have not concluded your merge (MERGE_HEAD exists).
Please, commit your changes before you merge.
PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\10-merge-feature-branch-to-main\employee-management-app> git add .
PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\10-merge-feature-branch-to-main\employee-management-app> git commit -m "merged main and feature -employee"
[main 7303dd6] merged main and feature -employee
PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\10-merge-feature-branch-to-main\employee-management-app> git merge feature/report
Merge made by the 'ort' strategy.
 employee-management-app/report.go | 7 +++++++
 1 file changed, 7 insertions(+)
 create mode 100644 employee-management-app/report.go
PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\10-merge-feature-branch-to-main\employee-management-app> git git branch
git: 'git' is not a git command. See 'git --help'.

The most similar command is
        init
PS C:\Training\Git\124772_Day3(Git)_Assignment_Reddem_Ganesh_Reddy\10-merge-feature-branch-to-main\employee-management-app> git branch
  feature/employee
  feature/report
* main
*/