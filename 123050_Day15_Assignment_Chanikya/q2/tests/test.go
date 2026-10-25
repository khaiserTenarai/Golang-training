/*
DAY 15 - GIN MIDDLEWARE
TESTING NOTES USING SOAPUI

BASE URL:
http://localhost:8081

1. HEALTH CHECK
------------------------------
Method: GET
URL: http://localhost:8081/health
Body: No body
Expected Status: 200 OK

Expected Response:

	{
	    "status": "healthy"
	}

2. MIDDLEWARE EXECUTION ORDER
------------------------------
Method: GET
URL: http://localhost:8081/middleware-order
Body: No body
Expected Status: 200 OK

Expected Response:

	{
	    "message": "Middleware execution order demonstrated"
	}

Verify middleware execution order in the Go terminal.

3. GET ALL EMPLOYEES
------------------------------
Method: GET
URL: http://localhost:8081/employees
Body: No body
Expected Status: 200 OK

Expected Response:
An array containing all employees.
The array may be empty if no employees exist.

4. CREATE EMPLOYEE
------------------------------
Method: POST
URL: http://localhost:8081/employees

Header:
Content-Type: application/json

Request Body:

	{
	    "name": "Sujith",
	    "email": "sujith@example.com",
	    "department": "IT",
	    "salary": 35000
	}

Expected Status: 201 Created

Expected Response:
Created employee details with generated ID.

Note: Use a unique email address for each employee.

5. GET EMPLOYEE BY ID
------------------------------
Method: GET
URL: http://localhost:8081/employees/1
Body: No body

Expected Status:
200 OK - Employee exists
404 Not Found - Employee does not exist

Note: Replace 1 with the ID returned by POST.

6. UPDATE EMPLOYEE
------------------------------
Method: PUT
URL: http://localhost:8081/employees/1

Header:
Content-Type: application/json

Request Body:

	{
	    "name": "Sujith Reddy",
	    "email": "sujith.updated@example.com",
	    "department": "Development",
	    "salary": 45000
	}

Expected Status:
200 OK - Employee updated
404 Not Found - Employee does not exist

7. DELETE EMPLOYEE
------------------------------
Method: DELETE
URL: http://localhost:8081/employees/1
Body: No body

Expected Status:
204 No Content - Employee deleted
404 Not Found - Employee does not exist

Expected Response:
No response body.

Note: Deletion permanently removes the employee record.

SOAPUI TESTING STEPS
------------------------------
1. Start the application using: go run .
2. Open SoapUI and create a REST request.
3. Select the HTTP method.
4. Enter the endpoint URL.
5. For POST and PUT, select JSON request format.
6. Add the Content-Type: application/json header.
7. Enter the request body where required.
8. Click Submit Request.
9. Verify the status code and response.
10. Check the Go terminal for middleware logs.

MIDDLEWARE LOG VERIFICATION
------------------------------
Expected log format:

Method: GET | Path: /middleware-order |
Status: 200 | Duration: 1.2ms

Actual duration and status may vary.

IMPORTANT:
- Run POST first to create an employee.
- Use the generated employee ID for GET, PUT, and DELETE.
- The health endpoint requires a working database.
- These are manual testing notes, not automated tests.
*/
package main