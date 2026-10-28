1. Health check and middleware
GET
1. Health check

http://localhost:8080/healthz
Expected response:
{
  "status": "ok"
}



GET
2. Custom middleware and execution order

http://localhost:8080/demo/middleware


GET
3. Abort middleware

http://localhost:8080/demo/abort
Expected status: 403 Forbidden
{
  "error": "request stopped by Abort middleware"
}



2. Login and JWT authentication
4. Login

URL: http://localhost:8080/login
JSON request body:
{
  "username": "admin",
  "password": "admin123"
}


Expected: a successful login response containing a JWT token. The exact JSON field name depends on your Login handler.

GET
5. Access employees without a token

URL: http://localhost:8080/employees
Expected status: 401 Unauthorized, with an authentication error response.

GET
6. Access employees with a valid token

URL: http://localhost:8080/employees
Add this request header:
Authorization: Bearer YOUR_JWT_TOKEN


3. Employee CRUD operations
These examples assume your employee table exists and the request body matches your Employee model.
GET
7. Get all employees

URL: http://localhost:8080/employees
Expected: a JSON list of employees.
[]


The empty list is an example if no records exist.

GET
8. Pagination

URL: http://localhost:8080/employees?page=1&size=5
Expected: up to five employee records for the first page.

GET
9. Filter by department

URL: http://localhost:8080/employees?department=IT
Expected: employees belonging to the IT department.

GET
10. Get employee by ID

URL: http://localhost:8080/employees/1
Expected if employee 1 exists:
{
  "id": 1,
  "name": "Ravi",
  "email": "ravi@example.com",
  "department": "IT",
  "salary": 35000
}


If that employee doesn't exist, expect a not-found response.

POST
11. Create an employee

URL: http://localhost:8080/employees
JSON body:
{
  "name": "Ravi",
  "email": "ravi@example.com",
  "department": "IT",
  "salary": 35000
}


Expected: a success response containing the created employee, usually with a generated ID.

PUT
12. Update employee

URL: http://localhost:8080/employees/1
JSON body:
{
  "name": "Ravi Kumar",
  "email": "ravi@example.com",
  "department": "IT",
  "salary": 40000
}


Expected: a success response if employee 1 exists and validation passes.

DELETE
13. Delete employee

URL: http://localhost:8080/employees/1
Expected: a success response if employee 1 exists. A repeated delete may return 404 Not Found.

4. File upload
POST
14. Upload a file

URL: http://localhost:8080/upload