package main

/*
DAY 15 - ADVANCED GIN
TEST.GO - MANUAL TESTING NOTES

BASE URL:
http://localhost:8081

==================================================
1. TEST GET ALL EMPLOYEES
==================================================

Method: GET
URL: http://localhost:8081/employees

Body:
No body required.

Steps:
1. Open Postman or SoapUI.
2. Select GET method.
3. Enter the URL.
4. Click Send.

Expected Result:
- Status: 200 OK
- Response: List of employees.
- Terminal displays the request method, path, status, and duration.


==================================================
2. TEST CREATE EMPLOYEE
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

Steps:
1. Select POST method.
2. Enter the URL.
3. Select Body -> raw -> JSON.
4. Enter the JSON data.
5. Click Send.

Expected Result:
- Employee should be created successfully.
- Response should contain the created employee details.
- Save the employee ID returned in the response.


==================================================
3. TEST GET EMPLOYEE BY ID
==================================================

Method: GET
URL: http://localhost:8081/employees/1

Note:
Replace 1 with the actual employee ID returned by POST.

Body:
No body required.

Expected Result:
- Status: 200 OK for an existing employee.
- Response contains the requested employee details.


==================================================
4. TEST UPDATE EMPLOYEE
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

Steps:
1. Select PUT method.
2. Enter the URL with the employee ID.
3. Select Body -> raw -> JSON.
4. Enter the updated employee details.
5. Click Send.

Expected Result:
- Employee details should be updated successfully.
- Response should reflect the updated values.


==================================================
5. TEST DELETE EMPLOYEE
==================================================

Method: DELETE
URL: http://localhost:8081/employees/1

Note:
Replace 1 with the actual employee ID.

Body:
No body required.

Steps:
1. Select DELETE method.
2. Enter the URL with the employee ID.
3. Click Send.

Expected Result:
- Employee should be deleted successfully.
- A subsequent GET request for the same ID should return
  the not-found response.


==================================================
6. TEST HEALTH CHECK
==================================================

Method: GET
URL: http://localhost:8081/health

Body:
No body required.

Expected Result:
- Successful response when the database is reachable.
- Error response when the database is unavailable.

Purpose:
Checks whether the application and database are healthy.


==================================================
7. TASK 1 - CUSTOM LOGGING MIDDLEWARE
==================================================

Method: GET
URL: http://localhost:8081/employees

Steps:
1. Send the GET request.
2. Check the terminal where the Go application is running.

Expected Terminal Output:
Method: GET | Path: /employees | Status: 200 | Duration: ...

Purpose:
Demonstrates custom Gin middleware logging:
- HTTP method
- Request path
- Response status code
- Request processing duration


==================================================
8. TASK 3 - DEMONSTRATE NEXT()
==================================================

Method: GET
URL: http://localhost:8081/middleware-next

Body:
No body required.

Steps:
1. Select GET method.
2. Enter the URL.
3. Click Send.
4. Check the response and terminal output.

Expected Response:
{
    "message": "Next() demonstration completed"
}

Expected Terminal Output Order:

Next Middleware: Before Next()
Handler: Executing route handler
Next Middleware: After Next()
Method: GET | Path: /middleware-next | Status: 200 | Duration: ...

Explanation:
1. The middleware starts executing.
2. c.Next() allows the remaining handler to execute.
3. After the handler finishes, execution returns to the middleware.
4. The middleware continues after c.Next().
5. The global logging middleware records the completed request.

Purpose:
Demonstrates how c.Next() executes the remaining handlers
and then returns control to the middleware.


==================================================
9. COMPLETE TESTING ORDER
==================================================

1. Start the application using: go run .
2. Test GET /health.
3. Test POST /employees to create an employee.
4. Test GET /employees to retrieve all employees.
5. Test GET /employees/:id to retrieve one employee.
6. Test PUT /employees/:id to update the employee.
7. Test GET /employees/:id again to verify the update.
8. Test DELETE /employees/:id to delete the employee.
9. Test GET /employees/:id again to verify deletion.
10. Test GET /middleware-next.
11. Verify the middleware execution order in the terminal.
12. Verify that request logs appear for the requests.


==================================================
IMPORTANT NOTES
==================================================

- Keep the Go server running while testing.
- Use the actual employee ID returned by the POST request.
- Set Content-Type to application/json for POST and PUT.
- GET and DELETE requests do not require a request body.
- Expected status codes and response formats depend on
  the implementation in controller/controller.go.
- This file contains manual testing notes only.
- It does not execute automated tests or send HTTP requests.
*/