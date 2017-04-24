package repository

import "github.com/ivan1993spb/myhomefinance/models"

type AccountRepository interface {
	CreateAccount(a *models.Account) error
	UpdateAccount(a *models.Account) error
	DeleteAccount(a *models.Account) error
	GetUserAccounts(userID uint64) ([]*models.Account, error)
}
