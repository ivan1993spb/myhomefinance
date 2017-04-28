package repository

import (
	"github.com/satori/go.uuid"

	"github.com/ivan1993spb/myhomefinance/models"
)

type AccountRepository interface {
	CreateAccount(a *models.Account) error
	UpdateAccount(a *models.Account) error
	DeleteAccount(a *models.Account) error
	GetUserAccounts(userUUID uuid.UUID) ([]*models.Account, error)
}
