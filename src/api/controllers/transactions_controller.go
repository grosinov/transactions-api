package controllers

import (
	"encoding/csv"
	"github.com/gin-gonic/gin"
	"github.com/grosinov/transactions-api/src/api/config"
	"github.com/grosinov/transactions-api/src/api/models"
	"github.com/grosinov/transactions-api/src/api/services"
	"io"
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

func (c *TransactionsController) MigrateTransactions(ctx *gin.Context) {
	csvReader := csv.NewReader(ctx.Request.Body)
	csvReader.Comma = ','
	defer ctx.Request.Body.Close()

	transactions, err := parseTransactions(csvReader)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Failed to parse CSV",
			"details": err.Error(),
		})
		return
	}

	createdTransactions, err := c.service.BulkCreateTransactions(transactions)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to save transactions",
			"details": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, createdTransactions)
}

func (c *TransactionsController) GetBalance(ctx *gin.Context) {}

func parseTransactions(csvReader *csv.Reader) (*[]models.Transaction, error) {
	var transactions []models.Transaction
	isHeader := true

	for {
		row, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		if isHeader {
			isHeader = false
			continue
		}

		id, _ := strconv.ParseUint(row[0], 10, 64)
		userId, _ := strconv.ParseUint(row[1], 10, 64)
		amount, _ := strconv.ParseFloat(row[2], 64)
		dateTime, _ := time.Parse(config.DateTimeLayout, row[1])

		transaction := models.Transaction{
			ID:       id,
			UserId:   userId,
			Amount:   amount,
			Datetime: dateTime,
		}

		transactions = append(transactions, transaction)
	}

	return &transactions, nil
}
