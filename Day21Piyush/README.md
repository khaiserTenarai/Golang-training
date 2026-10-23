# Day 21 — AWS + Go Integration (Piyush)

Go Employee API extended with AWS SDK for S3, Secrets Manager, CloudWatch, and RDS. Runs locally (in-memory + simulated AWS) or on AWS (real services).

## Quick Start (No AWS Needed)
```bash
cd app
go mod tidy
go run main.go
```
All AWS endpoints return simulated responses when credentials aren't configured.

## Project Structure
```
Day21Piyush/
├── app/
│   ├── main.go              # Task 15: Mini Project entry point
│   ├── config/config.go     # Task 10: Externalized AWS config
│   ├── models/employee.go   # Data models
│   ├── handler/employee.go  # CRUD handlers (in-memory or PostgreSQL)
│   └── awsutil/
│       ├── s3.go            # Tasks 3, 5: S3 SDK operations
│       ├── secrets.go       # Task 8: Secrets Manager
│       ├── cloudwatch.go    # Task 11: CloudWatch Logs
│       └── rds.go           # Task 6: RDS connectivity test
├── scripts/
│   └── aws_cli_commands.sh  # Tasks 1, 2, 4, 9, 11: CLI commands
├── docs/
│   ├── 07_iam_least_privilege.txt  # Task 7
│   ├── 12_high_availability.txt    # Task 12
│   ├── 13_security_checklist.txt   # Task 13
│   └── 14_cost_report.txt          # Task 14
├── decisions.md             # Architecture decisions
├── .gitignore               # Ignores flow.md, run.md
└── README.md
```

## Endpoints
| Method | Path | Task | Description |
|--------|------|------|-------------|
| GET | /health | — | Health check |
| GET | /api/employees | 15 | List employees |
| POST | /api/employees | 15 | Create employee |
| GET | /api/employees/{id} | 15 | Get by ID |
| DELETE | /api/employees/{id} | 15 | Delete |
| POST | /api/aws/s3/upload | 5 | Upload CSV to S3 |
| GET | /api/aws/s3/list | 5 | List S3 objects |
| GET | /api/aws/rds/test | 6 | Test RDS connectivity |
| GET | /api/aws/secrets | 8 | Retrieve from Secrets Manager |
| POST | /api/aws/cloudwatch/log | 11 | Write to CloudWatch |

## Task Mapping
| # | Task | Implementation |
|---|------|---------------|
| 1 | Configure CLI | scripts/aws_cli_commands.sh |
| 2 | Verify identity | scripts/aws_cli_commands.sh |
| 3 | AWS SDK for Go | app/awsutil/ (session setup) |
| 4 | S3 CLI ops | scripts/aws_cli_commands.sh |
| 5 | S3 from Go | app/awsutil/s3.go + endpoints |
| 6 | Test RDS | app/awsutil/rds.go |
| 7 | IAM review | docs/07_iam_least_privilege.txt |
| 8 | Secrets Manager | app/awsutil/secrets.go |
| 9 | EC2 deploy script | scripts/aws_cli_commands.sh |
| 10 | Externalize config | app/config/config.go |
| 11 | CloudWatch | app/awsutil/cloudwatch.go |
| 12 | HA architecture | docs/12_high_availability.txt |
| 13 | Security checklist | docs/13_security_checklist.txt |
| 14 | Cost report | docs/14_cost_report.txt |
| 15 | Mini Project | App uses S3 + Secrets Manager via Go SDK |
