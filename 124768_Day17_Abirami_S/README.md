# Day 17 - Employee Management Platform

Beginner-friendly Employee Management project using:

- Go
- Gin
- GORM
- PostgreSQL
- GraphQL
- JWT authentication
- REST + GraphQL using the same service layer

## Architecture

```text
REST API ──────┐
               ↓
        EmployeeService
               ↓
       EmployeeRepository
               ↓
        GORM + PostgreSQL
               ↑
GraphQL API ───┘
```

Both REST and GraphQL call the **same service layer**.

## Project structure

```text
day17-employee-platform/
├── config/
│   └── config.go
├── database/
│   └── database.go
├── model/
│   └── employee.go
├── repository/
│   ├── employee_repository.go
│   └── department_repository.go
├── service/
│   ├── employee_service.go
│   └── department_service.go
├── middleware/
│   ├── auth.go
│   └── jwt.go
├── handler/
│   ├── auth_handler.go
│   └── employee_handler.go
├── graphql/
│   ├── server.go
│   ├── types.go
│   ├── query.go
│   ├── mutation.go
│   └── examples.txt
├── tests/
│   ├── graphql_test.go
│   └── rest_test.go
├── main.go
├── .env.example
├── README.md
└── go.mod
```

## 1. Create PostgreSQL database

Create this database in PostgreSQL / pgAdmin:

```sql
CREATE DATABASE day17_employee_db;
```

You do **not** need to create the `employees` or `departments` tables manually. GORM creates them with `AutoMigrate()`.

## 2. Configure `.env`

Copy `.env.example` to `.env`.

```text
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=Root
DB_NAME=day17_employee_db
SERVER_PORT=8080
JWT_SECRET=day17-secret-key
```

Change `DB_PASSWORD` if your PostgreSQL password is different.

## 3. Install dependencies

From the project folder:

```powershell
go mod tidy
```

## 4. Run the project

```powershell
go run .
```

The application uses **one Gin server**:

```text
REST:    http://localhost:8080/employees
Login:   http://localhost:8080/login
GraphQL: http://localhost:8080/graphql
```

## Demo users

Admin:

```text
username: admin
password: admin123
role: admin
```

Normal user:

```text
username: user
password: user123
role: user
```

## Login

Send:

```text
POST http://localhost:8080/login
Content-Type: application/json
```

Body:

```json
{
  "username": "admin",
  "password": "admin123"
}
```

The response contains a JWT token.

## REST APIs

### Public GET

```text
GET http://localhost:8080/employees
GET http://localhost:8080/employees/1
```

Pagination:

```text
GET http://localhost:8080/employees?page=1&pageSize=5
```

Filtering:

```text
GET http://localhost:8080/employees?name=Alice
```

Sorting:

```text
GET http://localhost:8080/employees?sortBy=name&sortOrder=desc
```

### Admin-only write operations

Use:

```text
Authorization: Bearer YOUR_ADMIN_TOKEN
```

Then:

```text
POST   http://localhost:8080/employees
PUT    http://localhost:8080/employees/1
DELETE http://localhost:8080/employees/1
```

Example POST body:

```json
{
  "name": "Arun",
  "email": "arun@example.com",
  "salary": 45000,
  "departmentId": 1
}
```

## GraphQL

Open:

```text
http://localhost:8080/graphql
```

GraphiQL is enabled.

### Query employees

```graphql
query {
  employees(page: 1, pageSize: 5) {
    items {
      id
      name
      email
      salary
      department {
        id
        name
      }
    }
    total
    page
    pageSize
    totalPages
  }
}
```

### Query one employee

```graphql
query {
  employee(id: 1) {
    id
    name
    email
    salary
    department {
      id
      name
    }
  }
}
```

### Filter and sort

```graphql
query {
  employees(
    page: 1
    pageSize: 10
    name: "a"
    sortBy: "name"
    sortOrder: "desc"
  ) {
    items {
      id
      name
      email
    }
    total
    totalPages
  }
}
```

### GraphQL mutation

GraphQL mutations require an **admin JWT**.

In GraphiQL, add this header:

```json
{
  "Authorization": "Bearer YOUR_ADMIN_TOKEN"
}
```

Then run:

```graphql
mutation {
  createEmployee(
    input: {
      name: "GraphQL Test"
      email: "graphql-test@example.com"
      salary: 30000
      departmentId: 1
    }
  ) {
    id
    name
    email
    salary
    department {
      id
      name
    }
  }
}
```

The GraphQL HTTP middleware is intentionally optional: a request without a token can still open GraphiQL and execute read-only queries. A valid token adds the user role to the GraphQL context, and mutation resolvers require the `admin` role.

## Testing

Run:

```powershell
go test ./...
```

The GraphQL tests use `gql.Do(...)`, which is the correct execution API for `github.com/graphql-go/graphql`.
