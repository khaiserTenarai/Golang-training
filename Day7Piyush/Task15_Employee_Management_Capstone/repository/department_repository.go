package repository

import (
	"database/sql"
	"task15_employee_management_capstone/models"
)

// DepartmentRepository interface
type DepartmentRepository interface {
	Create(dept models.Department) (models.Department, error)
	GetAll() ([]models.Department, error)
	GetByID(id int) (models.Department, error)
	Update(dept models.Department) (models.Department, error)
	Delete(id int) error
}

// PostgresDepartmentRepository implements DepartmentRepository
type PostgresDepartmentRepository struct {
	DB *sql.DB
}

func NewPostgresDepartmentRepository(db *sql.DB) DepartmentRepository {
	return &PostgresDepartmentRepository{DB: db}
}

func (r *PostgresDepartmentRepository) Create(dept models.Department) (models.Department, error) {
	err := r.DB.QueryRow(
		`INSERT INTO cap_departments (name, location) VALUES ($1, $2)
		RETURNING id, name, location, created_at`,
		dept.Name, dept.Location).
		Scan(&dept.ID, &dept.Name, &dept.Location, &dept.CreatedAt)
	return dept, err
}

func (r *PostgresDepartmentRepository) GetAll() ([]models.Department, error) {
	rows, err := r.DB.Query(`SELECT id, name, location, created_at FROM cap_departments ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var depts []models.Department
	for rows.Next() {
		var d models.Department
		if err := rows.Scan(&d.ID, &d.Name, &d.Location, &d.CreatedAt); err != nil {
			return nil, err
		}
		depts = append(depts, d)
	}
	return depts, rows.Err()
}

func (r *PostgresDepartmentRepository) GetByID(id int) (models.Department, error) {
	var d models.Department
	err := r.DB.QueryRow(
		`SELECT id, name, location, created_at FROM cap_departments WHERE id=$1`, id).
		Scan(&d.ID, &d.Name, &d.Location, &d.CreatedAt)
	return d, err
}

func (r *PostgresDepartmentRepository) Update(dept models.Department) (models.Department, error) {
	err := r.DB.QueryRow(
		`UPDATE cap_departments SET name=$1, location=$2 WHERE id=$3
		RETURNING id, name, location, created_at`,
		dept.Name, dept.Location, dept.ID).
		Scan(&dept.ID, &dept.Name, &dept.Location, &dept.CreatedAt)
	return dept, err
}

func (r *PostgresDepartmentRepository) Delete(id int) error {
	result, err := r.DB.Exec(`DELETE FROM cap_departments WHERE id=$1`, id)
	if err != nil {
		return err
	}
	n, _ := result.RowsAffected()
	if n == 0 {
		return sql.ErrNoRows
	}
	return nil
}
