package model

// Department belongs to the Department bounded context. Tasks 3 and 5.
type Department struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
