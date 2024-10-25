package main

import (
	"BankSystem/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

var accounts []model.Account

func main() {
	// Initializing the router
	router := gin.Default()

	router.POST("/accounts", createAccount)
	router.GET("/accounts", getAccounts)

	router.Run(":9090")
}

// the *gin.Context is a reference to the actual Gin Context not a copy.
func createAccount(context *gin.Context) {
	var newAccount model.Account

	// ShouldBindJSON takes the data from the client's request and checks if
	// it is in the correct JSON format, if it is then it will return nil and fill the newData.
	// we use &(address of) because the ShouldBindJSON fills the newData with the
	// Client's request
	if err := context.ShouldBindJSON(&newAccount); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newAccount.ID = int64(len(accounts) + 1)
	accounts = append(accounts, newAccount)

	context.JSON(http.StatusCreated, newAccount)
}

func getAccounts(context *gin.Context) {
	context.JSON(http.StatusOK, accounts)
}
