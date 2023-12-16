package repository

import (
	"errors"
	"strings"

	"github.com/BerdanAkbulut/task-app-backend/entity"
)

type userRepository struct {
}

type UserRepository interface {
	Save(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (u *userRepository) Save(user *entity.User) error {
	user.HashPassword()
	res := DB().Create(&user)
	if strings.Contains(res.Error.Error(), "Duplicate entry") {
		return errors.New("User already exists with this email address: " + user.Email)
	}
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (u *userRepository) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	result := DB().Where("email = ?", email).Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("User not found by email address: " + email)
	}
	return &user, nil
}
