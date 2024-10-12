package repository

import (
	"auth-service/internal/domain"

	"gorm.io/gorm"
)

type GormUserRepository struct {
	DB *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) UserRepository {
	return &GormUserRepository{DB: db}
}

func (r *GormUserRepository) Save(user *domain.User) error {
	return r.DB.Create(&user).Error
}

func (r *GormUserRepository) GetByEmail(email string) (*domain.User, error) {
	var user domain.User
	if err := r.DB.Where("email =?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
