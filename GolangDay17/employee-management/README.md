# Employee Platform - REST + GraphQL + Security + Tests (Day 17)

Gin + GORM + PostgreSQL + GraphQL (graphql-go) + JWT login + role based access + automated tests.

---

## 1. Day 17 checklist - where is each item?

| #  | Topic | Where to look |
|----|-------|---------------|
| 1  | Organize GraphQL schema into modules | `graphql/schema/` - `types.go`, `inputs.go`, `queries.go`, `mutations.go`, `schema.go` |
| 2  | Organize resolver packages | `graphql/resolver/` - one file per topic (`employee_resolver.go`, `department_resolver.go`, `auth_resolver.go` ...) |
| 3  | Inject services into resolvers | `resolver.New(employeeSvc, departmentSvc, authSvc)` in `graphql/handler.go` |
| 4  | Nested Employee -> Department queries | `EmployeeDepartment` resolver (and `DepartmentEmployees` the other way round) |
| 5  | Optimize nested queries | `Preload("Department")` in `repository/employee_repository.go` (see section 7) |
| 6  | Authentication | `auth/` package (JWT + bcrypt), `middleware/auth.go`, `login` mutation, `POST /login` |
| 7  | Authorization | `resolver.AdminOnly(...)` in `schema/mutations.go`, `middleware.RequireRole` in `main.go` |
| 8  | Validate GraphQL input | `resolver/input.go` (ids, types) + `service/employee_service.go` (business rules) + enums |
| 9  | Pagination | `page`, `pageSize` -> `EmployeePage` (`totalCount`, `totalPages`, `hasNext`) |
| 10 | Filtering | `search`, `departmentId`, `minSalary`, `maxSalary` |
| 11 | Sorting | `sortBy` (enum), `sortOrder` (enum) |
| 12 | Test GraphQL queries | `graphql/graphql_test.go` |
| 13 | Test GraphQL mutations | `graphql/graphql_test.go` |
| 14 | Same service through REST and GraphQL | `controller/` and `graphql/resolver/` both call `service.EmployeeService` |
| 15 | Mini project | this whole project |

---

## 2. Project structure

```
employee-management/
├── main.go                      starts everything and connects the layers
├── .env                         settings (database, JWT secret)
├── config/config.go             reads .env
├── model/                       data shapes: Employee, Department, User, EmployeeFilter, EmployeePage
├── repository/                  talks to the database (GORM)
├── service/                     business rules - used by REST AND GraphQL
│   ├── employee_service.go      validation, pagination, sorting
│   ├── department_service.go
│   ├── auth_service.go          login + demo users
│   └── errors.go
├── auth/                        security helpers
│   ├── token.go                 create / check JWT tokens
│   ├── password.go              bcrypt hashing
│   ├── context.go               "who is the current user?"
│   └── auth.go                  roles + errors
├── middleware/
│   ├── logging.go               prints every request
│   └── auth.go                  Auth (401), OptionalAuth, RequireRole (403)
├── controller/                  REST layer
├── graphql/
│   ├── handler.go               builds the HTTP handler (wires services -> resolver -> schema)
│   ├── schema/                  the GraphQL schema, split into modules
│   ├── resolver/                the code behind every GraphQL field
│   ├── examples.graphql         ready-to-run operations
│   └── *_test.go                GraphQL tests
└── *_test.go (in auth, middleware, service)
```

---

## 3. Big picture (how a request travels)

```
Client
  -> Gin
  -> middleware  (logging, read the JWT token)
  -> REST controller   OR   GraphQL schema -> resolver (+ AdminOnly / LoginRequired)
  -> Service      (validation, paging, sorting rules)
  -> Repository   (SQL through GORM)
  -> PostgreSQL
```

The Service is the single place for the rules. REST and GraphQL are only two "doors" into it,
so they always give the same results.

---

## 4. Run it

1. Create the database once: `CREATE DATABASE employee_db;`
2. Put your PostgreSQL password into `.env` (`DB_PASSWORD`).
3. Install the new libraries (JWT + bcrypt - needs internet) and start:

```powershell
go mod tidy
go run .
```

Your existing Day 16 tables work as they are. Only a new `users` table is added.
Two demo users are created automatically on the first start:

| username | password | role  | can do |
|----------|----------|-------|--------|
| admin    | admin123 | ADMIN | read + create + update + delete |
| user     | user123  | USER  | read only |

(Real applications must never hard-code passwords like this - it is only for learning.)

---

## 5. How login works (authentication)

1. Client sends username + password to `POST /login` (or the `login` mutation).
2. The service loads the user and compares the password with the stored **bcrypt hash**.
3. If it matches, the server returns a **JWT token** (signed with `JWT_SECRET`, valid for `JWT_EXPIRY_MINUTES`).
4. The client sends the token in every later request: `Authorization: Bearer <token>`.
5. The middleware checks the signature and expiry and stores the user (id, username, role) in the request.
6. Controllers / resolvers read that user to decide what is allowed.

401 / UNAUTHENTICATED = "who are you?" (no or bad token).
403 / FORBIDDEN = "I know who you are, but you may not do this" (wrong role).

### Who can do what (authorization)

| Operation | not logged in | USER | ADMIN |
|-----------|:---:|:---:|:---:|
| login | yes | yes | yes |
| me, employee, employees, department, departments | no | yes | yes |
| createEmployee, updateEmployee, deleteEmployee, createDepartment | no | no | yes |

---

## 6. Try it

### GraphQL (Playground)

Open http://localhost:8080/playground and follow the steps at the top of `graphql/examples.graphql`:
run `login`, copy the token, then open **HTTP HEADERS** (bottom of the Playground) and paste

```json
{ "Authorization": "Bearer PASTE_YOUR_TOKEN_HERE" }
```

### REST (PowerShell)

```powershell
$login   = Invoke-RestMethod -Method Post -Uri http://localhost:8080/login -ContentType "application/json" -Body '{"username":"admin","password":"admin123"}'
$headers = @{ Authorization = "Bearer $($login.token)" }

# create a department, then an employee (ADMIN)
Invoke-RestMethod -Method Post -Uri http://localhost:8080/departments -Headers $headers -ContentType "application/json" -Body '{"name":"IT","description":"Software team"}'
Invoke-RestMethod -Method Post -Uri http://localhost:8080/employees -Headers $headers -ContentType "application/json" -Body '{"name":"Amit","email":"amit@example.com","salary":75000,"departmentId":1}'

# list with pagination + filtering + sorting
Invoke-RestMethod -Uri "http://localhost:8080/employees?page=1&pageSize=5&search=am&sortBy=salary&sortOrder=desc" -Headers $headers
```

### REST endpoints

| Method | URL | Who |
|--------|-----|-----|
| POST | /login | everyone |
| GET | /employees?page&pageSize&search&departmentId&minSalary&maxSalary&sortBy&sortOrder | logged in |
| GET | /employees/:id | logged in |
| GET | /departments | logged in |
| POST | /employees | ADMIN |
| PUT | /employees/:id | ADMIN |
| DELETE | /employees/:id | ADMIN |
| POST | /departments | ADMIN |

`sortBy` = id, name, email, salary, createdAt. `sortOrder` = asc or desc. `pageSize` max 100.

---

## 7. Optimizing nested queries (the "N+1 problem")

Query: `employees { items { name department { name } } }`

Without optimization: 1 query for the employees + 1 query PER employee for its department.
For 10 employees that is 11 queries; for 100 employees it is 101.

With `Preload("Department")`: GORM loads all needed departments in ONE extra query.
The whole request needs only 3 queries (count, employees, departments), no matter how many employees.

See it yourself: set `DEBUG_SQL=true` in `.env`, restart, run the query and count the SQL lines in the terminal.

---

## 8. Input validation

Three layers protect the API:

1. **GraphQL schema** - required fields (`String!`), correct types, and enums (`sortBy: SALARY`). Bad requests never reach your code.
2. **resolver/input.go** - converts ids and arguments into safe Go values (`"abc"` is not an id).
3. **service** - business rules: name 2-100 characters, valid email, salary 0 - 10,000,000,
   department must exist, `pageSize` 1-100, `minSalary` not above `maxSalary`, sort field must be allowed.

Because the rules live in the service, REST gets exactly the same validation.

---

## 9. Error codes (GraphQL `errors[].extensions.code`)

| code | meaning |
|------|---------|
| UNAUTHENTICATED | no / bad token, or wrong login |
| FORBIDDEN | logged in, but the role is not allowed |
| NOT_FOUND | employee / department does not exist |
| BAD_USER_INPUT | validation failed |
| CONFLICT | duplicate email or department name |
| INTERNAL_SERVER_ERROR | unexpected problem (details only in the server log) |

REST uses 401, 403, 404, 400, 409 and 500 for the same situations.

---

## 10. Tests

```powershell
go test ./...
```

No database is needed - tests use in-memory fake services.

| Package | What is tested |
|---------|----------------|
| `graphql` | queries (nested, pagination, filtering, sorting, variables), mutations (create, update, delete), errors, login, `me`, authentication and authorization |
| `service` | validation rules, default paging values, page-count math, bad paging / sorting input |
| `auth` | token create / parse, wrong secret, expired token, password hashing, role checks |
| `middleware` | 401 without token, 403 for wrong role, 200 for allowed users |

---

## 11. Security notes

- Change `JWT_SECRET` in `.env` before using this anywhere real. Anyone who knows it can forge tokens.
- Passwords are stored as bcrypt hashes, never as plain text.
- Wrong username and wrong password give the same error, so attackers cannot discover usernames.
- `sortBy` only accepts a fixed list, so it cannot be used for SQL injection.
- REST requests are bound to small request structs, so clients cannot set `id` or `createdAt`.
- Internal errors are logged on the server and the client only sees "something went wrong".
- For production also add HTTPS, rate limiting on login, and refresh tokens.
