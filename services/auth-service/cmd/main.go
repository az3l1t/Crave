package main

import (
	"auth-service/configs"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg, err := configs.LoadConfig("/configs/config.yaml")
	if err != nil {
		log.Fatalf("Error loading config %v", err)
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.Name)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database %v", err)
	}

	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatalf("Error closing database connection %v", err)
		}
		if err := sqlDB.Close(); err != nil {
			log.Fatalf("Error closing database")
		}
	}()
}
