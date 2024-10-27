package repository

import (
	"BankSystem/internal/model"
	"time"
)

type TransactionRepository struct {
	Transactions []model.Transaction
}

// Create a Transaction
func (r *TransactionRepository) Create(transaction model.Transaction) (model.Transaction, error) {
	transaction.TxID = int64(len(r.Transactions) + 1)
	transaction.Timestamp = time.Now()
	r.Transactions = append(r.Transactions, transaction)

	return transaction, nil
}

// Get All Transactions
func (r *TransactionRepository) GetAll() []model.Transaction {
	return r.Transactions
}

// Get Transactions by Account  ID
func (r *TransactionRepository) GetByAccountID(accountID int64) []model.Transaction {
	var accountTransactions []model.Transaction
	for _, transaction := range r.Transactions {
		if accountID == transaction.AccountID {
			accountTransactions = append(accountTransactions, transaction)
		}
	}
	return accountTransactions
}
