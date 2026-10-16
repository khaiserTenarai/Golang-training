package repository

import (
	"employee-management/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(*model.User) error
	GetByUsername(string) (*model.User, error)
	Count() (int64, error)
}
type userRepository struct{ db *gorm.DB }

func NewUserRepository(db *gorm.DB) UserRepository { return &userRepository{db} }
func (r *userRepository) Create(u *model.User) error { return r.db.Create(u).Error }
func (r *userRepository) GetByUsername(name string) (*model.User, error) {
	var u model.User
	if err := r.db.Where("username = ?", name).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}
func (r *userRepository) Count() (int64, error) {
	var n int64
	err := r.db.Model(&model.User{}).Count(&n).Error
	return n, err
}
