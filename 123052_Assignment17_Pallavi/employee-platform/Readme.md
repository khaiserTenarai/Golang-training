### Commands to Run and Test

1. **Initialize Dependencies:**
   ```bash
   go mod tidy
  

2. **Run Server:**
   ```bash
   go run main.go
   

3. **Run Automated Tests:**
   ```bash
   go test -v .
   

4. **Execute Nested GraphQL Query in GraphiQL UI (`http://localhost:8080/graphql`):**
   ```graphql
   query {
     employees(department: "Engineering", sortBy: "salary_desc", limit: 5, offset: 0) {
       id
       name
       salary
       department {
         id
         name
       }
     }
   }


5. **Execute Authorized Mutation:**
   Set the HTTP Header in GraphiQL or cURL:  
   `Authorization`: `Bearer admin-token`

   ```graphql
   mutation {
     createEmployee(
       name: "Pallavi"
       email: "pallavi@example.com"
       departmentId: 101
       salary: 90000
     ) {
       id
       name
       department {
         name
       }
     }
   }