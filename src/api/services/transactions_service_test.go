package services

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/grosinov/transactions-api/src/api/dtos"
	"github.com/grosinov/transactions-api/src/api/models"
	"github.com/grosinov/transactions-api/src/api/repositories"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	mockTransaction     = models.Transaction{Id: 1, UserId: 1, Amount: 10, Datetime: time.Now()}
	mockTransactionList = &[]models.Transaction{mockTransaction}
)

func TestTransactionsServiceImpl_BulkCreateTransactions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	transactionsRepositoryMock := repositories.NewMockTransactionsRepository(ctrl)
	transactionsServiceImpl := NewTransactionsServiceImpl(transactionsRepositoryMock)

	t.Run("success", func(t *testing.T) {

		transactionsRepositoryMock.EXPECT().BulkCreateTransactions(mockTransactionList).Return(mockTransactionList, nil)

		transactions, err := transactionsServiceImpl.BulkCreateTransactions(mockTransactionList)

		assert.Nil(t, err)
		assert.Equal(t, mockTransactionList, transactions)
	})

	t.Run("failure", func(t *testing.T) {
		transactionsRepositoryMock.EXPECT().BulkCreateTransactions(mockTransactionList).Return(nil, fmt.Errorf("error"))

		_, err := transactionsServiceImpl.BulkCreateTransactions(mockTransactionList)
		assert.NotNil(t, err)
	})
}

func TestTransactionsServiceImpl_GetTransaction(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	transactionsRepositoryMock := repositories.NewMockTransactionsRepository(ctrl)
	transactionsServiceImpl := NewTransactionsServiceImpl(transactionsRepositoryMock)

	t.Run("success", func(t *testing.T) {
		transactionsRepositoryMock.EXPECT().GetTransactions(uint64(1), nil, nil).Return(mockTransactionList, nil)

		transactions, err := transactionsServiceImpl.GetTransactions(uint64(1), nil, nil)

		assert.Nil(t, err)
		assert.Equal(t, mockTransactionList, transactions)
	})

	t.Run("failure", func(t *testing.T) {
		transactionsRepositoryMock.EXPECT().GetTransactions(uint64(1), nil, nil).Return(nil, fmt.Errorf("error"))

		_, err := transactionsServiceImpl.GetTransactions(uint64(1), nil, nil)

		assert.NotNil(t, err)
	})
}

func TestTransactionsServiceImpl_GetBalance(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	transactionsRepositoryMock := repositories.NewMockTransactionsRepository(ctrl)
	transactionsServiceImpl := NewTransactionsServiceImpl(transactionsRepositoryMock)

	t.Run("success", func(t *testing.T) {
		transactionsRepositoryMock.EXPECT().GetTransactions(uint64(1), nil, nil).Return(mockTransactionList, nil)

		balance, err := transactionsServiceImpl.GetBalance(uint64(1), nil, nil)

		assert.Nil(t, err)
		assert.Equal(t, &dtos.Balance{Balance: 10, TotalCredit: 1, TotalDebit: 0}, balance)
	})

	t.Run("failure", func(t *testing.T) {
		transactionsRepositoryMock.EXPECT().GetTransactions(uint64(1), nil, nil).Return(nil, fmt.Errorf("error"))
		_, err := transactionsServiceImpl.GetBalance(uint64(1), nil, nil)
		assert.NotNil(t, err)
	})
}
