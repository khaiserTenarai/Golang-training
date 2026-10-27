# Employee Microservices (Beginner-Friendly Go Project)


## Requirements

- Go 1.22 or later
- Internet access the first time, so Go can download Gin dependencies
- No database or deployment setup is required; this starter uses in-memory demo data

## Run the project

Open four separate terminals at the project root. Run each command in its own terminal and keep it running:

Terminal 1:
```bash
go run ./employee
```

Terminal 2:
```bash
go run ./department
```

Terminal 3:
```bash
go run ./auth
```

Terminal 4:
```bash
go run ./gateway
```

First-time setup (optional): run `go mod tidy` once from the project root.

## API endpoints

| Method | URL | Purpose |
|---|---|---|
| GET | `http://localhost:8080/health` | Gateway health |
| GET | `http://localhost:8080/employees` | List employees through gateway |
| POST | `http://localhost:8080/employees` | Create employee through gateway |
| GET | `http://localhost:8080/departments` | List demo departments |
| POST | `http://localhost:8080/login` | Test demo login |
| GET | `http://localhost:8081/health` | Employee Service health |
| GET | `http://localhost:8082/health` | Department Service health |
| GET | `http://localhost:8083/health` | Auth Service health |

