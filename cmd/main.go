package main

import (
	"BankSystem/internal/handler"
	"BankSystem/internal/repository"
	"BankSystem/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initializing the router
	router := gin.Default()

	// Initialize repository, service, and handler for Account
	accountRepo := &repository.AccountRepository{}
	accountService := &service.AccountService{Repo: accountRepo}
	accountHandler := &handler.AccountHandler{Service: accountService}

	// Initialize repository, service, and handler for Account
	transactionRepo := &repository.TransactionRepository{}
	TransactionService := &service.TransactionService{TransactionRepo: transactionRepo, AccountService: accountService}
	transactionHandler := &handler.TransactionHandler{Service: TransactionService}

	// Routes for Accounts
	router.POST("/accounts", accountHandler.CreateAccount)
	router.GET("/accounts", accountHandler.GetAccounts)
	router.GET("/accounts/:id", accountHandler.GetAccountByID)
	router.PUT("/accounts/:id", accountHandler.UpdateAccount)
	router.DELETE("/accounts/:id", accountHandler.DeleteAccount)

	// Routes for Transactions
	router.POST("/transactions", transactionHandler.CreateTransaction)
	router.GET("/transactions/account/:account_id", transactionHandler.GetTransactionByAccountID)

	router.Run(":9090")
}
