package resolver

import (
	"employee-management/model"
	"fmt"
	"strconv"
	"strings"
)


func parseID(v interface{}) (uint, error) {
	id, err := strconv.ParseUint(fmt.Sprint(v), 10, 32)
	if err != nil || id == 0 {
		return 0, newError("BAD_USER_INPUT", "invalid id")
	}
	return uint(id), nil
}


func str(v interface{}) string {
	s, _ := v.(string)
	return strings.TrimSpace(s)
}

func employeeFromInput(m map[string]interface{}) (model.Employee, error) {
	deptID, err := parseID(m["departmentId"])
	if err != nil {
		return model.Employee{}, err
	}
	salary, _ := m["salary"].(float64)
	return model.Employee{
		Name:         str(m["name"]),
		Email:        str(m["email"]),
		Salary:       salary,
		DepartmentID: deptID,
	}, nil
}


func filterFromArgs(args map[string]interface{}) (model.EmployeeFilter, error) {
	f := model.EmployeeFilter{}
	if v, ok := args["page"].(int); ok {
		f.Page = v
	}
	if v, ok := args["pageSize"].(int); ok {
		f.PageSize = v
	}
	if v, ok := args["search"].(string); ok {
		f.Search = v
	}
	if v, ok := args["sortBy"].(string); ok {
		f.SortBy = v
	}
	if v, ok := args["sortOrder"].(string); ok {
		f.SortOrder = v
	}
	if v, ok := args["minSalary"].(float64); ok {
		f.MinSalary = &v
	}
	if v, ok := args["maxSalary"].(float64); ok {
		f.MaxSalary = &v
	}
	if v, ok := args["departmentId"]; ok && v != nil {
		id, err := parseID(v)
		if err != nil {
			return f, err
		}
		f.DepartmentID = id
	}
	return f, nil
}
