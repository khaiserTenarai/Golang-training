# Employee API (Go + PostgreSQL + JWT)

## Requirements
- Go 1.22+
- PostgreSQL
- Postman (optional)

## Setup
1. Create a PostgreSQL database named `emp_db_soap`.
2. Run `schema.sql` against that database.
3. Copy `.env.example` to `.env` and update `DATABASE_URL`.
4. Set `JWT_SECRET` to a long random secret.
5. Run:

```powershell
go mod tidy
go run .
```

The server listens on `http://localhost:8080`.

## Demo login
`POST /login` with JSON:
- Admin: `{"username":"admin","password":"admin123"}`
- Normal user: `{"username":"user","password":"user123"}`

Use the returned token as a Bearer token for `/employees` endpoints.
Demo credentials are hardcoded for learning only. Do not use this authentication setup in production.

## Endpoints
- `POST /login` (public)
- `GET /employees` (admin or user)
- `GET /employees/{id}` (admin or user)
- `POST /employees` (admin only)
- `PUT /employees/{id}` (admin only)
- `DELETE /employees/{id}` (admin only)

List query parameters:
- `page` (default 1)
- `limit` (default 10, maximum 100)
- `department` (substring filter)
- `sort` (`id`, `name`, `salary`, or `department`)
- `order` (`asc` or `desc`)

Example:
`GET /employees?department=IT&page=1&limit=5&sort=salary&order=desc`
## API Endpoints

Base URL: `http://localhost:10006`

| Operation          | Method | URL                                  | Access      |
| ------------------ | ------ | ------------------------------------ | ----------- |
| Login              | POST   | `http://localhost:10006/login`       | Public      |
| Get all employees  | GET    | `http://localhost:10006/employees`   | Admin, User |
| Get employee by ID | GET    | `http://localhost:10006/employees/1` | Admin, User |
| Create employee    | POST   | `http://localhost:10006/employees`   | Admin only  |
| Update employee    | PUT    | `http://localhost:10006/employees/1` | Admin only  |
| Delete employee    | DELETE | `http://localhost:10006/employees/1` | Admin only  |

## Login Credentials

| Role        | Username | Password   |
| ----------- | -------- | ---------- |
| Admin       | `admin`  | `admin123` |
| Normal User | `user`   | `user123`  |

> **Note:** Login works only when the corresponding bcrypt password hashes are configured in the `.env` file.

## Authentication

1. Send a `POST` request to `/login` with the username and password in a JSON request body.
2. Copy the JWT token returned in the response.
3. For protected endpoints, set the `Authorization` header to `Bearer <your_token>`.
4. Admin users can create, update, and delete employees. Normal users have read-only access.

## Example Login Request

**Method:** `POST`

**URL:** `http://localhost:10006/login`

**Request Body:**

```json
{
  "username": "admin",
  "password": "admin123"
}
```

**Example Response:**

```json
{
  "token": "<your_jwt_token>",
  "role": "admin"
}
```

The token shown above is a placeholder. The API generates a real token when login succeeds.
