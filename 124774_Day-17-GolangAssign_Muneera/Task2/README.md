# Employee Management GraphQL API — Task 2: Organize Resolver Packages

This project organizes GraphQL resolvers into a separate `graphql/resolvers` package. It uses Gin for the HTTP server, GORM with PostgreSQL for persistence, and a service/repository structure. It does not include REST employee endpoints or a REST controller.

## Folder structure

```text
employee-management/
├── config/
│   └── config.go
├── graphql/
│   ├── handler.go
│   ├── schema.go
│   ├── types.go
│   └── resolvers/
│       ├── employee_query.go
│       └── employee_mutation.go
├── middleware/
│   └── logging.go
├── model/
│   └── employee.go
├── repository/
│   └── repository.go
├── service/
│   └── service.go
├── utility/
│   └── validation.go
├── .env.example
├── .gitignore
├── go.mod
└── main.go
```

## Run the project

1. Install and start PostgreSQL.
2. Create a database named `employee_db` (or set another name in `.env`).
3. Copy `.env.example` to `.env` and update `DB_PASSWORD` and other values for your local PostgreSQL setup.
4. Open a terminal in this folder and run:

```bash
go mod tidy
go run .
```

The API uses port `8081` by default.

- GraphQL endpoint: `http://localhost:8081/graphql`
- Playground: `http://localhost:8081/playground`

## Example query

```graphql
query {
  employees {
    id
    name
    email
    department
    salary
  }
}
```

## Example mutation

```graphql
mutation {
  createEmployee(input: {
    name: "Muneera"
    email: "muneera@example.com"
    department: "Engineering"
    salary: 45000
  }) {
    id
    name
    email
    department
    salary
  }
}
```

## Notes

- This project requires a running PostgreSQL database and valid environment settings.
- `.env` is ignored by Git so credentials are not committed.
- Employee operations are available through GraphQL only; no REST employee routes are included.
