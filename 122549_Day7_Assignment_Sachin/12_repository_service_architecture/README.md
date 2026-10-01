# 12_repository_service_architecture

Each project uses its own PostgreSQL database and follows Controller -> Service -> Repository -> PostgreSQL.

## Setup

1. Create the database named in ".env.example".
2. Run "database/init.sql" in that database.
3. Set the PostgreSQL environment variables from ".env.example" in your terminal.
4. Run "go mod tidy" when internet/module cache is available.
5. Run "go run .".

All application values are entered through the console.
