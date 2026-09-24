# Day 1 Mini Project — Employee Management CLI

Interactive, menu-driven CLI to **add, search, display and delete** employees
(stored in memory).

## Structure
```
cmd/employeeapp/main.go        entry point + seed data
internal/employee/employee.go  model, validation, sentinel errors
internal/employee/store.go     thread-safe in-memory store (map + RWMutex)
internal/cli/cli.go            menu UI (io.Reader/io.Writer → testable)
*_test.go                      unit + end-to-end tests
Makefile                       run / build / test / fmt / vet
```

## Run
```bash
make run              # or: go run ./cmd/employeeapp
SEED_DATA=false make run   # start empty
make test
```

## Sample session
```
====== Employee Management ======
 1. Add employee
 2. Search employee
 3. Display all employees
 4. Delete employee
 5. Exit
=================================
Choose an option (1-5): 1
Name: Rahul Sharma
Department: Engineering
Salary: 1500000
✔ Employee added with ID 4.

Choose an option (1-5): 3
ID  NAME          DEPARTMENT   SALARY
--  ----          ----------   ------
1   Asha Rao      Engineering  1200000.00
2   Vikram Iyer   Finance      950000.00
3   Meera Nair    HR           800000.00
4   Rahul Sharma  Engineering  1500000.00
(4 record(s))
```

## Input handling
- Invalid menu choices, non-numeric IDs/salaries, empty name/department and
  non-positive salary are all rejected with a clear message.
- Searching or deleting an unknown ID reports "not found".
- Ctrl+D (EOF) exits cleanly.
