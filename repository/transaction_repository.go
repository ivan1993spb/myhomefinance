package repository

import (
	"time"

	"github.com/ivan1993spb/myhomefinance/models"
)

type TransactionsRepository interface {
	GetTransactionsByTimeRange(from time.Time, to time.Time) ([]*models.Transaction, error)
	CreateTransaction(t *models.Transaction) error
	UpdateTransaction(t *models.Transaction) error
}
