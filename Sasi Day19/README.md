# Enterprise Go Microservices Architecture with API Gateway

This project implements an **API Gateway Architecture** in Go for routing, correlation ID propagation, and central entry-point management across Employee and Department microservices.

## Port Architecture:
- **API Gateway:** `http://localhost:8080` (Single Entry Point for Clients)
- **Employee Service:** `http://localhost:8081` (Internal Microservice)
- **Department Service:** `http://localhost:8082` (Internal Microservice)

## How Request Flow Works:

1. **Accessing Department via Gateway (`Port 8080`):**
   ```bash
   GET http://localhost:8080/api/v1/departments/1
   ```
   * *Gateway receives call on port 8080.*
   * *Gateway injects/passes `X-Correlation-ID` header.*
   * *Gateway routes call to Department Service on port 8082.*

2. **Accessing Employee via Gateway (`Port 8080`):**
   ```bash
   GET http://localhost:8080/api/v1/employees/1
   ```
   * *Gateway receives call on port 8080.*
   * *Gateway routes call to Employee Service on port 8081.*
   * *Employee Service calls Department Service (port 8082) using `http.Client` with Timeout, Retry, Backoff, and Circuit Breaker.*

## How to Run:
1. Start Department Service:
   ```bash
   cd department-service && go run main.go
   ```
2. Start Employee Service:
   ```bash
   cd employee-service && go run main.go
   ```
3. Start API Gateway:
   ```bash
   cd api-gateway && go run main.go
   ```
