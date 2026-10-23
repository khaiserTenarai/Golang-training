# Employee Management API — Gin + GORM + PostgreSQL + GraphQL

This project demonstrates a layered Go API with REST endpoints and a GraphQL API.

## Requirements covered

1. **Employee type** — GraphQL `Employee` object.
2. **Department type** — GraphQL `Department` object (`name` field); Employee's `department` field returns this object.
3. **Employee query** — `employee(id: ID!)`.
4. **Employee list query** — `employees`.
5. **Create mutation** — `createEmployee(input: EmployeeInput!)`.
6. **Update mutation** — `updateEmployee(id: ID!, input: EmployeeInput!)`.
7. **Delete mutation** — `deleteEmployee(id: ID!)`.
8. **GraphQL scalar types** — `ID`, `String`, `Float`, and `Boolean` (and GraphQL list/non-null wrappers).
9. **Input type** — `EmployeeInput`.
10. **Query arguments** — IDs and input arguments are declared in the schema.
11. **GraphQL variables** — examples with variables and Variables JSON are in `graphql/examples.graphql`.
12. **Resolvers** — query and mutation resolver functions are in `graphql/schema.go`.
13. **Error handling** — invalid IDs and inputs, missing employees, and service/database failures return GraphQL errors.

## Project architecture

Client → Gin → REST Controller / GraphQL Resolver → Service → Repository → GORM → PostgreSQL

## Database setup

Create the database in PostgreSQL:

```sql
CREATE DATABASE employee_db;
```

Edit `.env` and set your local PostgreSQL username/password as needed. The supplied file uses `localhost:5432`, database `employee_db`, and server port `8080`. GORM automatically migrates the Employee table when the application starts.

## Run the program (Windows PowerShell)

1. Extract the ZIP file.
2. Open PowerShell and change to the extracted project directory (the directory containing `go.mod`):

```powershell
cd employee-management-gin-gorm-postgresql-graphql-simple-ready-to-run
```

3. Check `.env` and make sure PostgreSQL is running and `employee_db` exists.
4. Download dependencies and start the server:

```powershell
go mod tidy
go run .
```

5. To run tests:

```powershell
go test ./...
```

## Links to open after starting the server

- **GraphQL endpoint / GraphiQL UI:** http://localhost:8080/graphql
- **GraphQL Playground UI:** http://localhost:8080/playground
- **REST employee list:** http://localhost:8080/employees
- **REST employee by ID:** http://localhost:8080/employees/1

The GraphQL UI runs locally with the server; these links will not respond until `go run .` is running successfully. No authentication token is configured in this sample.

## GraphQL examples

Open `http://localhost:8080/playground` and paste an operation into the query panel. For operations that use variables, enter the corresponding JSON in the **Query Variables** panel.

### List employees

```graphql
query ListEmployees {
  employees {
    id
    name
    email
    department { name }
    salary
  }
}
```

### Fetch one employee (uses variables)

```graphql
query GetEmployee($employeeID: ID!) {
  employee(id: $employeeID) {
    id
    name
    email
    department { name }
    salary
  }
}
```
Query Variables JSON:
```json
{ "employeeID": "1" }
```

### Create employee (uses input type + variables)

```graphql
mutation CreateEmployee($input: EmployeeInput!) {
  createEmployee(input: $input) {
    id
    name
    email
    department { name }
    salary
  }
}
```
Query Variables JSON:
```json
{ "input": { "name": "Amit", "email": "amit@example.com", "department": "IT", "salary": 75000 } }
```

### Update employee (uses arguments + variables)

```graphql
mutation UpdateEmployee($employeeID: ID!, $input: EmployeeInput!) {
  updateEmployee(id: $employeeID, input: $input) {
    id
    name
    email
    department { name }
    salary
  }
}
```
Query Variables JSON:
```json
{ "employeeID": "1", "input": { "name": "Amit Updated", "email": "amit.updated@example.com", "department": "HR", "salary": 85000 } }
```

### Delete employee (uses variables)

```graphql
mutation DeleteEmployee($employeeID: ID!) {
  deleteEmployee(id: $employeeID)
}
```
Query Variables JSON:
```json
{ "employeeID": "1" }
```

## REST endpoints

| Method | Endpoint | Purpose |
|---|---|---|
| GET | `/employees` | List employees |
| POST | `/employees` | Create employee |
| GET | `/employees/:id` | Get employee |
| PUT | `/employees/:id` | Update employee |
| DELETE | `/employees/:id` | Delete employee |

## Notes

- `Department` is represented in GraphQL as an object, while the current GORM model stores its name as a string. This keeps the existing database schema simple while demonstrating GraphQL object types.
- A GraphQL request can return an `errors` array for invalid input, invalid IDs, not-found records, or database/service errors.
- Make sure the `SERVER_PORT` in `.env` matches the port used in the links above.


mutation CreateEmployee {
  createEmployee(input: {
    name: "Sneha Patel"
    email: "sneha2@example.com"
    department: "IT"
    salary: 80000
  }) {
    id
    name
    email
    department {
      name
    }
    salary
  }
}