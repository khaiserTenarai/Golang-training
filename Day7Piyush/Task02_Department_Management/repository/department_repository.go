package repository

import (
	"database/sql"
	"task02_department_management/models"
)

type DepartmentRepository struct {
	DB *sql.DB
}

func NewDepartmentRepository(db *sql.DB) *DepartmentRepository {
	return &DepartmentRepository{DB: db}
}

func (r *DepartmentRepository) Create(dept models.Department) (int, error) {
	var id int
	query := `INSERT INTO departments (name, location) VALUES ($1, $2) RETURNING id`
	err := r.DB.QueryRow(query, dept.Name, dept.Location).Scan(&id)
	return id, err
}

func (r *DepartmentRepository) GetAll() ([]models.Department, error) {
	rows, err := r.DB.Query(`SELECT id, name, location, created_at FROM departments ORDER BY id`)
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

func (r *DepartmentRepository) GetByID(id int) (models.Department, error) {
	var d models.Department
	err := r.DB.QueryRow(`SELECT id, name, location, created_at FROM departments WHERE id=$1`, id).
		Scan(&d.ID, &d.Name, &d.Location, &d.CreatedAt)
	return d, err
}

func (r *DepartmentRepository) Update(dept models.Department) error {
	_, err := r.DB.Exec(`UPDATE departments SET name=$1, location=$2 WHERE id=$3`, dept.Name, dept.Location, dept.ID)
	return err
}

func (r *DepartmentRepository) Delete(id int) error {
	_, err := r.DB.Exec(`DELETE FROM departments WHERE id=$1`, id)
	return err
}
