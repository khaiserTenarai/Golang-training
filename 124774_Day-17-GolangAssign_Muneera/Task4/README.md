# Employee Management GraphQL API — Question 4
## Implement nested Employee → Department queries

This version follows the supplied reference structure and keeps the current `Employee` model, repository, service, and GraphQL resolver packages. The reference model stores `Department` as a string, so the GraphQL schema exposes it as a nested object with a `name` field.

### Example nested query

```graphql
query {
  employees {
    id
    name
    email
    department {
      name
    }
    salary
  }
}
```

### Query one employee

```graphql
query {
  employee(id: 1) {
    id
    name
    department {
      name
    }
  }
}
```

### Run

1. Configure PostgreSQL in `.env` using `.env.example` as a template.
2. Ensure PostgreSQL is running and the configured database exists.
3. Run:

```bash
go mod tidy
go run .
```

Open `http://localhost:8081/playground`.

### Important note

This nested GraphQL response follows the supplied reference model, where `Employee.Department` is a string. It wraps that value as `{ name: "..." }` in GraphQL; it does not add a separate Department database table or foreign-key relationship.
