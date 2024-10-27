package service

import (
	"BankSystem/internal/model"
	"BankSystem/internal/repository"
	"errors"
)

type AccountService struct {
	Repo *repository.AccountRepository
}

// CreateAccount handles business logic for account creation
func (s *AccountService) CreateAccount(account model.Account) (model.Account, error) {
	if account.Balance < 0 {
		return model.Account{}, errors.New("balance cannot be negative")
	}

	return s.Repo.Create(account)
}

// GetAllAccounts retrieves all accounts
func (s *AccountService) GetAllAccounts() []model.Account {
	return s.Repo.GetAll()
}

// GetAccountById
func (s *AccountService) GetByID(id int64) (model.Account, bool) {
	return s.Repo.GetByID(id)
}

func (s *AccountService) Update(account model.Account) (model.Account, error) {
	return s.Repo.Update(account)
}

func (s *AccountService) Delete(id int64) error {
	return s.Repo.Delete(id)
}
