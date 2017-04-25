package repository

import (
	"time"

	"github.com/ivan1993spb/myhomefinance/models"
)

type TransactionRepository interface {
	CreateTransaction(t *models.Transaction) error
	UpdateTransaction(t *models.Transaction) error
	DeleteTransaction(t *models.Transaction) error
	GetTransactionByID(ID uint64) (*models.Transaction, error)
	GetAccountTransactionsByTimeRange(accountID uint64, from, to time.Time) ([]*models.Transaction, error)
	GetAccountTransactionsByTimeRangeCategories(accountID uint64, from, to time.Time, categories []string) ([]*models.Transaction, error)
	GetAccountStatsByTimeRange(accountID uint64, from, to time.Time) (*models.StatsTimeRange, error)
	GetAccountStatsByTimeRangeCategories(accountID uint64, from, to time.Time, categories []string) (*models.StatsTimeRangeCategories, error)
	CountAccountCategoriesSumsByTimeRange(accountID uint64, from, to time.Time) ([]*models.CategorySum, error)
}
