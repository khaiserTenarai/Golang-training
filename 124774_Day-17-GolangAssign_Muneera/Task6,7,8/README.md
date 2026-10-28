# Employee Management GraphQL API

## Features

- GraphQL employee queries and mutations
- Bearer-token authentication
- ADMIN-only employee mutations
- Employee input validation
- PostgreSQL database using GORM
- Request logging

## Setup

1. Install Go and PostgreSQL.
2. Create a PostgreSQL database named employee_db.
3. Configure your database credentials and tokens in .env.
4. Install dependencies:

   go mod tidy

5. Start the application:

   go run .

## Endpoints

- GraphQL: http://localhost:8081/graphql
- Playground: http://localhost:8081/playground

## Authentication

Send this HTTP header with GraphQL requests:

Authorization: Bearer change-this-admin-token

Use the configured ADMIN_TOKEN for mutations.
Use the configured USER_TOKEN for authenticated read queries.

## Example query

query {
  employees {
    id
    name
    email
    department
    salary
  }
}

## Example mutation

mutation {
  createEmployee(input: {
    name: "Muneera"
    email: "muneera@example.com"
    department: "AI"
    salary: 30000
  }) {
    id
    name
    email
    department
    salary
  }
}