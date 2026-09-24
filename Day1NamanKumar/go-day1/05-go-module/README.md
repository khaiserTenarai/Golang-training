# Task 05 — Go Module

Created with:
```bash
mkdir 05-go-module && cd 05-go-module
go mod init employeeapp
mkdir employee   # separate package
```
The import path of a local package is **module name + folder**: `employeeapp/employee`.

Run: `go run .` → `Employee #1: Asha Rao (Salary: 85000.00)`

Only identifiers starting with a **capital letter** (`New`, `Employee`, `Describe`) are exported and visible from `main`.
