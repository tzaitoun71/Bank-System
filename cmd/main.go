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

	// Initialize repository, service, and handler

	accountRepo := &repository.AccountRepository{}
	accountService := &service.AccountService{Repo: accountRepo}
	accountHandler := &handler.AccountHandler{Service: accountService}

	router.POST("/accounts", accountHandler.CreateAccount)
	router.GET("/accounts", accountHandler.GetAccounts)
	router.GET("/accounts/:id", accountHandler.GetAccountByID)
	router.PUT("/accounts/:id", accountHandler.UpdateAccount)
	router.DELETE("/accounts/:id", accountHandler.DeleteAccount)

	router.Run(":9090")
}
