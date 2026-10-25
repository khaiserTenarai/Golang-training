package main

/*
DAY 15 - ADVANCED GIN
TESTING.GO - MANUAL TESTING NOTES

BASE URL:
http://localhost:8081

==================================================
1. HEALTH CHECK
==================================================

Method: GET
URL: http://localhost:8081/health

Body:
No body required.

Expected Result:
- Status: 200 OK when the database is reachable.
- Response:
  {
      "status": "healthy"
  }

==================================================
2. TASK 1 - CUSTOM LOGGING MIDDLEWARE
==================================================

Method: GET
URL: http://localhost:8081/employees

Body:
No body required.

Steps:
1. Send the request using Postman or SoapUI.
2. Check the terminal.

Expected Terminal Output:
Method: GET | Path: /employees | Status: 200 | Duration: ...

Purpose:
Demonstrates logging of the HTTP method, request path,
response status, and request duration.


==================================================
3. TASK 3 - DEMONSTRATE NEXT()
==================================================

Method: GET
URL: http://localhost:8081/middleware-next

Body:
No body required.

Expected Response:
{
    "message": "Next() demonstration completed"
}

Expected Terminal Log Order:

Next Middleware: Before Next()
Handler: Executing route handler
Next Middleware: After Next()
Method: GET | Path: /middleware-next | Status: 200 | Duration: ...

Explanation:
1. The middleware starts executing.
2. c.Next() allows the remaining handler to execute.
3. Control returns to the middleware after the handler finishes.
4. The middleware continues executing after c.Next().
5. The global logging middleware records the completed request.


==================================================
4. TASK 4 - DEMONSTRATE ABORT()
==================================================

Method: GET
URL: http://localhost:8081/middleware-abort

Body:
No body required.

Expected Response:
{
    "message": "Request aborted by middleware"
}

Expected Terminal Log Order:

Abort Middleware: Before Abort()
Abort Middleware: After Abort()
Method: GET | Path: /middleware-abort | Status: 403 | Duration: ...

Important:
The following message should NOT appear:

Handler: This should not execute

Explanation:
c.Abort() prevents the remaining handlers in the Gin
handler chain from executing. It does not stop the current
middleware, so code after c.Abort() still runs.


==================================================
5. CREATE EMPLOYEE
==================================================

Method: POST
URL: http://localhost:8081/employees

Headers:
Content-Type: application/json

Body (raw JSON):
{
    "name": "Sujith",
    "email": "sujith@example.com",
    "department": "IT",
    "salary": 35000
}

Expected Result:
- Employee should be created successfully.
- Response should contain the employee details and ID.

Save the returned employee ID for the following requests.


==================================================
6. GET ALL EMPLOYEES
==================================================

Method: GET
URL: http://localhost:8081/employees

Body:
No body required.

Expected Result:
- Status: 200 OK when successful.
- Response contains the list of employees.


==================================================
7. GET EMPLOYEE BY ID
==================================================

Method: GET
URL: http://localhost:8081/employees/1

Note:
Replace 1 with the actual employee ID.

Body:
No body required.

Expected Result:
- Employee details should be returned when the ID exists.


==================================================
8. UPDATE EMPLOYEE
==================================================

Method: PUT
URL: http://localhost:8081/employees/1

Note:
Replace 1 with the actual employee ID.

Headers:
Content-Type: application/json

Body (raw JSON):
{
    "name": "Sujith Updated",
    "email": "sujith.updated@example.com",
    "department": "Engineering",
    "salary": 45000
}

Expected Result:
- Employee should be updated successfully.
- Response should reflect the updated values.


==================================================
9. DELETE EMPLOYEE
==================================================

Method: DELETE
URL: http://localhost:8081/employees/1

Note:
Replace 1 with the actual employee ID.

Body:
No body required.

Expected Result:
- Employee should be deleted successfully.
- A subsequent GET request for the same ID should return
  the not-found response.


==================================================
10. COMPLETE TESTING ORDER
==================================================

1. Start the application using: go run .
2. Test GET /health.
3. Test GET /employees.
4. Test GET /middleware-next and verify its log order.
5. Test GET /middleware-abort and verify the handler is skipped.
6. Test POST /employees to create an employee.
7. Test GET /employees/:id.
8. Test PUT /employees/:id.
9. Test GET /employees/:id again to verify the update.
10. Test DELETE /employees/:id.
11. Test GET /employees/:id again to verify deletion.

==================================================
IMPORTANT NOTES
==================================================

- Keep the server running while testing.
- Use the actual employee ID returned by POST.
- POST and PUT require JSON request bodies.
- GET and DELETE do not require request bodies.
- Set Content-Type to application/json for POST and PUT.
- Actual status codes and response bodies depend on your
  controller implementation.
- This file contains manual testing notes only.
- It does not execute automated tests or send HTTP requests.


Purpose of c.Abort()
c.Abort() is used to stop the remaining handlers from executing in Gin.

For example, if a user is not authorized, middleware can stop the request before it reaches the route handler.
Flow of execution
Request: GET /middleware-abort
          |
          v
1. Logging middleware starts
          |
          v
2. AbortDemo() starts
          |
          v
3. Print "Before Abort()"
          |
          v
4. c.Abort()
   Stops the remaining handlers
          |
          v
5. Print "After Abort()"
          |
          v
6. Send 403 Forbidden response
          |
          v
7. Logging middleware logs request details


Important: The route handler does not execute, but the code after c.Abort() inside the current middleware still executes.

In simple words: c.Abort() closes the path for the next handlers, but it does not stop the current middleware immediately.
*/