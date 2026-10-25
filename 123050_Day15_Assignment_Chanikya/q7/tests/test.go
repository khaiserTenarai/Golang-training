package main

/*
=========================================================
DAY 15 - ADVANCED GIN
TESTING NOTES
=========================================================

SERVER:
http://localhost:8081

TOOL:
Postman or SoapUI

---------------------------------------------------------
TASK 1: CUSTOM LOGGING MIDDLEWARE
---------------------------------------------------------

METHOD: GET
URL: http://localhost:8081/employees

PURPOSE:
Verify that request details appear in the terminal.

EXPECTED:
- HTTP method
- Request path
- Response status
- Request duration
- Request ID

---------------------------------------------------------
TASK 2: MIDDLEWARE EXECUTION ORDER
---------------------------------------------------------

NOTE:
Test this only if the middleware-order route is
registered in main.go.

METHOD: GET
URL: http://localhost:8081/middleware-order

PURPOSE:
Observe the order in which middleware and the
route handler execute.

EXPECTED:
Terminal messages appear in the order defined
by the registered middleware.

---------------------------------------------------------
TASK 3: NEXT() MIDDLEWARE
---------------------------------------------------------

METHOD: GET
URL: http://localhost:8081/middleware-next

PURPOSE:
Verify that c.Next() continues to the next handler.

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
Verify that c.Abort() prevents the remaining
handlers from executing.

EXPECTED RESPONSE:

{
  "message": "Request aborted by middleware"
}

STATUS:
403 Forbidden

EXPECTED:
The message below must not appear:

Handler: This should not execute

---------------------------------------------------------
TASK 5: REQUEST ID MIDDLEWARE
---------------------------------------------------------

METHOD: GET
URL: http://localhost:8081/employees

PURPOSE:
Verify that each request receives a unique ID.

STEPS:
1. Send the request.
2. Open the response Headers section.
3. Find X-Request-ID.
4. Send the request again.
5. Verify that the ID changes.

EXPECTED RESPONSE HEADER:

X-Request-ID: <generated-request-id>

EXPECTED TERMINAL:
Request ID: <generated-request-id> |
Method: GET |
Path: /employees |
Status: 200 |
Duration: <duration>

---------------------------------------------------------
TASK 6: RECOVERY MIDDLEWARE
---------------------------------------------------------

METHOD: GET
URL: http://localhost:8081/panic

PURPOSE:
Verify that the server recovers from a panic.

EXPECTED RESPONSE:

{
  "error": "Internal Server Error"
}

STATUS:
500 Internal Server Error

EXPECTED TERMINAL:
Panic recovered: testing recovery middleware

VERIFY:
After the panic test, send:

GET http://localhost:8081/health

Expected: 200 OK

---------------------------------------------------------
TASK 7: CORS MIDDLEWARE
---------------------------------------------------------

METHOD: GET
URL: http://localhost:8081/employees

PURPOSE:
Verify that the server allows requests from
configured frontend origins.

In Postman, add this request header:

Origin: http://localhost:5173

EXPECTED RESPONSE HEADER:

Access-Control-Allow-Origin: http://localhost:5173

Also verify that X-Request-ID is present.

IMPORTANT:
Postman does not enforce browser CORS restrictions.
This test checks the response headers only.

For a browser test, send a request from a frontend
running at http://localhost:5173.

---------------------------------------------------------
EMPLOYEE CRUD TESTING
---------------------------------------------------------

1. CREATE EMPLOYEE
------------------

METHOD: POST
URL: http://localhost:8081/employees

BODY TYPE: JSON

EXAMPLE BODY:

{
  "name": "Sujith",
  "email": "sujith@example.com",
  "department": "IT",
  "salary": 35000
}

EXPECTED:
A successful response if the data passes validation.

Save the employee ID returned by the API.

2. GET ALL EMPLOYEES
--------------------

METHOD: GET
URL: http://localhost:8081/employees

EXPECTED:
List of employees.

3. GET EMPLOYEE BY ID
---------------------

METHOD: GET
URL: http://localhost:8081/employees/1

EXPECTED:
Employee details if ID 1 exists.

Replace 1 with an existing employee ID.

4. UPDATE EMPLOYEE
------------------

METHOD: PUT
URL: http://localhost:8081/employees/1

BODY TYPE: JSON

EXAMPLE BODY:

{
  "name": "Sujith Updated",
  "email": "sujith.updated@example.com",
  "department": "Development",
  "salary": 45000
}

EXPECTED:
Updated employee details if the update succeeds.

5. DELETE EMPLOYEE
------------------

METHOD: DELETE
URL: http://localhost:8081/employees/1

EXPECTED:
A successful deletion response.

VERIFY:
Send GET /employees/1 again.
The employee should no longer be found.

---------------------------------------------------------
ERROR HANDLING TESTS
---------------------------------------------------------

1. EMPLOYEE NOT FOUND

METHOD: GET
URL: http://localhost:8081/employees/999999

EXPECTED:
404 Not Found if handled that way by the controller.

2. INVALID JSON

METHOD: POST
URL: http://localhost:8081/employees

BODY:

{
  "name":

EXPECTED:
An error response if JSON validation is implemented.

3. INVALID EMPLOYEE DATA

METHOD: POST
URL: http://localhost:8081/employees

BODY TYPE: JSON

{
  "name": "",
  "email": "invalid-email",
  "department": "",
  "salary": -100
}

EXPECTED:
An error response if these values violate validation.

---------------------------------------------------------
IMPORTANT NOTES
---------------------------------------------------------

1. Start the server:

   go run .

2. Keep the server terminal open.

3. Use the correct HTTP method for each API.

4. Select JSON as the body type for POST and PUT.

5. Replace example employee IDs with real IDs.

6. Request IDs appear in the X-Request-ID response header.

7. Recovery should prevent a test panic from terminating
   the server.

8. CORS allows only the origins configured in cors.go.

9. The middleware-order endpoint requires its route
   and middleware to be registered in main.go.

10. Actual response codes depend on your controller
    and validation implementation.

=========================================================
END OF TESTING NOTES
=========================================================
*/
