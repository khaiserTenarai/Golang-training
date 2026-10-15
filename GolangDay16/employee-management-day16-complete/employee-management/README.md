# Employee Management - Gin + GORM + PostgreSQL + GraphQL (Day 16)

Day 16 - GraphQL Fundamentals mini project: a complete Employee GraphQL API.

## Day 16 checklist and where it lives

| # | Topic | File |
|---|-------|------|
| 1 | GraphQL schema | `graphql/schema.go` (`buildSchema`) |
| 2 | Employee type | `graphql/schema.go` |
| 3 | Department type | `graphql/schema.go`, `model/department.go` |
| 4 | Employee query | `employee(id)` - `resolveEmployee` |
| 5 | Employee list query | `employees(...)` - `resolveEmployees` |
| 6 | Create Employee mutation | `createEmployee` |
| 7 | Update mutation | `updateEmployee` |
| 8 | Delete mutation | `deleteEmployee` |
| 9 | Scalar types | `ID`, `String`, `Int`, `Float`, `Boolean`, `DateTime` |
| 10 | Input types | `EmployeeInput`, `DepartmentInput` |
| 11 | Query arguments | `employees(departmentId, search, limit, offset)` |
| 12 | Variables | `graphql/examples.graphql` |
| 13 | Resolvers | `graphql/resolvers.go` (query, mutation and nested resolvers) |
| 14 | Error handling | `graphql/errors.go` (`extensions.code`) |
| 15 | Mini project | this whole project |

## Architecture

Client
  -> Gin
  -> REST Controller OR GraphQL Resolver
  -> Employee / Department Service (validation)
  -> Employee / Department Repository
  -> GORM
  -> PostgreSQL

## Database

```sql
CREATE DATABASE employee_db;
```

Update `.env` with your PostgreSQL password.

IMPORTANT: if you ran the earlier version of this project, the old `employees`
table has a `department` text column that no longer exists in the model.
Drop it once, then restart so GORM creates the new tables:

```sql
DROP TABLE employees;
```

## Run

```powershell
go mod tidy
go run .
```

Open http://localhost:8080/playground (no token, no header needed).
Ready-made operations are in `graphql/examples.graphql`.

## GraphQL schema (SDL view)

```graphql
type Department { id: ID!  name: String!  description: String  employees: [Employee!]! }
type Employee   { id: ID!  name: String!  email: String!  salary: Float!  createdAt: DateTime!  department: Department! }

input EmployeeInput   { name: String!  email: String!  salary: Float!  departmentId: ID! }
input DepartmentInput { name: String!  description: String }

type Query {
  employee(id: ID!): Employee
  employees(departmentId: ID, search: String, limit: Int = 50, offset: Int = 0): [Employee!]!
  department(id: ID!): Department
  departments: [Department!]!
}
type Mutation {
  createDepartment(input: DepartmentInput!): Department!
  createEmployee(input: EmployeeInput!): Employee!
  updateEmployee(id: ID!, input: EmployeeInput!): Employee!
  deleteEmployee(id: ID!): Boolean!
}
```

## Error codes

| code | meaning |
|------|---------|
| `NOT_FOUND` | employee / department does not exist |
| `BAD_USER_INPUT` | validation failed, bad id, bad limit |
| `CONFLICT` | duplicate email or department name |
| `INTERNAL_SERVER_ERROR` | unexpected error (real cause is only logged on the server) |

## REST endpoints (body now uses `departmentId`)

GET    /employees
POST   /employees        { "name": "Amit", "email": "amit@example.com", "salary": 75000, "departmentId": 1 }
GET    /employees/:id
PUT    /employees/:id
DELETE /employees/:id

## Not included

JWT, authentication, authorization, Swagger, OpenAPI
