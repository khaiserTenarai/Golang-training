package repository

import (
 "context"
 "student-management/model"
)

type StudentRepository interface {
 Create(context.Context, model.Student) error
 List(context.Context, string, string, string, string, int, int) ([]model.Student, error)
 Count(context.Context, string, string) (int, error)
}
