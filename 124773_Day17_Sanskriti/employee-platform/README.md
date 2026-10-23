# Employee Platform — REST + GraphQL

Beginner-friendly Go project using Gin, PostgreSQL/pgx, JWT, and GraphQL. Both APIs share the same service and repository.

## Database

Uses your existing database `emp_db_soap` and tables:

- `departments(id, name)`
- `employees(id, name, email, age, salary, department_id)`

No migration is required if these tables already exist as described.

## Setup (PowerShell)

1. Copy `.env.example` to `.env`.
2. Update `DATABASE_URL` and `JWT_SECRET`. Use a random secret of at least 32 characters.
3. Install dependencies and run:

```powershell
go mod tidy
go run .
```

Server defaults to `http://localhost:8080`.

## Demo login

For local development only, defaults are:
- Admin: `admin` / `ChangeMe123!`
- User: `user` / `ChangeMe123!`

Override with `ADMIN_USERNAME`, `ADMIN_PASSWORD`, `USER_USERNAME`, and `USER_PASSWORD` in `.env`. Replace this demo credential approach with a users table and bcrypt password hashes before production.

## REST examples

Login:

```http
POST http://localhost:8080/auth/login
Content-Type: application/json

{"username":"admin","password":"ChangeMe123!"}
```

Use the returned token as `Authorization: Bearer YOUR_TOKEN`.

```http
GET http://localhost:8080/employees?page=1&limit=10&name=ann&department=IT&sortBy=name&order=ASC
GET http://localhost:8080/employees/1
GET http://localhost:8080/departments
POST http://localhost:8080/employees
Content-Type: application/json
Authorization: Bearer YOUR_TOKEN

{"name":"Anita Sharma","email":"anita@example.com","age":28,"salary":65000,"department_id":1}
```

Admin role is required for create, update, and delete REST routes.

## GraphQL UI

Open `http://localhost:8080/graphql/playground`. The endpoint is `POST http://localhost:8080/graphql`. Log in first and paste the JWT into the UI's headers JSON:

```json
{"Authorization":"Bearer YOUR_TOKEN"}
```

Example query:

```graphql
query {
  employees(page: 1, limit: 10, name: "an", sortBy: "name", order: "ASC") {
    total
    page
    limit
    data {
      id
      name
      email
      age
      salary
      department { id name }
    }
  }
}
```

Example mutation:

```graphql
mutation {
  createEmployee(input: {
    name: "Anita Sharma"
    email: "anita@example.com"
    age: 28
    salary: 65000
    department_id: "1"
  }) {
    id
    name
    department { name }
  }
}
```

## Testing

```powershell
go test ./...
```

## Important security notes

- JWT is signed with `JWT_SECRET`; keep the secret private.
- GraphQL requests can be sent with a bearer token. In a production deployment, enforce role authorization on GraphQL mutations too; the initial scaffold's REST mutations are role-protected but GraphQL mutation role checks should be added before exposing it outside a trusted development environment.
- Demo credentials are for local development only.
- Sorting fields are allowlisted and SQL values are parameterized. Do not concatenate user-controlled sort fields without allowlisting.
- The `graphql/schema/*.graphql` files document the modular schema; the executable schema is currently constructed in `graphql/graphql.go` using graphql-go.
