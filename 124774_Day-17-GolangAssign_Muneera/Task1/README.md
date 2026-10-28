# Question 1: Organize GraphQL Schema into Modules

## Objective
To organize the GraphQL schema into separate files for types, queries, mutations, schema initialization, and HTTP handling.

## Technologies Used
- Go (Golang)
- GraphQL
- Gin
- GORM
- PostgreSQL

## Project Structure


01-Create-GraphQL-Schema/
├── config/
├── controller/
├── graphql/
│   ├── types.go
│   ├── query.go
│   ├── mutation.go
│   ├── schema.go
│   └── handler.go
├── middleware/
├── model/
├── repository/
├── service/
├── utility/
├── tests/
├── .env
├── go.mod
├── main.go
└── README.md
```

## How to Run

### 1. Install dependencies

```bash
go mod tidy
```

### 2. Configure PostgreSQL

Create the database `employee3_db` and configure the database credentials in `.env`.

### 3. Start the application

```bash
go run .
```

## GraphQL Request URL
http://localhost:8081/graphql




## GraphQL Queries

### Get All Employees

```graphql
{
  employees {
    id
    name
    email
    department
    salary
  }
}
```

### Get Employee by ID

```graphql
{
  employee(id: "101") {
    id
    name
    email
    department
    salary
  }
}
```

## GraphQL Mutations

### Create Employee

```graphql
mutation {
  createEmployee(input: {
    name: "Muneera"
    email: "muneera@example.com"
    department: "IT"
    salary: 55000
  }) {
    id
    name
    email
    department
    salary
  }
}
```

### Update Employee

```graphql
mutation {
  updateEmployee(
    id: "101"
    input: {
      name: "Rajesh Kumar"
      email: "rajesh@example.com"
      department: "IT"
      salary: 60000
    }
  ) {
    id
    name
    email
    department
    salary
  }
}
```

### Delete Employee

```graphql
mutation {
  deleteEmployee(id: "101")
}
```

