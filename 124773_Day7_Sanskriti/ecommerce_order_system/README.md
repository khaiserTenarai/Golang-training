# E-Commerce Management

Simple Go + PostgreSQL layered project.

## Structure

config/config.go
database/database.go
model/model.go
utility/utility.go
repository/repository.go
service/service.go
controller/controller.go
view/view.go
main.go

Each layer has only ONE Go file.

## Setup

1. Create database:
   CREATE DATABASE ecommerce_db;

2. Run schema.sql in PostgreSQL.

3. Update .env with your PostgreSQL password.

4. Run:
   go mod tidy
   go run .

## Transaction

CreateOrder uses one PostgreSQL transaction:

BEGIN
-> check stock
-> INSERT order
-> INSERT order items
-> reduce stock
-> COMMIT

If stock is insufficient, the function returns an error before commit and the deferred Rollback cancels the whole operation.
