# Day 16 — GraphQL Fundamentals (Piyush)

No PostgreSQL. In-memory store with 5 employees + 4 departments.

## Run
```bash
go mod tidy
go run main.go
```
Open http://localhost:8080/graphql in browser for the Playground.

## Task Mapping

| Task | What | File |
|------|------|------|
| 1 | GraphQL schema | resolver/resolver.go (graphql.NewSchema) |
| 2 | Employee type | types/types.go (EmployeeType) |
| 3 | Department type | types/types.go (DepartmentType) |
| 4 | Employee query | resolver/resolver.go ("employee" field) |
| 5 | Employee list query | resolver/resolver.go ("employees" field) |
| 6 | Create mutation | resolver/resolver.go ("createEmployee") |
| 7 | Update mutation | resolver/resolver.go ("updateEmployee") |
| 8 | Delete mutation | resolver/resolver.go ("deleteEmployee") |
| 9 | Scalar types | types/types.go (Int, Float, String, Boolean, NonNull) |
| 10 | Input types | types/types.go (CreateEmployeeInput, UpdateEmployeeInput) |
| 11 | Query arguments | resolver/resolver.go (id, limit, offset, is_active) |
| 12 | Variables | main.go (VariableValues in graphql.Do) |
| 13 | Resolvers | resolver/resolver.go (all Resolve functions) |
| 14 | Error handling | resolver/resolver.go (validation + not-found errors) |
| 15 | Mini Project | Everything combined |

## Quick Test (curl)
```bash
# List employees
curl -X POST http://localhost:8080/graphql -H "Content-Type: application/json" -d '{"query":"{ employees { employees { id name salary department { name } } count } }"}'

# Single employee
curl -X POST http://localhost:8080/graphql -H "Content-Type: application/json" -d '{"query":"{ employee(id: 1) { name email salary } }"}'

# Create (with variables)
curl -X POST http://localhost:8080/graphql -H "Content-Type: application/json" -d '{"query":"mutation($input: CreateEmployeeInput!) { createEmployee(input: $input) { id name } }","variables":{"input":{"name":"Eve","email":"eve@test.com","age":26,"department_id":2,"salary":52000}}}'
```
