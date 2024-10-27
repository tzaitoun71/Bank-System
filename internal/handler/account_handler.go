package handler

import (
	"BankSystem/internal/model"
	"BankSystem/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AccountHandler struct {
	Service *service.AccountService
}

// Create Account
func (h *AccountHandler) CreateAccount(c *gin.Context) {
	// Initialize an Account
	var account model.Account

	// This checks if the JSON is valid and if it is it loads it into the account variable
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Call the service with this account as the input
	newAccount, err := h.Service.CreateAccount(account)

	// Check if the newAccount is valid from the service
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Successful Account creation
	c.JSON(http.StatusCreated, newAccount)
}

// Get all Accounts
func (h *AccountHandler) GetAccounts(c *gin.Context) {
	accounts := h.Service.GetAllAccounts()
	c.JSON(http.StatusOK, accounts)
}

// Get Account by ID
func (h *AccountHandler) GetAccountByID(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	account, found := h.Service.GetByID(id)

	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		return
	}
	c.JSON(http.StatusOK, account)
}

// Update an account
func (h *AccountHandler) UpdateAccount(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var account model.Account
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	account.ID = id
	updatedAccount, _ := h.Service.Update(account)
	c.JSON(http.StatusOK, updatedAccount)
}

// Delete an Account

func (h *AccountHandler) DeleteAccount(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	err := h.Service.Delete(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": "Account Deleted"})
}
