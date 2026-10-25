package main

/*


# EMS Project Testing — Admin Test Cases

We will test the Employee Management System (EMS) using SoapUI, starting with the admin account only. We will test login, JWT authentication, employee CRUD operations, pagination, validation, and middleware.

Base URL: `http://localhost:10005`

Before testing, make sure PostgreSQL is running, the `employee_db` database is accessible, and your Go server is started with `go run .`.

## Test 1: Admin Login

POST

Login with valid admin credentials

Endpointhttp://localhost:10005/login

`` Headers

| Key          | Value            |
| ------------ | ---------------- |
| Content-Type | application/json |

Request Body — Raw JSON

```
{
  "username": "admin",
  "password": "Admin@12345"
}
```

Expected result: `200 OK`

Expected response structure:

```
{
  "token": "<JWT_TOKEN>",
  "tokenType": "Bearer"
}
```

Copy the actual token returned by your server. Use it in the following tests as `ADMIN_TOKEN`.

## Test 2: Admin Login with Incorrect Password

POST

Invalid password

- Endpoint: `http://localhost:10005/login`
- Header: `Content-Type: application/json`

```
{
  "username": "admin",
  "password": "WrongPassword123"
}
```

Expected: `401 Unauthorized` (assuming your login controller returns 401 for invalid credentials).

## Test 3: Health Check

GET

Check server status

- Endpoint: `http://localhost:10005/health`
- Body: None
- Expected: `200 OK`

## Test 4: Get All Employees

GET

Retrieve employee records

- Endpoint: `http://localhost:10005/employees?page=1&pageSize=10`
- Header: `Authorization: Bearer ADMIN_TOKEN`
- Body: None

Expected: `200 OK`, with employee data and pagination information. Your existing database has employee records, so the response should contain them.

## Test 5: Get Employees — Second Page

- Method: `GET`
- Endpoint: `http://localhost:10005/employees?page=2&pageSize=2`
- Header: `Authorization: Bearer ADMIN_TOKEN`
- Expected: `200 OK`, with up to two employees on the second page if enough records exist.

## Test 6: Get Employee by ID

- Method: `GET`
- Endpoint: `http://localhost:10005/employees/5`
- Header: `Authorization: Bearer ADMIN_TOKEN`
- Expected: `200 OK`, if employee ID `5` exists.

Repeat with `/employees/99999` to test a nonexistent employee. Expected: `404 Not Found`, if your controller maps missing records to 404.

## Test 7: Create an Employee

POST

Create a new employee

- Endpoint: `http://localhost:10005/employees`
- Headers:

| Key           | Value              |
| ------------- | ------------------ |
| Content-Type  | application/json   |
| Authorization | Bearer ADMIN_TOKEN |

Request Body

```
{
  "name": "Admin Test Employee",
  "email": "admin.test.employee@example.com",
  "department": "IT",
  "salary": 55000
}
```

Expected: `201 Created`.

Record the new employee's `id` from the response. Use that ID in Tests 8–10.

## Test 8: Update an Employee

- Method: `PUT`
- Endpoint: `http://localhost:10005/employees/NEW_ID`
- Headers: `Content-Type: application/json` and `Authorization: Bearer ADMIN_TOKEN`

Replace `NEW_ID` with the ID returned by Test 7.

```
{
  "name": "Admin Test Employee Updated",
  "email": "admin.test.employee@example.com",
  "department": "Finance",
  "salary": 60000
}
```

Expected: a successful response if the employee exists and the controller accepts the supplied values.

## Test 9: Create an Employee with Duplicate Email

- Method: `POST`
- Endpoint: `http://localhost:10005/employees`
- Header: `Authorization: Bearer ADMIN_TOKEN` and `Content-Type: application/json`

```
{
  "name": "Duplicate Email Test",
  "email": "admin.test.employee@example.com",
  "department": "IT",
  "salary": 40000
}
```

Expected: `409 Conflict` if your controller maps duplicate-email errors to 409. This email was used in Test 7, so run Test 7 first.

## Test 10: Delete an Employee

- Method: `DELETE`
- Endpoint: `http://localhost:10005/employees/NEW_ID`
- Header: `Authorization: Bearer ADMIN_TOKEN`
- Body: None

Use the employee ID created in Test 7.

Expected: a successful deletion response. Then repeat the GET request for that ID to verify it no longer exists.

## Test 11: Validation — Missing Employee Name

- Method: `POST`
- Endpoint: `http://localhost:10005/employees`
- Headers: `Content-Type: application/json` and `Authorization: Bearer ADMIN_TOKEN`

```
{
  "name": "",
  "email": "missing.name@example.com",
  "department": "IT",
  "salary": 40000
}
```

Expected: `400 Bad Request` if the controller maps service validation errors to HTTP 400.

## Test 12: Validation — Invalid Email

```
{
  "name": "Invalid Email Test",
  "email": "not-an-email",
  "department": "IT",
  "salary": 40000
}
```

- Method: `POST`
- Endpoint: `http://localhost:10005/employees`
- Header: `Authorization: Bearer ADMIN_TOKEN` and `Content-Type: application/json`

Expected: `400 Bad Request`.

## Test 13: Validation — Negative Salary

```
{
  "name": "Negative Salary Test",
  "email": "negative.salary@example.com",
  "department": "IT",
  "salary": -5000
}
```

- Method: `POST`
- Endpoint: `http://localhost:10005/employees`
- Header: `Authorization: Bearer ADMIN_TOKEN` and `Content-Type: application/json`

Expected: `400 Bad Request`.

## Test 14: Invalid Pagination

- Method: `GET`
- Endpoint: `http://localhost:10005/employees?page=0&pageSize=200`
- Header: `Authorization: Bearer ADMIN_TOKEN`

Expected: `400 Bad Request` if your service rejects invalid pagination values.

## Test 15: JWT Authentication

Test these requests using SoapUI:

| Scenario                               | Expected result    |
| -------------------------------------- | ------------------ |
| GET `/employees` without Authorization | `401 Unauthorized` |
| GET `/employees` with an invalid token | `401 Unauthorized` |
| GET `/employees` with the admin token  | `200 OK`           |

## Test 16: Middleware Demonstrations

These routes were included in the proposed project setup:

| Method | Endpoint            | Purpose                        |
| ------ | ------------------- | ------------------------------ |
| GET    | `/middleware-next`  | Demonstrate `Next()` execution |
| GET    | `/middleware-abort` | Demonstrate `Abort()`          |
| GET    | `/panic`            | Test recovery middleware       |

Use `http://localhost:10005` as the base URL. Expected status codes and response bodies depend on the exact handlers and middleware implementations in your project.

## Admin test checklist

Progress

0 of 16 tested

Admin login succeeds

Invalid admin password rejected

Health check succeeds

Get all employees

Pagination works

Get existing employee

Create employee

Update employee

Duplicate email rejected

Delete employee

Missing name rejected

Invalid email rejected

Negative salary rejected

Invalid pagination rejected

Missing/invalid JWT rejected

Middleware demonstrations

Start with Test 1 (Admin Login). Once it succeeds, copy the JWT token and use it in the `Authorization` header for the protected employee endpoints. We can then verify each response in SoapUI before moving to user-role testing.
*/

/*
# EMS Project Testing — User Test Cases

We will now test the normal user account in SoapUI. The user should be able to log in and view employees, but should not be allowed to create, update, or delete employees.

Base URL: `http://localhost:10005`

## Test 1: User Login — Valid Credentials

POST

Login as user

Endpoint

`http://localhost:10005/login`

Headers

| Key          | Value            |
| ------------ | ---------------- |
| Content-Type | application/json |

Request Body — Raw JSON

```
{
  "username": "user",
  "password": "User@12345"
}
```

Expected result: `200 OK`

The response should contain a JWT token. Copy it and use it as `USER_TOKEN` in the following tests.

## Test 2: User Login — Incorrect Password

- Method: `POST`
- Endpoint: `http://localhost:10005/login`
- Header: `Content-Type: application/json`

```
{
  "username": "user",
  "password": "WrongPassword123"
}
```

Expected: `401 Unauthorized`.

## Test 3: User Login — Invalid Username

- Method: `POST`
- Endpoint: `http://localhost:10005/login`
- Header: `Content-Type: application/json`

```
{
  "username": "unknown",
  "password": "User@12345"
}
```

Expected: `401 Unauthorized`.

## Test 4: User Login — Missing Password

- Method: `POST`
- Endpoint: `http://localhost:10005/login`
- Header: `Content-Type: application/json`

```
{
  "username": "user"
}
```

Expected: `400 Bad Request` if the login controller rejects the missing field during JSON binding.

## Test 5: Get All Employees

- Method: `GET`
- Endpoint: `http://localhost:10005/employees`
- Header: `Authorization: Bearer USER_TOKEN`
- Body: None

Expected: `200 OK`, with employee data and pagination information.

## Test 6: Get Employees with Pagination

- Method: `GET`
- Endpoint: `http://localhost:10005/employees?page=1&pageSize=2`
- Header: `Authorization: Bearer USER_TOKEN`

Expected: `200 OK`, with up to two employees and pagination details.

## Test 7: Get Employee by Existing ID

- Method: `GET`
- Endpoint: `http://localhost:10005/employees/5`
- Header: `Authorization: Bearer USER_TOKEN`

Expected: `200 OK` if employee ID `5` exists.

## Test 8: Get Employee by Nonexistent ID

- Method: `GET`
- Endpoint: `http://localhost:10005/employees/99999`
- Header: `Authorization: Bearer USER_TOKEN`

Expected: `404 Not Found` if the controller maps missing employees to HTTP 404.

## Test 9: User Attempts to Create an Employee

POST

Access-control test

- Endpoint: `http://localhost:10005/employees`
- Headers: `Content-Type: application/json` and `Authorization: Bearer USER_TOKEN`

```
{
  "name": "User Create Test",
  "email": "user.create.test@example.com",
  "department": "IT",
  "salary": 40000
}
```

Expected: `403 Forbidden`.

The user is authenticated, but the operation requires the admin role.

## Test 10: User Attempts to Update an Employee

- Method: `PUT`
- Endpoint: `http://localhost:10005/employees/5`
- Headers: `Content-Type: application/json` and `Authorization: Bearer USER_TOKEN`

```
{
  "name": "Unauthorized Update Test",
  "email": "unauthorized.update@example.com",
  "department": "HR",
  "salary": 50000
}
```

Expected: `403 Forbidden`.

Verify that employee ID `5` was not modified by this request.

## Test 11: User Attempts to Delete an Employee

- Method: `DELETE`
- Endpoint: `http://localhost:10005/employees/5`
- Header: `Authorization: Bearer USER_TOKEN`
- Body: None

Expected: `403 Forbidden`.

Verify that employee ID `5` still exists in the database.

## Test 12: Access Employees Without a Token

- Method: `GET`
- Endpoint: `http://localhost:10005/employees`
- Headers: No Authorization header
- Body: None

Expected: `401 Unauthorized`.

## Test 13: Access Employees with an Invalid Token

- Method: `GET`
- Endpoint: `http://localhost:10005/employees`
- Header: `Authorization: Bearer invalid-token-123`

Expected: `401 Unauthorized`.

## Test 14: Invalid Pagination

- Method: `GET`
- Endpoint: `http://localhost:10005/employees?page=0&pageSize=200`
- Header: `Authorization: Bearer USER_TOKEN`

Expected: `400 Bad Request` if your service rejects invalid pagination values.

## Test 15: User Cannot Access Admin Operations

Test all three operations with the user token:

| Method | Endpoint       | Expected result |
| ------ | -------------- | --------------- |
| POST   | `/employees`   | `403 Forbidden` |
| PUT    | `/employees/5` | `403 Forbidden` |
| DELETE | `/employees/5` | `403 Forbidden` |

## User test checklist

Progress

0 of 14 tested

Valid user login

Incorrect password rejected

Invalid username rejected

Missing password rejected

Get all employees

Pagination works

Get existing employee

Get nonexistent employee

Create employee denied

Update employee denied

Delete employee denied

Missing token rejected

Invalid token rejected

Invalid pagination rejected
*/
