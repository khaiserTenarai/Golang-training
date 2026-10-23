# API Gateway

The API Gateway runs on port 8080.

- /api/auth/* -> Authentication Service
- /api/employees/* -> Employee Service
- /api/departments/* -> Department Service

The gateway checks JWTs for employee and department routes.
It propagates correlation IDs and forwards HTTP requests.