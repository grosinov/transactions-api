package services

import (
	"github.com/grosinov/transactions-api/src/api/dtos"
	"github.com/grosinov/transactions-api/src/api/models"
	"github.com/grosinov/transactions-api/src/api/repositories"
	"time"
)

type TransactionsService interface {
	BulkCreateTransactions(*[]models.Transaction) (*[]models.Transaction, error)
	GetBalance(userId uint64, from, to *time.Time) (*dtos.Balance, error)
	GetTransactions(userId uint64, from, to *time.Time) (*[]models.Transaction, error)
}

type TransactionsServiceImpl struct {
	repository repositories.TransactionsRepository
}

func (s *TransactionsServiceImpl) BulkCreateTransactions(transactions *[]models.Transaction) (*[]models.Transaction, error) {
	createdTransactions, err := s.repository.BulkCreateTransactions(transactions)
	if err != nil {
		return nil, err
	}

	return createdTransactions, nil
}

func (s *TransactionsServiceImpl) GetBalance(userId uint64, from, to *time.Time) (*dtos.Balance, error) {
	transactions, err := s.GetTransactions(userId, from, to)
	if err != nil {
		return nil, err
	}

	balance := &dtos.Balance{}

	for _, transaction := range *transactions {
		balance.Balance += transaction.Amount

		if transaction.Amount > 0 {
			balance.TotalCredit++
			continue
		}

		balance.TotalDebit++
	}

	return balance, nil
}

func (s *TransactionsServiceImpl) GetTransactions(userId uint64, from, to *time.Time) (*[]models.Transaction, error) {
	transactions, err := s.repository.GetTransactions(userId, from, to)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func NewTransactionsServiceImpl(repository repositories.TransactionsRepository) TransactionsService {
	return &TransactionsServiceImpl{repository: repository}
}
