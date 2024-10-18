package main

import (
	"fmt"
	"log"
	"order-service/configs"
	"order-service/internal/delivery"
	"order-service/internal/domain"
	"order-service/internal/repository"
	"order-service/internal/usecase"
	"order-service/package/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config")
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.Name)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database %v", err)
	}

	if err := db.AutoMigrate(&domain.Order{}, &domain.Product{}); err != nil {
		log.Fatalf("migration error: %v", err)
	}

	orderRepo := repository.NewGormOrderRepository(db)
	orderService := usecase.NewOrderService(orderRepo)
	orderController := delivery.NewOrderController(orderService)

	r := gin.Default()

	r.Use(middleware.AuthMiddleware())
	authRoutes := r.Group("/orders")
	{
		authRoutes.POST("/create", orderController.CreateOrder)
		authRoutes.GET("/get", orderController.GetOrders)
	}

	r.Run(":8081")
}
