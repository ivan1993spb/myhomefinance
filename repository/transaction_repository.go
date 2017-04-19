package repository

import (
	"time"

	"github.com/ivan1993spb/myhomefinance/models"
)

type TransactionsRepository interface {
	CreateTransaction(t *models.Transaction) error
	UpdateTransaction(t *models.Transaction) error
	DeleteTransaction(t *models.Transaction) error
	GetAccountTransactionsByTimeRange(accountID uint64, from time.Time, to time.Time) ([]*models.Transaction, error)
	GetAccountTransactionsByTimeRangeCategories(accountID uint64, from time.Time, to time.Time, categories []string) ([]*models.Transaction, error)
	GetAccountStatsByTimeRange(accountID uint64, from time.Time, to time.Time) (float64, float64, float64, uint64)
	GetAccountStatsByTimeRangeCategories(accountID uint64, from time.Time, to time.Time, categories []string) (float64, float64, float64, uint64)
	CountAccountCategoriesSumsByTimeRange(accountID uint64, from time.Time, to time.Time) ([]*models.CategorySum, error)
}
