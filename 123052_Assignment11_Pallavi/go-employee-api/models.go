package main

type Employee struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

func (e *Employee) Validate() string {
	if e.Name == "" {
		return "Name field is required"
	}
	if e.Role == "" {
		return "Role field is required"
	}
	return ""
}