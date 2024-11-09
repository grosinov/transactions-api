package config

import (
	"fmt"
	"github.com/grosinov/transactions-api/src/api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

const (
	DateTimeLayout = "2006-01-02T15:04:05MST"
)

func ConnectDatabase() *gorm.DB {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	err = db.AutoMigrate(&models.Transaction{})
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	return db
}
