# Day 15 — Advanced Gin

## Setup (once)

I could not download Gin here, so there is no go.sum yet. Run this once in
this folder while you are online:

    go mod tidy

## How to run

Questions 1-13 are single files in this folder. Run one at a time:

    go run 01-custom-middleware.go

Every file has its own `package main`, `main()` and small helper
functions, so `go run .` or `go build ./...` on the whole folder will give
"redeclared" errors. That is expected: run one file at a time.

Each program asks you for its settings when it starts (port number,
secret key, usernames and passwords, allowed website, upload folder, and
so on). Nothing like that is written inside the code.

Question 14 and 15 are folders:

    cd 14-test-handlers-httptest
    go test -v
    go run .

    cd 15-secure-gin-api
    go test -v
    go run .

## Trying the questions

    # Q1-Q3, Q5, Q6 (use the port you typed)
    curl "http://localhost:8080/hello?name=Anita"
    curl "http://localhost:8080/divide?a=10&b=0"          # Q6, triggers the panic

    # Q4 (the password you typed goes in the header)
    curl -H "X-Password: yourpassword" http://localhost:8080/private/data

    # Q8-Q10, Q15: log in, then use the token
    curl -X POST http://localhost:8080/login -H "Content-Type: application/json" \
         -d '{"username":"yourname","password":"yourpassword"}'
    curl -H "Authorization: Bearer <token>" http://localhost:8080/api/profile

    # Q11
    curl -X POST http://localhost:8080/employees -H "Content-Type: application/json" \
         -d '{"name":"Anita","email":"anita@example.com","age":28,"salary":45000}'

    # Q12
    curl "http://localhost:8080/employees?page=2&limit=5"

    # Q13
    curl -F "file=@photo.png" http://localhost:8080/upload

## Question 15 routes

| Route | Who can use it |
|-------|----------------|
| POST /login | anyone |
| GET /api/profile | any logged in user |
| GET /api/employees?page=&limit= | any logged in user |
| POST /api/employees | admin only, validated |
| DELETE /api/employees/:id | admin only |

Middleware on every route: request ID, logger, recovery.

## Not yet run

`gofmt` passed on every file (no syntax errors). I could not run
`go test` or the servers, because the network here blocked the Go module
download. Please run `go mod tidy` and then `go test -v` in folders 14 and
15 and tell me if anything fails.
