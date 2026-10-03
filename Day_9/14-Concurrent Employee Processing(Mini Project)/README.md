# Concurrent Employee Processing System

## Day 9 - Go Assignment

A simple **Employee Management System** developed using Go and PostgreSQL.

This project is an extension of the **Day 7 Employee Management CRUD project**. The existing CRUD operations are preserved, and **concurrent employee processing** is added using Go concurrency concepts.

---

## Technologies Used

* Go
* PostgreSQL
* pgx/v5
* Goroutines
* Channels
* Select
* WaitGroup
* File Handling
* Layered Architecture

---

## Project Objectives

The main objective of this project is to understand and implement:

1. Goroutines
2. Channels
3. Buffered Channels
4. Producer
5. Consumer
6. Workers
7. Select
8. Backpressure
9. Concurrent Employee Processing
10. Layered Architecture

---

## Project Structure

```text
15-concurrent-employee-processing/
│
├── config/
│   └── config.go
│
├── controller/
│   ├── controller.go
│   └── controller_impl.go
│
├── database/
│   └── database.go
│
├── model/
│   └── employee.go
│
├── repository/
│   ├── repository.go
│   └── repository_impl.go
│
├── service/
│   ├── service.go
│   └── service_impl.go
│
├── utility/
│   └── validation.go
│
├── view/
│   ├── view.go
│   └── view_impl.go
│
├── config.env
├── go.mod
├── go.sum
└── main.go
```

---

## Architecture

The project follows a layered architecture.

```text
                 View
                   ↓
              Controller
                   ↓
                Service
                   ↓
              Repository
                   ↓
               Database
                   ↓
              PostgreSQL
```

The Day 9 concurrent processing is implemented inside the **Service Layer**.

```text
                 Service
                    |
                    ↓
                 Producer
                    |
                    ↓
            Buffered Channel
                    |
          +---------+---------+
          |         |         |
          ↓         ↓         ↓
       Worker 1  Worker 2  Worker 3
          |         |         |
          +---------+---------+
                    ↓
              Employee Processing
```

---

# Existing Day 7 Features

The existing Employee Management CRUD functionality is preserved.

### 1. Save Employee

Adds a new employee to PostgreSQL.

### 2. Find Employee

Finds an employee using employee ID.

### 3. Find All Employees

Retrieves all employees from PostgreSQL.

### 4. Update Employee

Updates employee information.

### 5. Delete Employee

Deletes an employee using employee ID.

---

# Day 9 Concurrent Processing

A new menu option is added:

```text
6. Process Employees Concurrently
```

When this option is selected:

1. Employees are retrieved from PostgreSQL.
2. The Producer sends employees to a channel.
3. The channel temporarily stores employees.
4. Multiple Workers receive employees.
5. Workers process employees concurrently.
6. `select` is used for channel communication.
7. A buffered channel demonstrates backpressure.

---

# Goroutines

Goroutines allow multiple employees to be processed concurrently.

Example:

```go
go s.worker(
    workerID,
    employeeChannel,
    &wg,
)
```

Multiple workers can run at the same time.

For example:

```text
Worker 1 → Employee 1
Worker 2 → Employee 2
Worker 3 → Employee 3
```

---

# Channels

A channel is used to transfer employees between the Producer and Workers.

```go
employeeChannel := make(
    chan model.Employee,
    s.buffer,
)
```

The channel carries:

```text
model.Employee
```

---

# Buffered Channel

The channel has a limited capacity.

The configuration contains:

```text
BUFFER=2
```

Therefore:

```text
Channel Capacity = 2
```

The Producer can temporarily place two employees in the channel.

```text
+-------------------+
| Employee 1        |
| Employee 2        |
+-------------------+
     Capacity = 2
```

---

# Producer

The Producer sends employees into the channel.

```text
Database
    ↓
Repository
    ↓
Service
    ↓
Producer
    ↓
Channel
```

Example:

```go
employeeChannel <- employee
```

The Producer runs as a goroutine.

---

# Consumer

The Workers act as consumers.

They receive employees from the channel:

```go
employee, ok := <-employeeChannel
```

The received employee is then processed.

---

# Workers

The application can create multiple workers.

The configuration contains:

```text
WORKERS=3
```

Therefore, three workers are created:

```text
Worker 1
Worker 2
Worker 3
```

Each worker runs as a separate goroutine.

```go
for i := 1; i <= s.workers; i++ {

    wg.Add(1)

    go s.worker(
        i,
        employeeChannel,
        &wg,
    )
}
```

---

# Select

`select` is used to wait for channel operations.

Example:

```go
select {

case employeeChannel <- employee:

    fmt.Println(
        "Producer sent Employee:",
        employee.ID,
    )

case <-time.After(2 * time.Second):

    fmt.Println(
        "Producer timeout for Employee:",
        employee.ID,
    )
}
```

The Worker also uses `select` while receiving employees.

---

# Backpressure

Backpressure occurs when the Producer is producing data faster than Workers can consume it.

For example:

```text
BUFFER=2
```

The channel can contain only two employees at a time.

```text
Producer
   |
   ↓
+-----------+
| Employee1 |
| Employee2 |
+-----------+
| Buffer=2  |
+-----------+
      |
      ↓
   Workers
```

If the channel is full, the Producer must wait until a Worker consumes an employee.

This prevents the Producer from continuously adding unlimited data to the channel.

---

# WaitGroup

`sync.WaitGroup` is used to wait for all Workers to finish.

```go
var wg sync.WaitGroup
```

Before starting a Worker:

```go
wg.Add(1)
```

When the Worker finishes:

```go
defer wg.Done()
```

Finally:

```go
wg.Wait()
```

This ensures the main processing function waits until all Workers have completed.

---

# Employee Processing

Each Worker validates the employee and performs a simple salary calculation.

The project calculates a **10% salary increment** for demonstration.

Example:

```text
Old Salary = 45000

10% Increment = 4500

New Salary = 49500
```

The database salary is not modified by this processing operation. It is only used to demonstrate concurrent processing.

---

# Configuration

Application configuration is stored in `config.env`.

```text
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=Root
DB_NAME=go_training_db
DB_SSLMODE=disable

WORKERS=3
BUFFER=2
```

The application reads this file using Go file handling.

The project does not use:

* `bufio`
* JSON configuration
* External dotenv libraries

---

# PostgreSQL Database

Create the database:

```sql
CREATE DATABASE go_training_db;
```

Create the employee table:

```sql
CREATE TABLE employees (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    age INT NOT NULL,
    email VARCHAR(100) NOT NULL,
    salary NUMERIC(10,2) NOT NULL
);
```

Insert sample employees:

```sql
INSERT INTO employees (name, age, email, salary) VALUES
('Ganesh', 23, 'ganesh@gmail.com', 45000),
('Ravi', 25, 'ravi@gmail.com', 50000),
('Suresh', 28, 'suresh@gmail.com', 55000),
('Kiran', 24, 'kiran@gmail.com', 48000),
('Arun', 30, 'arun@gmail.com', 60000),
('Vijay', 27, 'vijay@gmail.com', 52000),
('Raj', 29, 'raj@gmail.com', 58000),
('Kumar', 26, 'kumar@gmail.com', 47000);
```

---

# Database Connection

The project uses:

```text
github.com/jackc/pgx/v5
```

and:

```text
pgxpool
```

for PostgreSQL connection pooling.

---

# How to Run

## Step 1: Clone or create the project

Open the project directory:

```powershell
cd 15-concurrent-employee-processing
```

---

## Step 2: Configure PostgreSQL

Make sure PostgreSQL is running.

Update `config.env`:

```text
DB_USER=postgres
DB_PASSWORD=your_password
```

---

## Step 3: Download dependencies

Run:

```powershell
go mod tidy
```

---

## Step 4: Run the application

```powershell
go run .
```

---

# Application Menu

The application displays:

```text
========== Employee Management System ==========

1. Save Employee
2. Find Employee
3. Find All Employees
4. Update Employee
5. Delete Employee
6. Process Employees Concurrently
7. Exit

Enter choice:
```

---

# Example Concurrent Processing

When option `6` is selected:

```text
========== Concurrent Processing ==========

Total Employees: 8
Workers: 3
Channel Buffer: 2

Producer sent Employee: 1
Producer sent Employee: 2
Producer sent Employee: 3

Worker 1 processing Employee 1 - Ganesh
Worker 2 processing Employee 2 - Ravi
Worker 3 processing Employee 3 - Suresh

Worker 1 completed Employee 1
Worker 1 processing Employee 4 - Kiran

Worker 2 completed Employee 2
Worker 2 processing Employee 5 - Arun

Worker 3 completed Employee 3
Worker 3 processing Employee 6 - Vijay

Producer completed.

Worker 1 stopped.
Worker 2 stopped.
Worker 3 stopped.

All employees processed.
```

The exact order may differ because the Workers execute concurrently.

---

# Key Go Concepts Learned

| Concept          | Purpose                                     |
| ---------------- | ------------------------------------------- |
| Goroutine        | Runs tasks concurrently                     |
| Channel          | Transfers data between goroutines           |
| Buffered Channel | Temporarily stores limited data             |
| Producer         | Sends employees to the channel              |
| Consumer         | Receives employees from the channel         |
| Worker           | Processes employees concurrently            |
| Select           | Handles channel communication               |
| Backpressure     | Controls producer when consumers are slower |
| WaitGroup        | Waits for workers to finish                 |

---

# Learning Outcome

After completing this project, you should understand how to:

* Create goroutines
* Create and use channels
* Use buffered channels
* Build a Producer
* Build Consumers/Workers
* Use `select`
* Understand backpressure
* Process multiple employees concurrently
* Use `sync.WaitGroup`
* Combine concurrency with layered architecture
* Connect Go applications to PostgreSQL
* Read configuration from a file
* Extend an existing CRUD application with concurrency

---

## Assignment Summary

```text
Day 7
Employee Management CRUD
        +
Day 9
Concurrent Processing
        ↓
Goroutines
Channels
Select
Producer
Consumer
Workers
Backpressure
        ↓
Concurrent Employee Processing System
```

**Project:** Concurrent Employee Processing System
**Language:** Go
**Database:** PostgreSQL
**Architecture:** Layered Architecture
**Focus:** Go Concurrency
