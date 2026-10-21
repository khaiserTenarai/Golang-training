package main

/*
=========================================================
DAY 15 - ADVANCED GIN
TESTING NOTES
=========================================================

SERVER URL:
http://localhost:8081

TOOL:
Postman or SoapUI

---------------------------------------------------------
TASK 1: CUSTOM LOGGING MIDDLEWARE
---------------------------------------------------------

METHOD: GET
URL: http://localhost:8081/employees

PURPOSE:
To verify that the logging middleware records request
details in the terminal.

EXPECTED:
- Status: 200 OK (if the request succeeds)
- Terminal displays:
  Request ID
  HTTP Method
  Request Path
  Response Status
  Request Duration

---------------------------------------------------------
TASK 2: MIDDLEWARE EXECUTION ORDER
---------------------------------------------------------

Use this endpoint if the order demonstration is
registered in main.go.

METHOD: GET
URL: http://localhost:8081/middleware-order

PURPOSE:
To observe the order in which middleware and the
route handler execute.

EXPECTED:
The terminal displays log messages in the order
defined by your middleware and route handler.

NOTE:
This endpoint requires the middleware-order route
and its corresponding middleware to be registered.

---------------------------------------------------------
TASK 3: NEXT() MIDDLEWARE
---------------------------------------------------------

METHOD: GET
URL: http://localhost:8081/middleware-next

PURPOSE:
To demonstrate that execution continues to the
next handler when c.Next() is called.

EXPECTED TERMINAL ORDER:

Next Middleware: Before Next()
Handler: Executing route handler
Next Middleware: After Next()

EXPECTED RESPONSE:

{
  "message": "Next() demonstration completed"
}

STATUS:
200 OK

---------------------------------------------------------
TASK 4: ABORT() MIDDLEWARE
---------------------------------------------------------

METHOD: GET
URL: http://localhost:8081/middleware-abort

PURPOSE:
To demonstrate that c.Abort() prevents the remaining
handlers in the chain from executing.

EXPECTED RESPONSE:

{
  "message": "Request aborted by middleware"
}

STATUS:
403 Forbidden

EXPECTED:
The route handler message
"Handler: This should not execute"
must not appear in the terminal.

---------------------------------------------------------
TASK 5: REQUEST ID MIDDLEWARE
---------------------------------------------------------

METHOD: GET
URL: http://localhost:8081/employees

PURPOSE:
To verify that every request receives a unique ID.

STEPS:
1. Send the request.
2. Open the response Headers section.
3. Find the X-Request-ID header.
4. Send the request again.
5. Verify that the new request has a different ID.

EXPECTED RESPONSE HEADER:

X-Request-ID: <generated-request-id>

EXPECTED TERMINAL LOG:

Request ID: <generated-request-id> |
Method: GET |
Path: /employees |
Status: 200 |
Duration: <duration>

NOTE:
The actual ID and duration will differ for each request.

---------------------------------------------------------
TASK 6: RECOVERY MIDDLEWARE
---------------------------------------------------------

METHOD: GET
URL: http://localhost:8081/panic

PURPOSE:
To test whether the server recovers from a panic.

The /panic route deliberately executes panic().

EXPECTED RESPONSE:

{
  "error": "Internal Server Error"
}

STATUS:
500 Internal Server Error

EXPECTED TERMINAL:
Panic recovered: testing recovery middleware

EXPECTED:
The server should continue running after recovering.

VERIFY:
Send this request after the panic test:

METHOD: GET
URL: http://localhost:8081/health

EXPECTED:
200 OK

---------------------------------------------------------
EMPLOYEE CRUD API TESTING
---------------------------------------------------------

1. CREATE EMPLOYEE
------------------

METHOD: POST
URL: http://localhost:8081/employees

BODY TYPE:
JSON

EXAMPLE BODY:

{
  "name": "Sujith",
  "email": "sujith@example.com",
  "department": "IT",
  "salary": 35000
}

EXPECTED:
A successful response if the supplied data passes
your application's validation.

Save the employee ID returned in the response.

2. GET ALL EMPLOYEES
--------------------

METHOD: GET
URL: http://localhost:8081/employees

PURPOSE:
Retrieve all employees.

3. GET EMPLOYEE BY ID
---------------------

METHOD: GET
URL: http://localhost:8081/employees/1

PURPOSE:
Retrieve the employee whose ID is 1.

NOTE:
Replace 1 with an existing employee ID.

4. UPDATE EMPLOYEE
------------------

METHOD: PUT
URL: http://localhost:8081/employees/1

BODY TYPE:
JSON

EXAMPLE BODY:

{
  "name": "Sujith Updated",
  "email": "sujith.updated@example.com",
  "department": "Development",
  "salary": 45000
}

PURPOSE:
Update an existing employee.

NOTE:
Replace 1 with the employee ID you created.

5. DELETE EMPLOYEE
------------------

METHOD: DELETE
URL: http://localhost:8081/employees/1

PURPOSE:
Delete the specified employee.

VERIFY:
Send GET /employees/1 again.
The employee should no longer be found.

---------------------------------------------------------
ERROR HANDLING TESTS
---------------------------------------------------------

1. NON-EXISTENT EMPLOYEE

METHOD: GET
URL: http://localhost:8081/employees/999999

EXPECTED:
404 Not Found if the controller handles missing
employees with HTTP 404.

2. INVALID JSON

METHOD: POST
URL: http://localhost:8081/employees

BODY:

{
  "name":

EXPECTED:
An error response if the controller validates JSON.

3. INVALID EMPLOYEE DATA

METHOD: POST
URL: http://localhost:8081/employees

BODY:

{
  "name": "",
  "email": "invalid-email",
  "department": "",
  "salary": -100
}

EXPECTED:
An error response if these values violate your
application's validation rules.

---------------------------------------------------------
IMPORTANT NOTES
---------------------------------------------------------

1. Start the server using:

   go run .

2. Keep the server terminal open while testing.

3. Use the correct HTTP method for each endpoint.

4. Select JSON as the request body type for POST and PUT.

5. Request IDs are returned in the X-Request-ID header.

6. Recovery should prevent a test panic from terminating
   the server.

7. HTTP status codes depend on your controller logic
   and validation rules.

8. The /middleware-order endpoint works only if you
   registered the corresponding route in main.go.

=========================================================
END OF TESTING NOTES
=========================================================
*/