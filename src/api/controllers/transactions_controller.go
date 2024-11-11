package controllers

import (
	"encoding/csv"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/grosinov/transactions-api/src/api/config"
	"github.com/grosinov/transactions-api/src/api/dtos"
	"github.com/grosinov/transactions-api/src/api/models"
	"github.com/grosinov/transactions-api/src/api/services"
	"net/http"
	"strconv"
	"time"
)

type TransactionsController struct {
	service services.TransactionsService
}

func NewTransactionsController(service services.TransactionsService) *TransactionsController {
	return &TransactionsController{service: service}
}

// @BasePath /api/v1

// MigrateTransactions godoc
// @Summary Migrate transactions
// @Schemes
// @Description Migrate transactions from CSV file
// @Tags example
// @Accept multipart/form-data
// @Produce json
// @Success 200 {array} models.Transaction
// @Failure 400 {object} dtos.ErrorResponse "Failed to parse CSV"
// @Failure 500 {object} dtos.ErrorResponse "Failed to save transactions"
// @Router /migrate [post]
func (c *TransactionsController) MigrateTransactions(ctx *gin.Context) {
	csvReader := csv.NewReader(ctx.Request.Body)
	csvReader.Comma = ','
	defer ctx.Request.Body.Close()

	transactions, err := parseTransactions(csvReader)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &dtos.ErrorResponse{
			Message: fmt.Sprintf("Error in CSV file: %s", err.Error()),
		})
		return
	}

	createdTransactions, err := c.service.BulkCreateTransactions(transactions)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &dtos.ErrorResponse{
			Message: fmt.Sprintf("Failed to create transactions: %s", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusOK, createdTransactions)
}

// GetBalance godoc
// @Summary Get Balance
// @Schemes
// @Description Get user's balance
// @Tags example
// @Accept json
// @Param userId path int true "User ID"
// @Param from query string false "Start date (yyyy-mm-ddThh:mm:ssZ)"
// @Param to query string false "End date (yyyy-mm-ddThh:mm:ssZ)"
// @Produce json
// @Success 200 {object} dtos.Balance
// @Failure 400 {object} dtos.ErrorResponse "Invalid user ID"
// @Failure 400 {object} dtos.ErrorResponse "Invalid from date"
// @Failure 400 {object} dtos.ErrorResponse "Invalid to date"
// @Failure 500 {object} dtos.ErrorResponse "Failed to get balance"
// @Router /users/{user_id}/balance [get]
func (c *TransactionsController) GetBalance(ctx *gin.Context) {
	userId, err := strconv.ParseUint(ctx.Param("user_id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &dtos.ErrorResponse{
			Message: fmt.Sprintf("Invalid user id %s", ctx.Param("user_id")),
		})
		return
	}

	fromDate, err := time.Parse(config.DateTimeLayout, ctx.DefaultQuery("from", time.Time{}.Format(config.DateTimeLayout)))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &dtos.ErrorResponse{
			Message: fmt.Sprintf("Invalid from date, format must be YYYY-MM-DDThh:mm:ssZ"),
		})
		return
	}

	toDate, err := time.Parse(config.DateTimeLayout, ctx.DefaultQuery("to", time.Time{}.Format(config.DateTimeLayout)))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &dtos.ErrorResponse{
			Message: fmt.Sprintf("Invalid from date, format must be YYYY-MM-DDThh:mm:ssZ"),
		})
		return
	}

	balance, err := c.service.GetBalance(userId, &fromDate, &toDate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &dtos.ErrorResponse{
			Message: fmt.Sprintf("Failed to get balance: %s", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusOK, balance)
}

func parseTransactions(csvReader *csv.Reader) (*[]models.Transaction, error) {
	var transactions []models.Transaction
	rows, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	for _, row := range rows[1:] {
		id, _ := strconv.ParseUint(row[0], 10, 64)
		userId, _ := strconv.ParseUint(row[1], 10, 64)
		amount, _ := strconv.ParseFloat(row[2], 64)
		dateTime, _ := time.Parse(config.DateTimeLayout, row[1])

		transaction := models.Transaction{
			Id:       id,
			UserId:   userId,
			Amount:   amount,
			Datetime: dateTime,
		}

		transactions = append(transactions, transaction)
	}

	return &transactions, nil
}
