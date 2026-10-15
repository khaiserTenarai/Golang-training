# Day 16 - Employee Management: Gin + GORM + PostgreSQL + GraphQL

Built on the reference project (same stack and layout), extended with a real
`Department` entity, input types, variables, filtering/pagination arguments,
a custom scalar and structured GraphQL error handling.

    Client -> Gin -> REST Controller | GraphQL resolver -> Service -> Repository -> GORM -> PostgreSQL

## Layout

    config/      .env loading
    model/       Employee, Department, input + filter structs
    repository/  GORM repositories (+ seed data)
    service/     validation and business rules
    controller/  REST handlers
    middleware/  request logging
    graphql/     schema.go, resolvers.go, scalars.go, handler.go, tests
    apperror/    errors with codes (NOT_FOUND, BAD_USER_INPUT, CONFLICT, INTERNAL)

## Run

    docker compose up -d          # or create the database yourself: CREATE DATABASE employee_db;
    go mod tidy
    go run .

UI: http://localhost:8080/playground (or /graphql). No token or header needed.
Tables are created automatically and 3 departments + 3 employees are seeded.

> Upgrading from the reference project? Its `employees` table has a text
> `department` column that conflicts with the new schema. Use a fresh database
> or run `DROP TABLE employees;` once.

## Tests (no database needed)

    go test ./...

The tests run the full GraphQL stack (schema -> resolvers -> services) against
in-memory repositories.

## Examples

See `graphql/examples.graphql` and `schema.graphql`.

    # variables
    query GetEmployee($id: ID!) { employee(id: $id) { id name department { name } } }
    { "id": "1" }

    mutation Create($input: CreateEmployeeInput!) {
      createEmployee(input: $input) { id name active department { name } }
    }
    { "input": { "name": "John", "email": "john@gmail.com", "salary": 70000, "departmentId": "3" } }

## Errors

    {
      "data": { "employee": null },
      "errors": [{
        "message": "employee with id 99 not found",
        "path": ["employee"],
        "extensions": { "code": "NOT_FOUND" }
      }]
    }

Validation errors add `"field"`. Unexpected failures (e.g. database down) are logged
and returned as a generic `INTERNAL` error. Malformed queries are rejected by GraphQL itself.

## REST (unchanged idea from the reference)

    GET/POST /employees    GET/PUT/DELETE /employees/:id    GET /departments

    curl -X POST localhost:8080/employees -H 'Content-Type: application/json' \
      -d '{"name":"John","email":"john@gmail.com","salary":70000,"departmentId":3}'
