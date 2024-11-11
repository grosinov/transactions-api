package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/grosinov/transactions-api/src/api/dtos"
	"github.com/grosinov/transactions-api/src/api/models"
	"github.com/grosinov/transactions-api/src/api/services"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var (
	now                 = time.Now()
	mockTransaction     = models.Transaction{Id: 1, UserId: 1, Amount: 10, Datetime: now}
	mockTransactionList = &[]models.Transaction{mockTransaction}
	mockBalance         = &dtos.Balance{Balance: 10, TotalCredit: 1, TotalDebit: 0}
)

func TestTransactionsController_MigrateTransactions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	transactionsServiceMock := services.NewMockTransactionsService(ctrl)
	transactionsControllerImpl := NewTransactionsController(transactionsServiceMock)

	t.Run("success", func(t *testing.T) {
		transactionsServiceMock.EXPECT().BulkCreateTransactions(gomock.Any()).Return(mockTransactionList, nil)

		gin.SetMode(gin.TestMode)

		csvBody := fmt.Sprintf("id,user_id,amount,datetime\n1,1,10,%s", now)
		req := httptest.NewRequest(http.MethodPost, "/migrate", bytes.NewBufferString(csvBody))
		req.Header.Set("Content-Type", "text/csv")

		recorder := httptest.NewRecorder()

		ctx, _ := gin.CreateTestContext(recorder)
		ctx.Request = req

		transactionsControllerImpl.MigrateTransactions(ctx)

		var transactionsResponse []models.Transaction

		err := json.Unmarshal(recorder.Body.Bytes(), &transactionsResponse)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, recorder.Code)

		assert.Equal(t, uint64(1), transactionsResponse[0].Id)
		assert.Equal(t, uint64(1), transactionsResponse[0].UserId)
		assert.Equal(t, 10.0, transactionsResponse[0].Amount)
		assert.True(t, now.Equal(transactionsResponse[0].Datetime))
	})

	t.Run("csv read failure", func(t *testing.T) {
		gin.SetMode(gin.TestMode)

		csvBody := fmt.Sprintf("id,user_id,amount,datetime\n,,%s", now)

		req := httptest.NewRequest(http.MethodPost, "/migrate", bytes.NewBufferString(csvBody))

		req.Header.Set("Content-Type", "text/csv")
		recorder := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(recorder)
		ctx.Request = req

		transactionsControllerImpl.MigrateTransactions(ctx)

		var errorResponse dtos.ErrorResponse
		err := json.Unmarshal(recorder.Body.Bytes(), &errorResponse)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)

		assert.Equal(t, dtos.ErrorResponse(dtos.ErrorResponse{Message: "Error in CSV file: record on line 2: wrong number of fields"}), errorResponse)
	})

	t.Run("service failure", func(t *testing.T) {
		transactionsServiceMock.EXPECT().BulkCreateTransactions(gomock.Any()).Return(nil, fmt.Errorf("error"))

		gin.SetMode(gin.TestMode)

		csvBody := fmt.Sprintf("id,user_id,amount,datetime\n1,1,10,%s", now)

		req := httptest.NewRequest(http.MethodPost, "/migrate", bytes.NewBufferString(csvBody))

		req.Header.Set("Content-Type", "text/csv")
		recorder := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(recorder)
		ctx.Request = req

		transactionsControllerImpl.MigrateTransactions(ctx)

		var errorResponse dtos.ErrorResponse
		err := json.Unmarshal(recorder.Body.Bytes(), &errorResponse)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, recorder.Code)

		assert.Equal(t, dtos.ErrorResponse(dtos.ErrorResponse{Message: "Failed to create transactions: error"}), errorResponse)
	})
}

func TestTransactionsController_GetBalance(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	transactionsServiceMock := services.NewMockTransactionsService(ctrl)
	transactionsControllerImpl := NewTransactionsController(transactionsServiceMock)

	t.Run("success", func(t *testing.T) {
		transactionsServiceMock.EXPECT().GetBalance(gomock.Any(), gomock.Any(), gomock.Any()).Return(mockBalance, nil)

		gin.SetMode(gin.TestMode)

		req := httptest.NewRequest(http.MethodGet, "/users/1/balance", nil)

		recorder := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(recorder)
		ctx.Request = req

		transactionsControllerImpl.GetBalance(ctx)

		assert.Equal(t, http.StatusOK, recorder.Code)

		var balance dtos.Balance
		err := json.Unmarshal(recorder.Body.Bytes(), &balance)
		assert.NoError(t, err)

		assert.Equal(t, 10.0, balance.Balance)
		assert.Equal(t, 1, balance.TotalCredit)
		assert.Equal(t, 0, balance.TotalDebit)
	})
}
