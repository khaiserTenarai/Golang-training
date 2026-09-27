package task3

type Employee struct{
	ID int 				`json:"id"`
	Name string			`json:"name"`
	Age int				`json:"age"`
	Salary float64		`json:"salary"`
	FirstName string	`json:"firstname"`
	LastName string		`json:"lastname"`
	Email string		`json:"email"`
	IsActive bool		`json:"isactive"`
}