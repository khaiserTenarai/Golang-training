package repository
import("context";"errors";"q12-repository-service-architecture/model";"github.com/jackc/pgx/v5";"github.com/jackc/pgx/v5/pgxpool")
var ErrNotFound=errors.New("employee not found");type EmployeeRepositoryImpl struct{db *pgxpool.Pool}
func NewEmployeeRepository(db *pgxpool.Pool)EmployeeRepository{return &EmployeeRepositoryImpl{db}}
func(r *EmployeeRepositoryImpl)CreateEmployee(e model.Employee)error{_,x:=r.db.Exec(context.Background(),`INSERT INTO employees(name,email,salary)VALUES($1,$2,$3)`,e.Name,e.Email,e.Salary);return x}
func(r *EmployeeRepositoryImpl)GetEmployee(id int)(*model.Employee,error){var e model.Employee;x:=r.db.QueryRow(context.Background(),`SELECT id,name,email,salary FROM employees WHERE id=$1`,id).Scan(&e.ID,&e.Name,&e.Email,&e.Salary);if x==pgx.ErrNoRows{return nil,ErrNotFound};return &e,x}
func(r *EmployeeRepositoryImpl)GetAllEmployees()([]model.Employee,error){rows,x:=r.db.Query(context.Background(),`SELECT id,name,email,salary FROM employees ORDER BY id`);if x!=nil{return nil,x};defer rows.Close();var out []model.Employee;for rows.Next(){var e model.Employee;if x=rows.Scan(&e.ID,&e.Name,&e.Email,&e.Salary);x!=nil{return nil,x};out=append(out,e)};return out,rows.Err()}
func(r *EmployeeRepositoryImpl)UpdateEmployee(e model.Employee)error{res,x:=r.db.Exec(context.Background(),`UPDATE employees SET name=$1,email=$2,salary=$3 WHERE id=$4`,e.Name,e.Email,e.Salary,e.ID);if x!=nil{return x};if res.RowsAffected()==0{return ErrNotFound};return nil}
func(r *EmployeeRepositoryImpl)DeleteEmployee(id int)error{res,x:=r.db.Exec(context.Background(),`DELETE FROM employees WHERE id=$1`,id);if x!=nil{return x};if res.RowsAffected()==0{return ErrNotFound};return nil}
