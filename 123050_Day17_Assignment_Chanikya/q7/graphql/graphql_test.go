package graphql

import (
	"testing"

	"employee-management/model"
	"employee-management/service"

	gql "github.com/graphql-go/graphql"
)

type fakeEmployeeService struct{}

func (fakeEmployeeService) Create(*model.Employee) error { return nil }

func (fakeEmployeeService) GetAll(
	_ model.EmployeeFilter,
	_ model.EmployeeSort,
	_, _ int,
) (*model.EmployeeListResponse, error) {
	return &model.EmployeeListResponse{
		Data: []model.Employee{
			{ID: 1, Name: "Amit", Email: "amit@example.com", Department: "IT", Salary: 50000},
		},
		Page: 1, PageSize: 10, Total: 1, TotalPages: 1,
	}, nil
}

func (fakeEmployeeService) GetByID(id uint) (*model.Employee, error) {
	return &model.Employee{
		ID: id, Name: "Amit", Email: "amit@example.com",
		Department: "IT", Salary: 50000,
	}, nil
}

func (fakeEmployeeService) Update(id uint, e *model.Employee) (*model.Employee, error) {
	e.ID = id
	return e, nil
}

func (fakeEmployeeService) Delete(uint) error { return nil }

var _ service.EmployeeService = fakeEmployeeService{}

func TestSchemaSupportsNestedDepartmentAndPagination(t *testing.T) {
	server, err := NewServer(fakeEmployeeService{})
	if err != nil {
		t.Fatalf("NewServer failed: %v", err)
	}

	result := gql.Do(gql.Params{
		Schema: *server.Schema(),
		RequestString: `
			query {
				employees(
					filter: { department: "IT" }
					sort: { field: "salary", direction: "desc" }
					pagination: { page: 1, pageSize: 10 }
				) {
					data {
						id
						name
						department { name }
						salary
					}
					page
					total
					totalPages
				}
			}
		`,
	})

	if len(result.Errors) > 0 {
		t.Fatalf("GraphQL execution errors: %+v", result.Errors)
	}
	if result.Data == nil {
		t.Fatal("expected GraphQL response data")
	}
}
