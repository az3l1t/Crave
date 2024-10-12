package usecase

import (
	"auth-service/internal/domain"
	"auth-service/internal/dto"
	"auth-service/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repository repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{Repository: repo}
}

func (s *UserService) RegisterUser(userDto *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userDto.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Username: userDto.Username,
		Email:    userDto.Email,
		Password: string(hashedPassword),
	}

	if err := s.Repository.Save(user); err != nil {
		return nil, err
	}

	return &dto.RegisterResponse{Message: "User registered successfully"}, nil
}
