package repository

import (
	"time"

	"github.com/satori/go.uuid"

	"github.com/ivan1993spb/myhomefinance/models"
)

type TransactionRepository interface {
	CreateTransaction(t *models.Transaction) error
	UpdateTransaction(t *models.Transaction) error
	DeleteTransaction(t *models.Transaction) error
	GetUserAccountTransaction(userUUID, accountUUID, transactionUUID uuid.UUID) (*models.Transaction, error)
	GetUserAccountTransactionsByTimeRange(userUUID, accountUUID uuid.UUID, from, to time.Time) ([]*models.Transaction, error)
	GetUserAccountTransactionsByTimeRangeCategories(userUUID, accountUUID uuid.UUID, from, to time.Time, categories []string) ([]*models.Transaction, error)
	GetUserAccountStatsByTimeRange(userUUID, accountUUID uuid.UUID, from, to time.Time) (*models.StatsTimeRange, error)
	GetUserAccountStatsByTimeRangeCategories(userUUID, accountUUID uuid.UUID, from, to time.Time, categories []string) (*models.StatsTimeRangeCategories, error)
	CountUserAccountCategorySumsByTimeRange(userUUID, accountUUID uuid.UUID, from, to time.Time) ([]*models.CategorySum, error)
}
