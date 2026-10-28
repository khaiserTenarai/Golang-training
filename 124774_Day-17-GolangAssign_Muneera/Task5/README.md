
## Question 5: Testing Nested GraphQL Queries

GraphQL endpoint: `http://localhost:8081/graphql`

### 1. Get all employees with nested department details

```graphql
query {
  employees {
    id
    name
    email
    departmentId
    department {
      id
      name
    }
    salary
  }
}
```

### 2. Get one employee with nested department details

Replace `1` with an existing employee ID.

```graphql
query {
  employee(id: 1) {
    id
    name
    email
    departmentId
    department {
      id
      name
    }
    salary
  }
}
```

### 3. Create an employee with a department ID

Use an existing department ID.

```graphql
mutation {
  createEmployee(input: {
    name: "Muneera"
    email: "muneera@example.com"
    departmentId: 1
    salary: 30000
  }) {
    id
    name
    email
    departmentId
    department {
      id
      name
    }
    salary
  }
}
```

### 4. Update an employee's department

Replace `1` with an existing employee ID and `2` with an existing department ID.

```graphql
mutation {
  updateEmployee(id: 1, input: {
    name: "Muneera"
    email: "muneera@example.com"
    departmentId: 2
    salary: 35000
  }) {
    id
    name
    departmentId
    department {
      id
      name
    }
  }
}
```

### 5. Delete an employee

Replace `1` with an existing employee ID.

```graphql
mutation {
  deleteEmployee(id: 1)
}
```

### Expected Results

- Employee queries return nested department ID and name.
- Creating or updating an employee uses a valid department ID.
- Invalid employee or department IDs should return appropriate errors.
- The repository uses GORM `Preload("Department")` in `GetAll()` and `GetByID()` to load related department records efficiently.
