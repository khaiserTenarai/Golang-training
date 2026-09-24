package employee

type Employee struct {
	ID     int
	Name   string
	Salary float64
}

func (e *Employee) GiveRaise(percent float64) {
	e.Salary += e.Salary * (percent / 100)
}