## Question 9: Pagination

### Test 1: First Page

```graphql
query {
  employees(page: 1, limit: 2) {
    id
    name
    email
    department
    salary
  }
}
```

### Test 2: Second Page

```graphql
query {
  employees(page: 2, limit: 2) {
    id
    name
    email
    department
    salary
  }
}
```

## Question 10: Filtering

### Test 1: Filter by AI Department

```graphql
query {
  employees(department: "AI") {
    id
    name
    email
    department
    salary
  }
}
```

### Test 2: Filter by Engineering Department

```graphql
query {
  employees(department: "Engineering") {
    id
    name
    email
    department
    salary
  }
}
```

## Question 11: Sorting

### Test 1: Sort Salary in Ascending Order

```graphql
query {
  employees(sortBy: "salary", sortOrder: "asc") {
    id
    name
    salary
  }
}
```

### Test 2: Sort Salary in Descending Order

```graphql
query {
  employees(sortBy: "salary", sortOrder: "desc") {
    id
    name
    salary
  }
}
```

### Test 3: Sort Employees by Name

```graphql
query {
  employees(sortBy: "name", sortOrder: "asc") {
    id
    name
    salary
  }
}
```

## Combined Test: Pagination, Filtering and Sorting

```graphql
query {
  employees(
    page: 1
    limit: 2
    department: "AI"
    sortBy: "salary"
    sortOrder: "desc"
  ) {
    id
    name
    email
    department
    salary
  }
}
```
