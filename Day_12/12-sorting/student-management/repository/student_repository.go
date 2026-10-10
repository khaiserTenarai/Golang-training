package repository

import (
	"context"
	"student-management/model"
)

type StudentRepository interface {
	Create(context.Context, model.Student) error
	GetFilteredSorted(context.Context, string, string, string, string, int, int) ([]model.Student, error)
	CountFiltered(context.Context, string, string) (int, error)
}
