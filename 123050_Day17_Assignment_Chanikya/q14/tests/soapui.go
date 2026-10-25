package main

/*

Admin:
----------------

# 13. Test admin authentication

First, log in as admin.

Method: `POST`

URL: `http://localhost:10005/auth/login`

Body:

```
{
  "username": "admin",
  "password": "Admin@12345"
}
```

Copy the token from the login response and pass it in the `Authorization` header for all protected GraphQL requests.


Base endpoint: `http://localhost:10005/graphql`

For every request, use `POST` and these headers:

```
Content-Type: application/json
Authorization: Bearer YOUR_ADMIN_JWT_TOKEN
```

Replace `YOUR_ADMIN_JWT_TOKEN` with the token received from `POST http://localhost:10005/auth/login`.

1. Get all employees
{
  "query": "{ employees { data { id name email department { name } salary } page pageSize total totalPages } }"
}


2. Get employee by ID
{
  "query": "{ employee(id: \"1\") { id name email department { name } salary } }"
}


3. Create employee
{
  "query": "mutation { createEmployee(input: { name: \"Amit Kumar\", email: \"amit@example.com\", department: \"IT\", salary: 50000 }) { id name email department { name } salary } }"
}


4. Update employee
{
  "query": "mutation { updateEmployee(id: \"1\", input: { name: \"Amit Updated\", email: \"amit.updated@example.com\", department: \"Engineering\", salary: 60000 }) { id name email department { name } salary } }"
}


5. Delete employee
{
  "query": "mutation { deleteEmployee(id: \"1\") }"
}


Warning: Use a disposable employee record for deletion.
6. Filter employees by name
{
  "query": "{ employees(filter: { name: \"Amit\" }) { data { id name email department { name } salary } page pageSize total totalPages } }"
}


7. Filter employees by department
{
  "query": "{ employees(filter: { department: \"IT\" }) { data { id name email department { name } salary } page pageSize total totalPages } }"
}


8. Sort employees by name — ascending
{
  "query": "{ employees(sort: { field: \"name\", direction: \"asc\" }) { data { id name email department { name } salary } page pageSize total totalPages } }"
}


9. Sort employees by salary — highest first
{
  "query": "{ employees(sort: { field: \"salary\", direction: \"desc\" }) { data { id name email department { name } salary } page pageSize total totalPages } }"
}


10. Pagination — first five employees
The correct argument structure is pagination: { ... }. The following assumes PaginationInput uses page and pageSize; confirm these field names using the introspection query below if necessary.
{
  "query": "{ employees(pagination: { page: 1, pageSize: 5 }) { data { id name email department { name } salary } page pageSize total totalPages } }"
}


11. Pagination — next five employees
{
  "query": "{ employees(pagination: { page: 2, pageSize: 5 }) { data { id name email department { name } salary } page pageSize total totalPages } }"
}


12. Filter + sort + pagination together
{
  "query": "{ employees(filter: { department: \"IT\" }, sort: { field: \"salary\", direction: \"desc\" }, pagination: { page: 1, pageSize: 5 }) { data { id name email department { name } salary } page pageSize total totalPages } }"
}


13. Verify pagination input fields
To confirm the exact names and types supported by PaginationInput, run:
{
  "query": "{ __type(name: \"PaginationInput\") { name inputFields { name type { kind name ofType { kind name } } } } }"
}


User :
--------
http://localhost:10005/auth/login
{
"username": "user",
"password": "User@12345"
}
*/
