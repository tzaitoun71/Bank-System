package service

import (
	"BankSystem/internal/model"
	"BankSystem/internal/repository"
	"errors"
)

type TransactionService struct {
	AccountService  *AccountService
	TransactionRepo *repository.TransactionRepository
}

func (s *TransactionService) CreateTransaction(transaction model.Transaction) (model.Transaction, error) {
	if transaction.Amount <= 0 {
		return model.Transaction{}, errors.New("transaction amount must be positive")
	}

	// Validate existence of account
	account, found := s.AccountService.GetByID(transaction.AccountID)
	if !found {
		return model.Transaction{}, errors.New("account not found")
	}

	switch transaction.TxType {
	case "Deposit":
		account.Balance += transaction.Amount
		s.AccountService.Update(account)

	case "Withdrawal":
		if account.Balance < transaction.Amount {
			return model.Transaction{}, errors.New("insufficient funds")
		}
		account.Balance -= transaction.Amount
		s.AccountService.Update(account)

	default:
		return model.Transaction{}, errors.New("invalid transaction type")
	}

	return s.TransactionRepo.Create(transaction)
}
