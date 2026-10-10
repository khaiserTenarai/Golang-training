package repository

import (
	"context"

	"student-management/model"
)

// StudentRepository defines database operations.
type StudentRepository interface {
	GetFiltered(ctx context.Context, name, grade string, limit, offset int) ([]model.Student, error)
	CountFiltered(ctx context.Context, name, grade string) (int64, error)
	Create(ctx context.Context, student model.Student) (model.Student, error)
}
