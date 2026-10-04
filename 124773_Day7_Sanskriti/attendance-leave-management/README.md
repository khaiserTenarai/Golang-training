# Attendance & Leave Management

Go + PostgreSQL project with one Go file per layer.

## Operations
1. Add Employee
2. Display Employees
3. Employee Check-In
4. Employee Check-Out
5. Attendance Report
6. Apply Leave
7. Display Leave Applications
8. Approve Leave
9. Reject Leave
10. Exit

## Setup

Create the database:
CREATE DATABASE attendance_leave_db;

Connect:
\c attendance_leave_db

Run `schema.sql`.

Update `.env` with your PostgreSQL password.

Then:
go mod tidy
go run .

Attendance prevents duplicate check-in and requires check-in before check-out.
Leave applications start as Pending and can be Approved or Rejected only once.
