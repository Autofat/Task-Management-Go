package repository

import (
	"task-management/internal/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) FindByID(id uint) (*model.User, error) {
	var user model.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Update(id uint, updates *model.User) error {

	user, err := r.FindByID(id)
	if err != nil {
		return err
	}
	if updates.Email != "" {
		user.Email = updates.Email
	}
	if updates.Fullname != "" {
		user.Fullname = updates.Fullname
	}
	if updates.Role != "" {
		user.Role = updates.Role
	}
	return r.db.Save(user).Error
}

func (r *UserRepository) DeleteById(id uint) error {
	return r.db.Unscoped().Delete(&model.User{}, id).Error
}