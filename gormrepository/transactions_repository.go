package gormrepository

import (
	"time"

	"github.com/jinzhu/gorm"

	"github.com/ivan1993spb/myhomefinance/models"
)

type TransactionsRepository struct {
	db *gorm.DB
}

func NewTransactionsRepository(db *gorm.DB) (*TransactionsRepository, error) {
	return nil, nil
}

func (r *TransactionsRepository) init() error {
	return nil
}

func (r *TransactionsRepository) GetTransactionsByTimeRange(from time.Time, to time.Time) ([]*models.Transaction, error) {
	return []*models.Transaction{}, nil
}

func (r *TransactionsRepository) CreateTransaction(transaction *models.Transaction) error {
	return nil
}
