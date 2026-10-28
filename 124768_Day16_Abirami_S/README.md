# Day 16 - Employee GraphQL API

This project implements all Day 16 GraphQL questions in one layered Go project.

## Structure

day16-employee-graphql/
├── config/
├── database/
├── graphql/
│   ├── server.go
│   └── http_handler.go
├── model/
├── repository/
├── service/
├── schema/
├── queries/
├── go.mod
└── main.go

## Architecture

Client
  ↓
/graphql
  ↓
GraphQL HTTP Handler
  ↓
GraphQL Server / Schema
  ↓
Resolver
  ↓
Service
  ↓
Repository
  ↓
In-memory Database

The GraphQL HTTP UI is provided by github.com/graphql-go/handler with:
- GraphiQL: true
- Playground: true

Open:
http://localhost:8080/graphql

## Run

go mod tidy
go run .

## Day 16 mapping

1. GraphQL schema - graphql/server.go + schema/schema.graphql
2. Employee type - Employee
3. Department type - Department
4. Employee query - employee(id: ID!)
5. Employee list query - employees
6. Create Employee mutation - createEmployee
7. Update mutation - updateEmployee
8. Delete mutation - deleteEmployee
9. Scalar types - ID, String, Int, Float, Boolean
10. Input type - EmployeeInput
11. Query arguments - id and input
12. GraphQL variables - queries/*.graphql
13. Resolvers - resolver functions inside graphql/server.go
14. Error handling - resolver/service errors returned by GraphQL
15. Complete Employee GraphQL API - entire project

## Sample queries

query {
  employees {
    id
    name
    email
    age
    salary
    department {
      id
      name
    }
  }
}

Create with variables:

mutation CreateEmployee($input: EmployeeInput!) {
  createEmployee(input: $input) {
    id
    name
    email
    age
    salary
    department {
      id
      name
    }
  }
}

Variables:

{
  "input": {
    "name": "Priya",
    "email": "priya@example.com",
    "age": 23,
    "salary": 48000,
    "departmentId": "1"
  }
}
