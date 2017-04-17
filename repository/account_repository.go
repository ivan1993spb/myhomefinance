package repository

import "github.com/ivan1993spb/myhomefinance/models"

type AccountRepository interface {
	CreateTransaction(a *models.Account) error
	UpdateTransaction(a *models.Account) error
	DeleteTransaction(a *models.Account) error
}
