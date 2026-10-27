# Day 18 — Microservices Architecture: Design Notes

This document contains the tasks that are primarily analysis/design work, plus architecture notes. The matching task numbers are included so a reviewer can map the assignment to the project.

## Task 1 — Analyze the Employee monolith

A typical Employee monolith places employee management, department management, authentication, and notifications in one deployable application. These features may share code and a database, so a change in one area can affect others.

Potential issues as the application grows:
- Modules are tightly coupled.
- A single deployment is needed even for a small feature change.
- Scaling one busy feature may require scaling the whole application.
- A shared database can make ownership of data unclear.

This is a conceptual analysis for the assignment. Replace it with observations from your actual monolith if you have an existing implementation.

## Task 2 — Identify microservice boundaries

Proposed boundaries:
- **Employee Service:** employee profile and employment information.
- **Department Service:** department records and department membership rules.
- **Authentication Service:** login and identity-related operations.
- **API Gateway:** a single client-facing entry point; it routes requests but should not own business data.
- **Notification component:** consumes events and sends notifications.

Each business service owns its business rules and data. Avoid allowing one service to read or update another service's tables directly.

## Task 3 — Define bounded contexts

- **Employee context:** Employee ID, name, and Department ID.
- **Department context:** Department ID and department name.
- **Authentication context:** username and credentials/identity data.
- **Notification context:** event type and notification delivery status.

IDs can be shared between contexts, but each context should define its own model. Do not reuse one giant model struct across all services.

## Task 7 — Database-per-service design

For a later PostgreSQL version, use separate databases such as:
- `employee_db` owned by Employee Service
- `department_db` owned by Department Service
- `auth_db` owned by Authentication Service

A service should access its own database through its repository. Another service should use an API or event rather than querying its database directly.

The current starter uses in-memory repositories to keep local learning simple. Therefore, this task is currently documented/design-only, not implemented with real databases.

## Architecture and communication

- **Task 8 — Synchronous REST:** Employee Service can call Department Service through `employee/client/department_client.go`. REST is appropriate when the caller needs an immediate response.
- **Task 9 — Asynchronous communication:** `shared/messaging/event_bus.go` uses a Go channel to demonstrate event publishing and consuming. Because it is in-process, it is not a real broker connecting separate services.
- **Task 10 — API Gateway:** `gateway/handler/gateway_handler.go` forwards client requests to downstream services.
- **Task 11 — Service discovery:** `shared/discovery/consul.go` shows optional Consul registration. It requires a local Consul agent and is not enabled by default.
- **Tasks 12 and 13 — Distributed transactions and Saga:** `shared/saga/employee_saga.go` shows orchestration and a compensating action. A real Saga needs durable state, idempotency, timeouts, and retry handling.
- **Task 14 — Observability:** `shared/observability/logger.go` adds structured logs. Metrics and distributed tracing are future extensions.

## Task 12/13 — Example Saga sequence

1. Create an employee.
2. Ask Department Service to assign a department.
3. If assignment fails, run a compensating action to mark/delete the created employee.
4. Record the outcome so it can be inspected and retried.

A Saga is not a single ACID transaction spanning databases. Each service commits its local operation, and the workflow coordinates compensating actions when later steps fail.
