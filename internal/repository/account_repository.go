package repository

import (
	"BankSystem/internal/model"
	"errors"
)

type AccountRepository struct {
	Accounts []model.Account
}

// Create an account
func (r *AccountRepository) Create(account model.Account) (model.Account, error) {
	account.ID = int64(len(r.Accounts) + 1)
	r.Accounts = append(r.Accounts, account)

	return account, nil
}

// Get all accounts
func (r *AccountRepository) GetAll() []model.Account {
	return r.Accounts
}

// Get account by id
func (r *AccountRepository) GetByID(id int64) (model.Account, bool) {
	for _, account := range r.Accounts {
		if account.ID == id {
			return account, true
		}
	}
	return model.Account{}, false
}

// Update an account
func (r *AccountRepository) Update(account model.Account) (model.Account, error) {
	for i, acc := range r.Accounts {
		if acc.ID == account.ID {
			r.Accounts[i] = account
			return account, nil
		}
	}
	return model.Account{}, errors.New("accouint was not found")
}

// Delete an account
func (r *AccountRepository) Delete(id int64) error {
	for i, account := range r.Accounts {
		if account.ID == id {
			// creates a slice using all the accounts before the index and after the index
			r.Accounts = append(r.Accounts[:i], r.Accounts[i+1:]...)
			return nil
		}
	}
	return errors.New("account not found")
}
