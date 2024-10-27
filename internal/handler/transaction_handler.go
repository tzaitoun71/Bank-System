package handler

import (
	"BankSystem/internal/model"
	"BankSystem/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	Service *service.TransactionService
}

func (h *TransactionHandler) CreateTransaction(c *gin.Context) {
	var transaction model.Transaction
	if err := c.ShouldBindJSON(transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTransaction, err := h.Service.CreateTransaction(transaction)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newTransaction)
}

func (h *TransactionHandler) GetTransactionByAccountID(c *gin.Context) {
	accountID, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Account ID"})
		return
	}

	transactions := h.Service.TransactionRepo.GetByAccountID(accountID)
	c.JSON(http.StatusOK, transactions)
}
