package gormrepository

import (
	"time"

	"github.com/ivan1993spb/myhomefinance/models"
)

type TransactionsRepository struct {
	transactions []*models.Transaction
}

func NewTransactionsRepository() (*TransactionsRepository, error) {
	return &TransactionsRepository{}, nil
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

func (r *TransactionsRepository) UpdateTransaction(transaction *models.Transaction) error {
	return nil
}
