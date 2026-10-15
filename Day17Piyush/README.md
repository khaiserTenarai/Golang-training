# Day 17 — Advanced GraphQL

No PostgreSQL. In-memory store. REST + GraphQL on the same server.

## Run
```bash
go mod tidy
go run main.go
```
- GraphQL Playground: http://localhost:8080/graphql (open in browser)
- REST API: http://localhost:8080/api/employees

## Tests (Tasks 12, 13)
```bash
go test -v ./tests/...
```

## Task Mapping

| Task | What | Where |
|------|------|-------|
| 1 | Schema modules | schema/types.go, schema/schema.go |
| 2 | Resolver packages | schema/schema.go (queries + mutations organized) |
| 3 | Inject services | schema/schema.go (BuildSchema takes services) |
| 4 | Nested Employee→Department | schema/schema.go (department field resolver) |
| 5 | Optimize nested queries | service/department_service.go (GetByIDs batch) |
| 6 | Authentication | service/auth_service.go + JWT in context |
| 7 | Authorization | schema/schema.go (requireRole helper) |
| 8 | Validate input | service/employee_service.go (Validate method) |
| 9 | Pagination | employees query (page, page_size args) |
| 10 | Filtering | employees query (name, min_salary, max_salary) |
| 11 | Sorting | employees query (sort_by, sort_order) |
| 12 | Test queries | tests/graphql_test.go (TestQuery*) |
| 13 | Test mutations | tests/graphql_test.go (TestMutation*) |
| 14 | REST + GraphQL | main.go (both endpoints on same server) |
| 15 | Mini Project | Everything combined |

## Seeded Users
| Username | Password | Role |
|----------|----------|------|
| admin | admin123 | admin (full access) |
| manager | manager123 | manager (CRUD) |
| piyush | piyush123 | user (read only) |

## GraphQL Examples

### Login
```graphql
mutation {
  login(username: "admin", password: "admin123") {
    token username role
  }
}
```

### List Employees (with pagination, sorting, nested department)
```graphql
{
  employees(page: 1, page_size: 3, sort_by: "salary", sort_order: "desc") {
    employees {
      id name salary
      department { name location }
    }
    total total_pages
  }
}
```

### Create Employee (admin/manager only)
```graphql
mutation {
  createEmployee(input: {
    name: "Frank", email: "frank@test.com",
    age: 30, department_id: 1, salary: 70000
  }) {
    id name email
  }
}
```

### Delete Employee (admin only)
```graphql
mutation { deleteEmployee(id: 3) }
```

Set `Authorization: Bearer <token>` header for protected queries.
