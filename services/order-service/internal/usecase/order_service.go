package usecase

import (
	"order-service/internal/domain"
	"order-service/internal/dto"
	"order-service/internal/repository"
)

type OrderService struct {
	Repo repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) *OrderService {
	return &OrderService{Repo: repo}
}

func (s *OrderService) CreateOrder(order *domain.Order) (*dto.CreateOrderResponse, error) {
	if err := s.Repo.Create(order); err != nil {
		return nil, err
	}

	return &dto.CreateOrderResponse{ID: order.ID, Message: "Order created successfully"}, nil
}

func (s *OrderService) GetByID(id uint) (*dto.GetOrderResponse, error) {
	order, err := s.Repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return &dto.GetOrderResponse{
		ID:       order.ID,
		UserID:   order.UserID,
		Products: order.Products,
	}, nil
}

func (s *OrderService) GetOrders(userID uint) ([]domain.Order, error) {
	return s.Repo.GetAll(userID)
}
