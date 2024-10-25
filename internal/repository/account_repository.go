package repository

import "BankSystem/internal/model"

type AccountRepository struct {
	Accounts []model.Account
}

func (r *AccountRepository) Create(account model.Account) (model.Account, error) {
	account.ID = int64(len(r.Accounts) + 1)
	r.Accounts = append(r.Accounts, account)

	return account, nil
}
