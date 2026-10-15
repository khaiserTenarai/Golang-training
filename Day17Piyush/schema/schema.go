// Tasks 1, 3, 4, 5: Schema assembly with service injection and nested queries

package schema

import (
	"fmt"

	"graphql_employee/models"
	"graphql_employee/service"

	"github.com/graphql-go/graphql"
)

// Task 3: Services injected into schema builder
func BuildSchema(empSvc *service.EmployeeService, deptSvc *service.DepartmentService, authSvc *service.AuthService) (graphql.Schema, error) {

	// Task 4: Add nested Department field to Employee (resolved via dept service)
	// Task 5: Optimized — resolves per-employee using service method
	EmployeeType.AddFieldConfig("department", &graphql.Field{
		Type: DepartmentType,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			emp, ok := p.Source.(models.Employee)
			if !ok {
				return nil, nil
			}
			dept, err := deptSvc.GetByID(emp.DepartmentID)
			if err != nil {
				return nil, nil
			}
			return dept, nil
		},
	})

	// ---- QUERIES ----
	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{

			// Single employee
			"employee": &graphql.Field{
				Type: EmployeeType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// Task 6: Check auth from context
					if err := requireAuth(p); err != nil {
						return nil, err
					}
					id := p.Args["id"].(int)
					return empSvc.GetByID(id)
				},
			},

			// Employee list with pagination, filtering, sorting (Tasks 9, 10, 11)
			"employees": &graphql.Field{
				Type: EmployeeListType,
				Args: graphql.FieldConfigArgument{
					"page":       &graphql.ArgumentConfig{Type: graphql.Int, DefaultValue: 1},
					"page_size":  &graphql.ArgumentConfig{Type: graphql.Int, DefaultValue: 10},
					"name":       &graphql.ArgumentConfig{Type: graphql.String},
					"min_salary": &graphql.ArgumentConfig{Type: graphql.Float},
					"max_salary": &graphql.ArgumentConfig{Type: graphql.Float},
					"sort_by":    &graphql.ArgumentConfig{Type: graphql.String, DefaultValue: "id"},
					"sort_order": &graphql.ArgumentConfig{Type: graphql.String, DefaultValue: "asc"},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if err := requireAuth(p); err != nil {
						return nil, err
					}
					page := p.Args["page"].(int)
					pageSize := p.Args["page_size"].(int)
					filter := models.EmployeeFilter{
						SortBy:    p.Args["sort_by"].(string),
						SortOrder: p.Args["sort_order"].(string),
					}
					if v, ok := p.Args["name"].(string); ok {
						filter.Name = v
					}
					if v, ok := p.Args["min_salary"].(float64); ok {
						filter.MinSalary = v
					}
					if v, ok := p.Args["max_salary"].(float64); ok {
						filter.MaxSalary = v
					}

					employees, total := empSvc.GetAll(filter, models.PaginationInput{Page: page, PageSize: pageSize})
					return map[string]interface{}{
						"employees":   employees,
						"total":       total,
						"page":        page,
						"page_size":   pageSize,
						"total_pages": service.TotalPages(total, pageSize),
					}, nil
				},
			},

			// Departments
			"departments": &graphql.Field{
				Type: graphql.NewList(DepartmentType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return deptSvc.GetAll(), nil
				},
			},
		},
	})

	// ---- MUTATIONS ----
	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{

			// Task 6: Login mutation
			"login": &graphql.Field{
				Type: AuthResponseType,
				Args: graphql.FieldConfigArgument{
					"username": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
					"password": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					token, user, err := authSvc.Login(
						p.Args["username"].(string),
						p.Args["password"].(string),
					)
					if err != nil {
						return nil, err
					}
					return map[string]interface{}{
						"token": token, "user_id": user.ID,
						"username": user.Username, "role": user.Role,
					}, nil
				},
			},

			// Register
			"register": &graphql.Field{
				Type: AuthResponseType,
				Args: graphql.FieldConfigArgument{
					"username": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
					"password": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
					"role":     &graphql.ArgumentConfig{Type: graphql.String, DefaultValue: "user"},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					token, user, err := authSvc.Register(
						p.Args["username"].(string),
						p.Args["password"].(string),
						p.Args["role"].(string),
					)
					if err != nil {
						return nil, err
					}
					return map[string]interface{}{
						"token": token, "user_id": user.ID,
						"username": user.Username, "role": user.Role,
					}, nil
				},
			},

			// Create employee (Task 7: admin/manager only, Task 8: validation)
			"createEmployee": &graphql.Field{
				Type: EmployeeType,
				Args: graphql.FieldConfigArgument{
					"input": &graphql.ArgumentConfig{Type: graphql.NewNonNull(EmployeeInputType)},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if err := requireRole(p, "admin", "manager"); err != nil {
						return nil, err
					}
					input := p.Args["input"].(map[string]interface{})
					emp := models.Employee{
						Name:         input["name"].(string),
						Email:        input["email"].(string),
						Age:          input["age"].(int),
						DepartmentID: input["department_id"].(int),
						Salary:       input["salary"].(float64),
					}
					if r, ok := input["role"].(string); ok {
						emp.Role = r
					}
					return empSvc.Create(emp)
				},
			},

			// Update employee
			"updateEmployee": &graphql.Field{
				Type: EmployeeType,
				Args: graphql.FieldConfigArgument{
					"id":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
					"input": &graphql.ArgumentConfig{Type: graphql.NewNonNull(EmployeeInputType)},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if err := requireRole(p, "admin", "manager"); err != nil {
						return nil, err
					}
					id := p.Args["id"].(int)
					input := p.Args["input"].(map[string]interface{})
					emp := models.Employee{
						Name:         input["name"].(string),
						Email:        input["email"].(string),
						Age:          input["age"].(int),
						DepartmentID: input["department_id"].(int),
						Salary:       input["salary"].(float64),
					}
					if r, ok := input["role"].(string); ok {
						emp.Role = r
					}
					return empSvc.Update(id, emp)
				},
			},

			// Delete employee
			"deleteEmployee": &graphql.Field{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if err := requireRole(p, "admin"); err != nil {
						return nil, err
					}
					return empSvc.Delete(p.Args["id"].(int)) == nil, nil
				},
			},
		},
	})

	return graphql.NewSchema(graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	})
}

// Task 6: Auth helper — checks JWT info from context
func requireAuth(p graphql.ResolveParams) error {
	if p.Context.Value("role") == nil {
		return fmt.Errorf("authentication required")
	}
	return nil
}

// Task 7: Authorization helper
func requireRole(p graphql.ResolveParams, roles ...string) error {
	if err := requireAuth(p); err != nil {
		return err
	}
	userRole, _ := p.Context.Value("role").(string)
	for _, r := range roles {
		if r == userRole {
			return nil
		}
	}
	return fmt.Errorf("access denied: requires role %v", roles)
}
