package repository

import (
	"employee-platform/model"
	"gorm.io/gorm"
)

type UserRepository struct{ db *gorm.DB }

func NewUserRepository(db *gorm.DB) *UserRepository     { return &UserRepository{db: db} }
func (r *UserRepository) Create(user *model.User) error { return r.db.Create(user).Error }
func (r *UserRepository) ByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}
func (r *UserRepository) ByID(id uint) (*model.User, error) {
	var user model.User
	err := r.db.First(&user, id).Error
	return &user, err
}
