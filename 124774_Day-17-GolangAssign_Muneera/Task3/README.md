# Employee Management GraphQL API — Task 3: Inject Services into Resolvers

This version demonstrates dependency injection: `main.go` creates the repository and service, `NewServer` injects the service into `EmployeeResolver`, and the schema registers the resolver methods.

## Flow

1. `main.go` connects to PostgreSQL and creates the repository.
2. `service.NewEmployeeService(employeeRepo)` creates the service.
3. `gqlserver.NewServer(employeeService)` passes the service to the GraphQL server.
4. `resolvers.NewEmployeeResolver(svc)` stores the service in the resolver.
5. `schema.go` registers methods such as `r.EmployeeByID` and `r.CreateEmployee`.
6. GraphQL requests call those methods, which use the injected service.

## Run

1. Make sure PostgreSQL is running and create the database configured in `.env`.
2. Copy `.env.example` to `.env`, then set your database credentials.
3. Run:

```bash
go mod tidy
go run .
```

Open `http://localhost:8081/playground`.

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

This project exposes employee operations through GraphQL only; it does not add REST employee endpoints.
