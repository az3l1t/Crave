package domain

import "time"

type Order struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	Products  []Product `json:"products" gorm:"many2many:order_products"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Product struct {
	ID       uint    `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

type OrderProduct struct {
	OrderID   uint `gorm:"primaryKey"`
	ProductID uint `gorm:"primaryKey"`
}
