package repository

import (
	"auth-service/internal/domain"
	"auth-service/internal/dto"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type GormUserRepository struct {
	DB *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) UserRepository {
	return &GormUserRepository{DB: db}
}

func (r *GormUserRepository) RegisterUser(user *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Passsword), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	dbUser := &domain.User{
		Username: user.Username,
		Email:    user.Email,
		Password: string(hashedPassword),
	}

	if err := r.DB.Create(dbUser).Error; err != nil {
		return nil, err
	}

	return &dto.RegisterResponse{Message: ""}, nil
}

func (r *GormUserRepository) LoginUser(user *dto.LoginRequest) (*dto.LoginResponse, error) {
	return &dto.LoginResponse{Token: ""}, nil
}

func (r *GormUserRepository) GetByEmail(email string) (*dto.GetByEmailResponse, error) {
	return &dto.GetByEmailResponse{}, nil
}
