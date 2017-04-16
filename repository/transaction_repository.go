package repository

import (
	"time"

	"github.com/ivan1993spb/myhomefinance/models"
)

type TransactionsRepository interface {
	CreateTransaction(t *models.Transaction) error
	UpdateTransaction(t *models.Transaction) error
	DeleteTransaction(t *models.Transaction) error
	GetTransactionsByTimeRange(from time.Time, to time.Time) ([]*models.Transaction, error)
	GetTransactionsByTimeRangeCategories(from time.Time, to time.Time, categories []string) ([]*models.Transaction, error)
	StatsByTimeRange(from time.Time, to time.Time) (float64, float64, float64, uint64)

	// todo create foreach func
}
