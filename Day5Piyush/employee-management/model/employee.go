package model

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zip_code"`
}

type Department struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Manager string `json:"manager"`
}

type Employee struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	Email      string     `json:"email"`
	Age        int        `json:"age"`
	Salary     float64    `json:"salary"`
	Phone      string     `json:"phone"`
	Position   string     `json:"position"`
	Address    Address    `json:"address"`
	Department Department `json:"department"`
}

func (e Employee) Display() {
	println("ID:", e.ID)
	println("Name:", e.Name)
	println("Email:", e.Email)
	println("Age:", e.Age)
	println("Salary:", e.Salary)
	println("Phone:", e.Phone)
	println("Position:", e.Position)
	println("Department:", e.Department.Name)
}

func (e *Employee) IncreaseSalary(amount float64) {
	e.Salary += amount
}
