package dto

import "order-service/internal/domain"

type CreateOrderRequest struct {
	Products []domain.Product `json:"products"`
}

type CreateOrderResponse struct {
	ID      uint   `json:"id"`
	Message string `json:"message"`
}

type GetOrderResponse struct {
	ID       uint             `json:"id"`
	UserID   uint             `json:"user_id"`
	Products []domain.Product `json:"products"`
}
