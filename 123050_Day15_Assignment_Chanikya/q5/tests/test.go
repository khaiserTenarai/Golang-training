package main

/*
DAY 15 - ADVANCED GIN
COMPLETE MANUAL TESTING NOTES
TOOL: POSTMAN / SOAPUI
BASE URL: http://localhost:8081

==================================================
1. HEALTH CHECK
==================================================

Method: GET
URL: http://localhost:8081/health

Body: No body

Expected:
Status: 200 OK
{
    "status": "healthy"
}

If the database is unavailable, an error response is expected.


==================================================
2. TASK 1 - CUSTOM LOGGING MIDDLEWARE
==================================================

Method: GET
URL: http://localhost:8081/employees

Body: No body

Expected:
- Employee list is returned.
- Terminal logs the HTTP method, path, status, and duration.


==================================================
3. TASK 3 - DEMONSTRATE NEXT()
==================================================

Method: GET
URL: http://localhost:8081/middleware-next

Body: No body

Expected Status: 200 OK

Response:
{
    "message": "Next() demonstration completed"
}

Expected terminal log order:

Next Middleware: Before Next()
Handler: Executing route handler
Next Middleware: After Next()
Method: GET | Path: /middleware-next | Status: 200 | Duration: ...

Purpose:
c.Next() executes the remaining handlers and then returns
control to the middleware.


==================================================
4. TASK 4 - DEMONSTRATE ABORT()
==================================================

Method: GET
URL: http://localhost:8081/middleware-abort

Body: No body

Expected Status: 403 Forbidden

Response:
{
    "message": "Request aborted by middleware"
}

Expected terminal log order:

Abort Middleware: Before Abort()
Abort Middleware: After Abort()
Method: GET | Path: /middleware-abort | Status: 403 | Duration: ...

The following message must NOT appear:

Handler: This should not execute

Purpose:
c.Abort() prevents the remaining handlers from executing.


==================================================
5. TASK 5 - REQUEST ID MIDDLEWARE
==================================================

Method: GET
URL: http://localhost:8081/employees

Body: No body

Steps:
1. Send the request.
2. Open the response Headers section.
3. Find X-Request-ID.

Expected Response Header:
X-Request-ID: <unique-request-id>

Each request should receive its own generated ID.


==================================================
6. POST - CREATE EMPLOYEE
==================================================

Method: POST
URL: http://localhost:8081/employees

Headers:
Content-Type: application/json

Body -> raw -> JSON:

{
    "name": "Sujith",
    "email": "sujith@example.com",
    "department": "IT",
    "salary": 35000
}

Expected:
- Employee is created.
- Response contains employee details and ID.

Save the returned ID for the following tests.


==================================================
7. GET - RETRIEVE ALL EMPLOYEES
==================================================

Method: GET
URL: http://localhost:8081/employees

Body: No body

Expected:
- Employee list is returned.
- The newly created employee appears in the list.


==================================================
8. GET - RETRIEVE EMPLOYEE BY ID
==================================================

Method: GET
URL: http://localhost:8081/employees/1

Replace 1 with the actual employee ID.

Body: No body

Expected:
- The matching employee details are returned.
- A nonexistent ID should produce a not-found response.


==================================================
9. PUT - UPDATE EMPLOYEE
==================================================

Method: PUT
URL: http://localhost:8081/employees/1

Replace 1 with the actual employee ID.

Headers:
Content-Type: application/json

Body -> raw -> JSON:

{
    "name": "Sujith Updated",
    "email": "sujith.updated@example.com",
    "department": "Engineering",
    "salary": 45000
}

Expected:
- Employee details are updated.
- Response reflects the new values.


==================================================
10. GET - VERIFY EMPLOYEE UPDATE
==================================================

Method: GET
URL: http://localhost:8081/employees/1

Replace 1 with the actual employee ID.

Body: No body

Expected:
- Updated employee details are returned.
- Name: Sujith Updated
- Department: Engineering
- Salary: 45000


==================================================
11. DELETE - DELETE EMPLOYEE
==================================================

Method: DELETE
URL: http://localhost:8081/employees/1

Replace 1 with the actual employee ID.

Body: No body

Expected:
- Employee is deleted successfully.
- Response indicates successful deletion.


==================================================
12. GET - VERIFY EMPLOYEE DELETION
==================================================

Method: GET
URL: http://localhost:8081/employees/1

Replace 1 with the deleted employee ID.

Body: No body

Expected:
- A not-found response is returned.


==================================================
13. TEST INVALID EMPLOYEE ID
==================================================

Method: GET
URL: http://localhost:8081/employees/999999

Body: No body

Expected:
- A not-found or appropriate error response.


==================================================
14. TEST INVALID EMPLOYEE DATA
==================================================

Method: POST
URL: http://localhost:8081/employees

Headers:
Content-Type: application/json

Body -> raw -> JSON:

{
    "name": "",
    "email": "invalid-email",
    "department": "",
    "salary": -100
}

Expected:
- Validation should reject invalid data if the
  application validation rules cover these fields.


==================================================
15. TEST INVALID JSON
==================================================

Method: POST
URL: http://localhost:8081/employees

Headers:
Content-Type: application/json

Body -> raw -> JSON:

{
    "name": "Sujith",
    "email":

Expected:
- The request should be rejected if JSON parsing
  errors are handled by the controller.


==================================================
16. COMPLETE TESTING ORDER
==================================================

1. Start the application using: go run .
2. GET /health
3. GET /employees
4. GET /middleware-next
5. GET /middleware-abort
6. GET /employees and inspect X-Request-ID.
7. POST /employees
8. GET /employees
9. GET /employees/:id
10. PUT /employees/:id
11. GET /employees/:id to verify the update.
12. DELETE /employees/:id
13. GET /employees/:id to verify deletion.
14. GET /employees/999999 to test an invalid ID.
15. POST /employees with invalid data.
16. POST /employees with invalid JSON.

==================================================
IMPORTANT NOTES
==================================================

- Keep the server running while testing.
- Use the actual employee ID returned by POST.
- POST and PUT require JSON bodies.
- GET and DELETE normally do not require request bodies.
- Set Content-Type to application/json for POST and PUT.
- Exact status codes and response bodies depend on
  the implementation in controller/controller.go.
- These are manual testing notes, not automated tests.
*/