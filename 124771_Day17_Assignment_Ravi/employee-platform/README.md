# Day 17 Mini Project — Employee Platform (REST + GraphQL)

A beginner-friendly Go layered application using Gin, `graphql-go/graphql`, GORM, PostgreSQL, JWT authentication, and role-based authorization.

## What is included

- One Go application with REST and GraphQL endpoints sharing the same service layer.
- PostgreSQL connection and GORM `AutoMigrate` for `users`, `departments`, and `employees`.
- Layered architecture: model -> repository -> service -> REST controller / GraphQL resolver.
- JWT login and bcrypt password hashing.
- Role-based authorization: public registration creates `employee`; admin-only create/update/delete operations.
- Nested Employee -> Department GraphQL field.
- Pagination, filtering, sorting for employee listing.
- Classic GraphiQL browser IDE at `/graphql`, enabled through `github.com/graphql-go/handler` (`GraphiQL: true`, `Playground: true`).
- Comments in code identify the assignment task areas.

## Requirements

- Go 1.23 or newer
- PostgreSQL 14+ (or compatible version)

## 1. Create database

Open psql or pgAdmin and run:

```sql
CREATE DATABASE employee_platform;
```

The application creates the tables automatically with GORM `AutoMigrate` after it connects. The database itself must exist first.


## 2. Install dependencies and run

From the project root:

```bash
go mod tidy
go run .
```

The server starts at `http://localhost:8080`. If PostgreSQL is not running, the application prints a database connection error.


The default registration endpoint intentionally does not let users choose their role. To create a local admin, stop the server and set these variables in `.env`:

```env
SEED_DEMO_ADMIN=true
DEMO_ADMIN_EMAIL=admin@example.com
DEMO_ADMIN_PASSWORD=Admin123!
DEMO_ADMIN_NAME=Demo Admin
```

Restart the app. Change the demo password for your own environment. Do not use demo credentials in production.

## 4. REST endpoints to test

Use Postman, curl, or another API client.

### Health check

```http
GET http://localhost:8080/health
```

