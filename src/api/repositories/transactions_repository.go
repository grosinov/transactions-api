package repositories

import (
	"github.com/grosinov/transactions-api/src/api/models"
	"gorm.io/gorm"
	"time"
)

type TransactionsRepository interface {
	BulkCreateTransactions(*[]models.Transaction) (*[]models.Transaction, error)
	GetTransactions(userId uint64, from, to *time.Time) (*[]models.Transaction, error)
}

type TransactionsRepositoryImpl struct {
	db *gorm.DB
}

func (r *TransactionsRepositoryImpl) BulkCreateTransactions(transactions *[]models.Transaction) (*[]models.Transaction, error) {
	if err := r.db.Create(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *TransactionsRepositoryImpl) GetTransactions(userId uint64, from, to *time.Time) (*[]models.Transaction, error) {
	var transactions []models.Transaction
	query := r.db.Model(&models.Transaction{})

	if from != nil && !from.IsZero() {
		query = query.Where("date >= ?", *from)
	}

	if to != nil && !to.IsZero() {
		query = query.Where("date <= ?", *to)
	}

	if err := query.Find(&transactions).Error; err != nil {
		return nil, err
	}

	return &transactions, nil
}

func NewTransactionsRepository(db *gorm.DB) TransactionsRepository {
	return &TransactionsRepositoryImpl{db: db}
}
