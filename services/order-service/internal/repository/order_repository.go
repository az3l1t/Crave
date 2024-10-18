package repository

import (
	"order-service/internal/domain"

	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(order *domain.Order) error
	GetByID(id uint) (*domain.Order, error)
	GetAll(userID uint) ([]domain.Order, error)
}

type GormOrderRepository struct {
	DB *gorm.DB
}

func NewGormOrderRepository(db *gorm.DB) OrderRepository {
	return &GormOrderRepository{DB: db}
}

func (r *GormOrderRepository) Create(order *domain.Order) error {
	return r.DB.Create(order).Error
}

func (r *GormOrderRepository) GetByID(id uint) (*domain.Order, error) {
	var order domain.Order
	if err := r.DB.Preload("Products").First(&order, id).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *GormOrderRepository) GetAll(userID uint) ([]domain.Order, error) {
	var orders []domain.Order
	if err := r.DB.Preload("Products").Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}
