# Employee Microservices (Go)

A beginner-friendly architecture demo using Gin, PostgreSQL, REST, an in-memory Go-channel Event Bus, a simple service registry, a basic API Gateway, and a Saga example. No Kafka is used.

## Important limitations

- This starter runs the gateway and three service HTTP servers in **one Go process** so the in-memory Event Bus can deliver events. It demonstrates logical service boundaries, not independently deployable async messaging.
- The registry is in-memory and not health-aware. For multiple processes/hosts, use a shared registry such as Consul/Kubernetes and a shared durable transport such as a PostgreSQL queue or RabbitMQ.
- The Saga endpoint is educational: it shows orchestration and the failure branch, but it does **not** delete/undo the employee row on failure. A production Saga needs a real, idempotent compensating operation and persisted Saga state.
- This starter does not implement JWT authentication, rate limiting, OpenTelemetry exporters, or durable event delivery.

## Requirements

- Go 1.23+
- PostgreSQL
- Three databases: `employee_db`, `department_db`, `salary_db`

Create databases, for example in `psql`:

```sql
CREATE DATABASE employee_db;
CREATE DATABASE department_db;
CREATE DATABASE salary_db;
```

Update `.env` with your PostgreSQL username/password. The default values are examples.

Run the migrations against their matching databases:

```powershell
psql -d employee_db -f migrations/employee.sql
psql -d department_db -f migrations/department.sql
psql -d salary_db -f migrations/salary.sql
```

Then install dependencies and start the application:

```powershell
go mod tidy
go run .
```

## URLs

Gateway: `http://localhost:8080`

| Method | URL | Purpose |
|---|---|---|
| GET | `/health` | Gateway health |
| GET | `/employees` | List employees |
| POST | `/employees` | Create employee |
| GET | `/departments` | List departments |
| GET | `/departments/1` | Get department |
| GET | `/salaries/101` | List salary rows for employee 101 |
| POST | `/sagas/employee` | Demonstrate Saga orchestration |

Direct service health URLs:
- Employee: `http://localhost:8081/health`
- Department: `http://localhost:8082/health`
- Salary: `http://localhost:8083/health`

## SoapUI examples

Create employee:

```http
POST http://localhost:8080/employees
Content-Type: application/json

{
  "name": "Anita Sharma",
  "email": "anita@example.com",
  "department_id": 1,
  "salary": 65000
}
```

Run Saga success path:

```http
POST http://localhost:8080/sagas/employee
Content-Type: application/json

{
  "employee_name": "Ravi Kumar",
  "email": "ravi@example.com",
  "department_id": 1,
  "fail_department": false
}
```

Run the demo failure path (uses a unique email each time):

```http
POST http://localhost:8080/sagas/employee
Content-Type: application/json

{
  "employee_name": "Test Failure",
  "email": "test-failure-001@example.com",
  "department_id": 1,
  "fail_department": true
}
```

The failure endpoint returns `COMPENSATION_REQUIRED`; it intentionally does not pretend that compensation occurred. Implement an actual delete/release operation before treating this as a complete production Saga.

## Event Bus

The Employee Service publishes `EmployeeCreated` to `eventbus.Bus`. Subscribers run asynchronously in goroutines. The sample Salary consumer logs the event and initializes a salary row only when a positive salary amount is present. Because this is in-memory, events can be lost if the process stops; it is suitable for learning, not durable production messaging.

## Project layout

- `gateway/`: API Gateway and request ID middleware
- `employee/`, `department/`, `salary/`: each service's model, controller, service, repository
- `eventbus/`: in-process asynchronous event bus
- `saga/`: simple Saga coordinator endpoint
- `discovery/`: demo in-memory service registry
- `observability/`: basic structured-prefix logger
- `config/`: environment configuration
- `migrations/`: database schemas

## Build and test

```powershell
go fmt ./...
go vet ./...
go test ./...
go build ./...
```
