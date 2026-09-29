package main

type Employee struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	JobTitle  string    `json:"job_title"`
	Salary    float64   `json:"salary"`
	IsActive  bool      `json:"is_active"`
}

func (e Employee) FullName() string {
	return e.FirstName + " " + e.LastName
}

func (e *Employee) GiveRaise(percent float64) {
	e.Salary += e.Salary * (percent / 100.0)
}

func (e *Employee) Deactivate() {
	e.IsActive = false
}