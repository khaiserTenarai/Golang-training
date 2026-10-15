// Day 16 — GraphQL Fundamentals (All 15 Tasks)
// Task 1:  GraphQL schema (assembled in resolver/resolver.go)
// Task 12: GraphQL variables (supported via variables field in request)
// Task 15: Mini Project — Complete Employee GraphQL API
//
// Run:  go mod tidy && go run main.go
// Playground: http://localhost:8080/graphql (open in browser)

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"graphql_fundamentals/resolver"
	"graphql_fundamentals/store"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

func main() {
	dataStore := store.NewStore()
	schema, err := resolver.BuildSchema(dataStore)
	if err != nil {
		panic("Failed to build schema: " + err.Error())
	}

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "healthy", "service": "graphql-fundamentals"})
	})

	// GraphQL endpoint
	r.POST("/graphql", func(c *gin.Context) {
		var req struct {
			Query     string                 `json:"query"`
			Variables map[string]interface{} `json:"variables"` // Task 12: Variables support
			Operation string                 `json:"operationName"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request: " + err.Error()})
			return
		}

		result := graphql.Do(graphql.Params{
			Schema:         schema,
			RequestString:  req.Query,
			VariableValues: req.Variables, // Task 12: Pass variables
			OperationName:  req.Operation,
		})

		// Pretty print for readability
		c.Header("Content-Type", "application/json")
		enc := json.NewEncoder(c.Writer)
		enc.SetIndent("", "  ")
		enc.Encode(result)
	})

	// GraphQL Playground (browser UI)
	r.GET("/graphql", func(c *gin.Context) {
		c.Header("Content-Type", "text/html")
		c.String(http.StatusOK, `<!DOCTYPE html><html><head><title>GraphQL Playground</title>
<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/graphql-playground-react/build/static/css/index.css"/>
<script src="https://cdn.jsdelivr.net/npm/graphql-playground-react/build/static/js/middleware.js"></script>
</head><body><div id="root"></div><script>
window.addEventListener('load',function(){GraphQLPlayground.init(document.getElementById('root'),{endpoint:'/graphql'})})
</script></body></html>`)
	})

	fmt.Println("╔═══════════════════════════════════════════════════════╗")
	fmt.Println("║   Day 16: GraphQL Fundamentals — Employee API        ║")
	fmt.Println("╠═══════════════════════════════════════════════════════╣")
	fmt.Println("║  Playground: http://localhost:8080/graphql (browser)  ║")
	fmt.Println("║  Endpoint:   POST http://localhost:8080/graphql       ║")
	fmt.Println("╠═══════════════════════════════════════════════════════╣")
	fmt.Println("║  Queries:                                            ║")
	fmt.Println("║    employee(id)              → single employee       ║")
	fmt.Println("║    employees(limit,offset)   → list with pagination  ║")
	fmt.Println("║    department(id)            → single department     ║")
	fmt.Println("║    departments               → all departments       ║")
	fmt.Println("║  Mutations:                                          ║")
	fmt.Println("║    createEmployee(input)     → create new            ║")
	fmt.Println("║    updateEmployee(id,input)  → update existing       ║")
	fmt.Println("║    deleteEmployee(id)        → delete                ║")
	fmt.Println("╚═══════════════════════════════════════════════════════╝")

	r.Run(":8080")
}

/*
══════════════════════════════════════════════
  GRAPHQL QUERY EXAMPLES (paste in Playground)
══════════════════════════════════════════════

# Task 4: Single employee query
{ employee(id: 1) { id name email salary is_active } }

# Task 4 + nested department
{ employee(id: 1) { name salary department { name location } } }

# Task 5: List query with Task 11 arguments
{ employees(limit: 3, offset: 0) { employees { id name salary is_active } count } }

# Task 5: Filter by active status
{ employees(is_active: true) { employees { name is_active } count } }

# Task 3: Department queries
{ departments { id name location budget } }
{ department(id: 1) { name budget } }

# Task 6: Create mutation (Task 10: input type)
mutation {
  createEmployee(input: {
    name: "Eve"
    email: "eve@test.com"
    age: 26
    department_id: 2
    salary: 52000
    is_active: true
    join_date: "2024-06-01"
  }) {
    id name email department { name }
  }
}

# Task 7: Update mutation
mutation {
  updateEmployee(id: 1, input: {
    name: "Piyush Kumar"
    email: "piyush@updated.com"
    age: 26
    department_id: 1
    salary: 75000
    is_active: true
  }) {
    id name email salary
  }
}

# Task 8: Delete mutation
mutation { deleteEmployee(id: 5) { success message } }

# Task 12: Using GraphQL variables
# In the "Query Variables" panel of Playground, add:
#   { "empId": 1 }
# Then run:
query GetEmployee($empId: Int!) {
  employee(id: $empId) {
    id name email salary
    department { name location }
  }
}

# Task 12: Variables with mutation
# Variables: { "input": {"name":"Frank","email":"frank@test.com","age":30,"department_id":1,"salary":65000} }
mutation CreateEmp($input: CreateEmployeeInput!) {
  createEmployee(input: $input) {
    id name email
  }
}

# Task 14: Error handling — try non-existent ID
{ employee(id: 999) { name } }

# Task 14: Validation error — bad age
mutation {
  createEmployee(input: {
    name: "Bad"
    email: "bad@test.com"
    age: 10
    department_id: 1
    salary: -100
  }) { id }
}
*/
