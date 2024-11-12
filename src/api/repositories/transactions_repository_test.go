package repositories

import (
	"fmt"
	"github.com/grosinov/transactions-api/src/api/models"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"regexp"
	"testing"
	"time"
)

var (
	now                 = time.Now()
	mockTransaction     = models.Transaction{Id: 1, UserId: 1, Amount: 10, Datetime: now}
	mockTransactionList = &[]models.Transaction{mockTransaction}
)

func SetupMockDB() (*gorm.DB, sqlmock.Sqlmock, error) {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, nil, err
	}

	return gormDB, mock, nil
}

func TestBulkCreateTransactions(t *testing.T) {
	db, mock, err := SetupMockDB()
	if err != nil {
		t.Fatalf("Error initializing mock DB: %v", err)
	}
	defer db.Exec("")

	transactionRepository := &TransactionsRepositoryImpl{db: db}

	t.Run("success", func(t *testing.T) {
		addRow := sqlmock.NewRows([]string{"id", "user_id", "amount", "datetime"}).AddRow(mockTransaction.Id, mockTransaction.UserId, mockTransaction.Amount, mockTransaction.Datetime)
		expectedSQL := "INSERT INTO \"transactions\" (.+) VALUES (.+)"
		mock.ExpectBegin()
		mock.ExpectQuery(expectedSQL).WillReturnRows(addRow)
		mock.ExpectCommit()
		createdTransactions, err := transactionRepository.BulkCreateTransactions(mockTransactionList)
		assert.Nil(t, err)

		assert.Equal(t, mockTransactionList, createdTransactions)
	})

	t.Run("failure", func(t *testing.T) {
		expectedSQL := "INSERT INTO \"transactions\" (.+) VALUES (.+)"
		mock.ExpectBegin()
		mock.ExpectQuery(expectedSQL).WillReturnError(fmt.Errorf("error"))
		mock.ExpectCommit()
		_, err := transactionRepository.BulkCreateTransactions(mockTransactionList)
		assert.NotNil(t, err)
	})
}

func TestGetTransactions(t *testing.T) {
	db, mock, err := SetupMockDB()
	if err != nil {
		t.Fatalf("Error initializing mock DB: %v", err)
	}
	defer db.Exec("")

	transactionRepository := &TransactionsRepositoryImpl{db: db}

	t.Run("success", func(t *testing.T) {
		row := sqlmock.NewRows([]string{"id", "user_id", "amount", "datetime"}).AddRow(mockTransaction.Id, mockTransaction.UserId, mockTransaction.Amount, mockTransaction.Datetime)
		expectedSQL := regexp.QuoteMeta(`SELECT * FROM "transactions" WHERE user_id = $1`)
		mock.ExpectBegin()
		mock.ExpectQuery(expectedSQL).WillReturnRows(row)
		transactions, err := transactionRepository.GetTransactions(1, nil, nil)
		assert.Nil(t, err)

		assert.Equal(t, mockTransactionList, transactions)
	})

	t.Run("success when between dates", func(t *testing.T) {
		specificDate := time.Date(2024, 11, 11, 12, 0, 0, 0, time.FixedZone("UDT", -11*60*60))
		from := time.Date(2024, 11, 10, 12, 0, 0, 0, time.FixedZone("UDT", -11*60*60))
		to := time.Date(2024, 11, 12, 12, 0, 0, 0, time.FixedZone("UDT", -11*60*60))
		row := sqlmock.NewRows([]string{"id", "user_id", "amount", "datetime"}).AddRow(mockTransaction.Id, mockTransaction.UserId, mockTransaction.Amount, specificDate)
		expectedSQL := regexp.QuoteMeta(`SELECT * FROM "transactions" WHERE user_id = $1 AND datetime >= $2 AND datetime <= $3`)
		mock.ExpectBegin()
		mock.ExpectQuery(expectedSQL).WillReturnRows(row)
		transactions, err := transactionRepository.GetTransactions(1, &from, &to)
		assert.Nil(t, err)

		expectedTransactionList := &[]models.Transaction{{Id: 1, UserId: 1, Amount: 10, Datetime: specificDate}}
		assert.Equal(t, expectedTransactionList, transactions)
	})

	t.Run("failure", func(t *testing.T) {
		expectedSQL := regexp.QuoteMeta(`SELECT * FROM "transactions" WHERE user_id = $1`)
		mock.ExpectBegin()
		mock.ExpectQuery(expectedSQL).WillReturnError(fmt.Errorf("error"))
		mock.ExpectCommit()
		_, err := transactionRepository.GetTransactions(1, nil, nil)
		assert.NotNil(t, err)
	})
}
