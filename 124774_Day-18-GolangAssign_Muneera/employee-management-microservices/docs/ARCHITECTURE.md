# Employee Management Microservices Architecture

## Services

- Employee Service: port 8081
- Department Service: port 8082
- Authentication Service: port 8083
- API Gateway: port 8080

## Technology

- Go and Gin
- GORM
- PostgreSQL
- JWT and bcrypt
- REST APIs
- PostgreSQL outbox and Go background worker

## Constraints

- No Docker
- No Kafka
- Each service owns its own database.
- Services communicate through APIs, not shared database tables.