package repository

import "auth-service/internal/domain"

type UserRepository interface {
	Save(user *domain.User) error
	GetByEmail(email string) (*domain.User, error)
}
