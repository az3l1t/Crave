package repository

import (
	"auth-service/internal/dto"
)

type UserRepository interface {
	RegisterUser(user *dto.RegisterRequest) (*dto.RegisterResponse, error)
	LoginUser(user *dto.LoginRequest) (*dto.LoginResponse, error)
	GetByEmail(email string) (*dto.GetByEmailResponse, error)
}
