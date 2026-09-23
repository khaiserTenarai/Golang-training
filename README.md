# Employee Management System
A simple Employee Management System built with Go.

## Features

- Add employee
- Display employees
- Search employees
- Generate employee reports

## Technologies

- Go
- Git
- Github

## How to Run

'''bash
go run .




## My terminal activity throughout the assessment
// I am leaving this here in case you wanted to see how I approached it.
PS C:\124770_Day3_Piyush> mkdir Employee-Management


    Directory: C:\124770_Day3_Piyush


Mode                 LastWriteTime         Length Name                          
----                 -------------         ------ ----                          
d-----        23-09-2026     15:43                Employee-Management           


PS C:\124770_Day3_Piyush> cd Employee-Management
PS C:\124770_Day3_Piyush\Employee-Management> git init
Initialized empty Git repository in C:/124770_Day3_Piyush/Employee-Management/.git/
PS C:\124770_Day3_Piyush\Employee-Management> git status
On branch master

No commits yet

nothing to commit (create/copy files and use "git add" to track)
PS C:\124770_Day3_Piyush\Employee-Management> go mod init Employee-Management
go: creating new go.mod: module Employee-Management
PS C:\124770_Day3_Piyush\Employee-Management> git status
On branch master

No commits yet

Untracked files:
  (use "git add <file>..." to include in what will be committed)
        employee.go
        go.mod
        main.go
        report.go

nothing added to commit but untracked files present (use "git add" to track)
PS C:\124770_Day3_Piyush\Employee-Management> git add .
warning: in the working copy of 'go.mod', LF will be replaced by CRLF the next time Git touches it
PS C:\124770_Day3_Piyush\Employee-Management> git status
On branch master

No commits yet

Changes to be committed:
  (use "git rm --cached <file>..." to unstage)
        new file:   employee.go
        new file:   go.mod
        new file:   main.go
        new file:   report.go

PS C:\124770_Day3_Piyush\Employee-Management> git commit -m "Initial employee management project setup"
[master (root-commit) b3f088f] Initial employee management project setup
 Committer: Piyush Goyal <Piyush.Goyal@tenarai.com>
Your name and email address were configured automatically based
on your username and hostname. Please check that they are accurate.
You can suppress this message by setting them explicitly. Run the
following command and follow the instructions in your editor to edit
your configuration file:

    git config --global --edit

After doing this, you may fix the identity used for this commit with:

    git commit --amend --reset-author

 4 files changed, 4 insertions(+)
 create mode 100644 employee.go
 create mode 100644 go.mod
 create mode 100644 main.go
 create mode 100644 report.go
PS C:\124770_Day3_Piyush\Employee-Management> git status
On branch master
Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
        modified:   employee.go

no changes added to commit (use "git add" and/or "git commit -a")
PS C:\124770_Day3_Piyush\Employee-Management> git add employee.go
PS C:\124770_Day3_Piyush\Employee-Management> git status
On branch master
Changes to be committed:
  (use "git restore --staged <file>..." to unstage)
        modified:   employee.go

PS C:\124770_Day3_Piyush\Employee-Management> git diff
diff --git a/employee.go b/employee.go
index 92b3a7a..951b16c 100644
--- a/employee.go
+++ b/employee.go
@@ -4,4 +4,5 @@ type Employee struct{
        ID int 
        Name string
        Salary float64
+       Department string
 }
\ No newline at end of file
PS C:\124770_Day3_Piyush\Employee-Management> git add employee.go
PS C:\124770_Day3_Piyush\Employee-Management> git commit -m "Add employee structure"
[master 111a054] Add employee structure
 Committer: Piyush Goyal <Piyush.Goyal@tenarai.com>
Your name and email address were configured automatically based
on your username and hostname. Please check that they are accurate.
You can suppress this message by setting them explicitly. Run the
following command and follow the instructions in your editor to edit
your configuration file:

    git config --global --edit

After doing this, you may fix the identity used for this commit with:

    git commit --amend --reset-author

 1 file changed, 8 insertions(+)
PS C:\124770_Day3_Piyush\Employee-Management> git show
commit 111a0542b67eb97e28c83a4506dc591a9b98b52a (HEAD -> master)
Author: Piyush Goyal <Piyush.Goyal@tenarai.com>
Date:   Wed Sep 23 16:11:38 2026 +0530

    Add employee structure

diff --git a/employee.go b/employee.
go
index e69de29..951b16c 100644
--- a/employee.go
+++ b/employee.go
@@ -0,0 +1,8 @@
+package main
+
+type Employee struct{
+       ID int 
+       Name string
+       Salary float64
+       Department string
+}
\ No newline at end of file
PS C:\124770_Day3_Piyush\Employee-Management> git log --oneline
111a054 (HEAD -> master) Add employee structure
b3f088f Initial employee management project setup
PS C:\124770_Day3_Piyush\Employee-Management> git show 111a054
commit 111a0542b67eb97e28c83a4506dc591a9b98b52a (HEAD -> master)
Author: Piyush Goyal <Piyush.Goyal@tenarai.com>
Date:   Wed Sep 23 16:11:38 2026 +0530

    Add employee structure

diff --git a/employee.go b/employee.
PS C:\124770_Day3_Piyush\Employee-Management> git status
On branch master
Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
        modified:   employee.go

no changes added to commit (use "git add" and/or "git commit -a")
PS C:\124770_Day3_Piyush\Employee-Management> git restore employee.go
PS C:\124770_Day3_Piyush\Employee-Management> git status
On branch master
nothing to commit, working tree clean
PS C:\124770_Day3_Piyush\Employee-Management> git switch main
fatal: invalid reference: main
PS C:\124770_Day3_Piyush\Employee-Management> git branch -M main
PS C:\124770_Day3_Piyush\Employee-Management> git branch
* main
PS C:\124770_Day3_Piyush\Employee-Management> git switch -c feature/employee
Switched to a new branch 'feature/employee'
PS C:\124770_Day3_Piyush\Employee-Management> git branch
* feature/employee
  main
PS C:\124770_Day3_Piyush\Employee-Management> git status
On branch feature/employee
nothing to commit, working tree clean
PS C:\124770_Day3_Piyush\Employee-Management> git branch
* feature/employee
  main
PS C:\124770_Day3_Piyush\Employee-Management> git log --oneline
111a054 (HEAD -> feature/employee, main) Add employee structure
b3f088f Initial employee management project setup
PS C:\124770_Day3_Piyush\Employee-Management> git status
On branch feature/employee
Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
        modified:   employee.go

no changes added to commit (use "git add" and/or "git commit -a")
PS C:\124770_Day3_Piyush\Employee-Management> git add employee.go
PS C:\124770_Day3_Piyush\Employee-Management> git status
On branch feature/employee
Changes to be committed:
  (use "git restore --staged <file>..." to unstage)
        modified:   employee.go

PS C:\124770_Day3_Piyush\Employee-Management> git commit -m "Add employee display functionality"
[feature/employee 90ec2a4] Add employee display functionality
 Committer: Piyush Goyal <Piyush.Goyal@tenarai.com>
Your name and email address were configured automatically based
on your username and hostname. Please check that they are accurate.
You can suppress this message by setting them explicitly. Run the
following command and follow the instructions in your editor to edit
your configuration file:

    git config --global --edit

After doing this, you may fix the identity used for this commit with:

    git commit --amend --reset-author

 1 file changed, 9 insertions(+)
PS C:\124770_Day3_Piyush\Employee-Management> git log --oneline --all --decorate
90ec2a4 (HEAD -> feature/employee) Add employee display functionality
111a054 (main) Add employee structure
b3f088f Initial employee management project setup
PS C:\124770_Day3_Piyush\Employee-Management> git switch main
Switched to branch 'main'
PS C:\124770_Day3_Piyush\Employee-Management> git status
On branch main
nothing to commit, working tree clean
PS C:\124770_Day3_Piyush\Employee-Management> git switch -c feature/report
Switched to a new branch 'feature/report'
PS C:\124770_Day3_Piyush\Employee-Management> git branch
  feature/employee
* feature/report
  main
PS C:\124770_Day3_Piyush\Employee-Management> git status
On branch feature/report
Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
        modified:   report.go

no changes added to commit (use "git add" and/or "git commit -a")
PS C:\124770_Day3_Piyush\Employee-Management> git diff
diff --git a/report.go b/report.go
index 99fec5a..e89b9f8 100644
--- a/report.go
+++ b/report.go
@@ -1 +1,12 @@
-package employeemanagement
+package main
+
+import "fmt"
+
+func GenerateReport(employees []Emp
loyee) {
+       fmt.Println("Employee Report
")
+       fmt.Println("---------------
-")
+
+       for _, employee := range emp
loyees {
+               fmt.Println(employee
PS C:\124770_Day3_Piyush\Employee-Management> git branch  
  feature/employee
* feature/report
  main
PS C:\124770_Day3_Piyush\Employee-Management> git add report.go
PS C:\124770_Day3_Piyush\Employee-Management> git status
On branch feature/report
Changes to be committed:
  (use "git restore --staged <file>..." to unstage)
        modified:   report.go

PS C:\124770_Day3_Piyush\Employee-Management> git commit -m "ADD employee report functionality"
[feature/report 73e0749] ADD employee report functionality
 Committer: Piyush Goyal <Piyush.Goyal@tenarai.com>
Your name and email address were configured automatically based
on your username and hostname. Please check that they are accurate.
You can suppress this message by setting them explicitly. Run the
following command and follow the instructions in your editor to edit
your configuration file:

    git config --global --edit

After doing this, you may fix the identity used for this commit with:

    git commit --amend --reset-author

 1 file changed, 12 insertions(+), 1 deletion(-)
PS C:\124770_Day3_Piyush\Employee-Management> git log --oneline --all --decorate --graph
* 73e0749 (HEAD -> feature/report) ADD employee report functionality
| * 90ec2a4 (feature/employee) Add employee display functionality
|/  
* 111a054 (main) Add employee structure
* b3f088f Initial employee management project setup
PS C:\124770_Day3_Piyush\Employee-Management> 