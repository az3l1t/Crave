package usecase

import (
	"auth-service/internal/domain"
	"auth-service/internal/dto"
	"auth-service/internal/repository"
	"auth-service/package/utils"
	"errors"

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

/*
TODO LoginUser function implementation
*/
func (s *UserService) LoginUser(userDto *dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := s.Repository.GetByEmail(userDto.Email)
	if err != nil {
		return nil, errors.New("first")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userDto.Passsword)); err != nil {
		return nil, errors.New("2")
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return nil, errors.New("3")
	}

	return &dto.LoginResponse{Token: token}, nil
}
