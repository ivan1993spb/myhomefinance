package repository

import (
	"github.com/ivan1993spb/myhomefinance/models"
	"github.com/satori/go.uuid"
)

type AccountRepository interface {
	CreateAccount(a *models.Account) error
	UpdateAccount(a *models.Account) error
	DeleteAccount(a *models.Account) error
	GetAccountsByUserUUID(userUUID uuid.UUID) ([]*models.Account, error)
}
