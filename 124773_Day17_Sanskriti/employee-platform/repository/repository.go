package repository

import (
	"context"
	"fmt"
	"strings"

	"employee-platform/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct { DB *pgxpool.Pool }
func New(db *pgxpool.Pool) *Repository { return &Repository{DB: db} }

func (r *Repository) ListEmployees(ctx context.Context, page, limit int, name, department, sortBy, order string) ([]model.Employee, int, error) {
	if page < 1 { page = 1 }; if limit < 1 { limit = 10 }; if limit > 100 { limit = 100 }
	allowed := map[string]bool{"id":true,"name":true,"email":true,"age":true,"salary":true}
	if !allowed[sortBy] { sortBy = "id" }
	order = strings.ToUpper(order); if order != "DESC" { order = "ASC" }
	where := ` WHERE ($1 = '' OR e.name ILIKE '%' || $1 || '%') AND ($2 = '' OR d.name ILIKE '%' || $2 || '%')`
	var total int
	err := r.DB.QueryRow(ctx, `SELECT count(*) FROM employees e JOIN departments d ON d.id=e.department_id`+where, name, department).Scan(&total)
	if err != nil { return nil, 0, err }
	q := fmt.Sprintf(`SELECT e.id,e.name,e.email,e.age,e.salary,e.department_id,d.id,d.name
		FROM employees e JOIN departments d ON d.id=e.department_id %s ORDER BY e.%s %s LIMIT $3 OFFSET $4`, where, sortBy, order)
	rows, err := r.DB.Query(ctx, q, name, department, limit, (page-1)*limit)
	if err != nil { return nil, 0, err }; defer rows.Close()
	out := []model.Employee{}
	for rows.Next() {
		var e model.Employee; var d model.Department
		if err := rows.Scan(&e.ID,&e.Name,&e.Email,&e.Age,&e.Salary,&e.DepartmentID,&d.ID,&d.Name); err != nil { return nil,0,err }
		e.Department = &d; out = append(out,e)
	}
	return out,total,rows.Err()
}

func (r *Repository) GetEmployee(ctx context.Context, id int64) (model.Employee,error) {
	var e model.Employee; var d model.Department
	err := r.DB.QueryRow(ctx, `SELECT e.id,e.name,e.email,e.age,e.salary,e.department_id,d.id,d.name FROM employees e JOIN departments d ON d.id=e.department_id WHERE e.id=$1`, id).
		Scan(&e.ID,&e.Name,&e.Email,&e.Age,&e.Salary,&e.DepartmentID,&d.ID,&d.Name)
	if err == nil { e.Department = &d }; return e,err
}
func (r *Repository) CreateEmployee(ctx context.Context, in model.EmployeeInput) (model.Employee,error) {
	var id int64
	err := r.DB.QueryRow(ctx, `INSERT INTO employees(name,email,age,salary,department_id) VALUES($1,$2,$3,$4,$5) RETURNING id`, in.Name,in.Email,in.Age,in.Salary,in.DepartmentID).Scan(&id)
	if err != nil { return model.Employee{},err }; return r.GetEmployee(ctx,id)
}
func (r *Repository) UpdateEmployee(ctx context.Context, id int64, in model.EmployeeInput) (model.Employee,error) {
	_,err := r.DB.Exec(ctx, `UPDATE employees SET name=$1,email=$2,age=$3,salary=$4,department_id=$5 WHERE id=$6`,in.Name,in.Email,in.Age,in.Salary,in.DepartmentID,id)
	if err != nil { return model.Employee{},err }; return r.GetEmployee(ctx,id)
}
func (r *Repository) DeleteEmployee(ctx context.Context, id int64) error {
	tag,err := r.DB.Exec(ctx, `DELETE FROM employees WHERE id=$1`,id); if err != nil { return err }
	if tag.RowsAffected()==0 { return fmt.Errorf("employee not found") }; return nil
}
func (r *Repository) ListDepartments(ctx context.Context) ([]model.Department,error) {
	rows,err:=r.DB.Query(ctx,`SELECT id,name FROM departments ORDER BY name`); if err!=nil{return nil,err}; defer rows.Close()
	out:=[]model.Department{}; for rows.Next(){var d model.Department;if err:=rows.Scan(&d.ID,&d.Name);err!=nil{return nil,err};out=append(out,d)};return out,rows.Err()
}
func (r *Repository) GetDepartment(ctx context.Context,id int64)(model.Department,error){
	var d model.Department; err:=r.DB.QueryRow(ctx,`SELECT id,name FROM departments WHERE id=$1`,id).Scan(&d.ID,&d.Name);return d,err
}
