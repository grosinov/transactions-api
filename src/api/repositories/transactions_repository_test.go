package repositories

import (
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
)

func TestBulkCreateTransactions(t *testing.T) {
	mockDB, mock, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDB,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})
	transactionsRepository := &TransactionsRepositoryImpl{db}

	t.Run("Success", func(t *testing.T) {

	})
}
