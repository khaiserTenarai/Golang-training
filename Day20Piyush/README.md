# Day 20 — AWS Fundamentals (Piyush)

A Go Employee REST API designed for AWS deployment. Runs locally with in-memory store or connects to PostgreSQL (RDS) when `DATABASE_URL` is set.

## Project Structure

```
Day20Piyush/
├── .gitignore
├── README.md                    # This file
├── decisions.md                 # Architecture decisions
├── app/                         # Go Employee API
│   ├── go.mod
│   ├── main.go                  # Entry point (Tasks 5, 6, 12)
│   ├── config/config.go         # Env-based configuration
│   ├── models/employee.go       # Data models
│   └── handler/
│       ├── employee.go          # CRUD handlers (in-memory or PostgreSQL)
│       └── s3_handler.go        # S3 upload handler (Tasks 9, 10)
└── aws/                         # AWS deployment configs
    ├── architecture.txt         # Tasks 1, 2, 3, 14: AWS concepts + diagrams
    ├── iam-policy.json          # Task 4: IAM role policy
    └── setup-all.sh             # Tasks 4-14: Full deployment script
```

## Quick Start (Local — No AWS Needed)

```bash
cd app
go mod tidy
go run main.go
```

Test:
```bash
curl http://localhost:8080/health
curl http://localhost:8080/api/employees
curl -X POST http://localhost:8080/api/employees \
  -H "Content-Type: application/json" \
  -d '{"name":"Eve","email":"eve@test.com","age":26,"department":"HR","salary":52000}'
```

## Task Mapping

| # | Task | Where |
|---|------|-------|
| 1 | AWS Regions | aws/architecture.txt |
| 2 | Availability Zones | aws/architecture.txt |
| 3 | Architecture diagram | aws/architecture.txt (ASCII diagram) |
| 4 | IAM role | aws/iam-policy.json + setup-all.sh |
| 5 | Deploy on EC2 | aws/setup-all.sh (EC2 launch) |
| 6 | Configure EC2 service | aws/setup-all.sh (systemd unit) |
| 7 | Design VPC | aws/setup-all.sh (VPC + 4 subnets) |
| 8 | Security Groups | aws/setup-all.sh (ALB/EC2/RDS SGs) |
| 9 | Create S3 bucket | aws/setup-all.sh + app/handler/s3_handler.go |
| 10 | Upload report to S3 | POST /api/s3/upload-report |
| 11 | Create RDS PostgreSQL | aws/setup-all.sh (db.t3.micro) |
| 12 | Connect to RDS | app/main.go (DATABASE_URL env var) |
| 13 | CloudWatch logging | aws/setup-all.sh (log groups) |
| 14 | Load-balanced architecture | aws/architecture.txt (ALB + multi-AZ) |
| 15 | Mini Project | Everything combined |

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | /health | Health check (shows DB status) |
| GET | /api/employees | List all employees |
| GET | /api/employees/{id} | Get by ID |
| POST | /api/employees | Create employee |
| PUT | /api/employees/{id} | Update employee |
| DELETE | /api/employees/{id} | Delete employee |
| GET | /api/employees/report | Download CSV report |
| POST | /api/s3/upload-report | Upload report to S3 |
