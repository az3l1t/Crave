package main

import (
	"auth-service/configs"
	"auth-service/internal/delivery"
	"auth-service/internal/domain"
	"auth-service/internal/repository"
	"auth-service/internal/usecase"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config %v", err)
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.Name)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database %v", err)
	}

	if err := db.AutoMigrate(&domain.User{}); err != nil {
		log.Fatalf("migration error: %v", err)
	}

	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatalf("Error retrieving database connection %v", err)
		}
		if err := sqlDB.Close(); err != nil {
			log.Fatalf("Error closing database connection %v", err)
		}
	}()

	userRepo := repository.NewGormUserRepository(db)
	userService := usecase.NewUserService(userRepo)
	authController := delivery.NewAuthController(userService)

	r := gin.Default()

	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/register", authController.RegisterUser)
		authRoutes.POST("/login", authController.LoginUser)
	}

	r.Run(":8080")
}
