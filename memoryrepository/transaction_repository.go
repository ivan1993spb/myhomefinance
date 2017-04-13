package gormrepository

import (
	"sync"
	"time"

	"github.com/ivan1993spb/myhomefinance/models"
	"github.com/ivan1993spb/myhomefinance/repository"
)

type TransactionsRepository struct {
	transactions []*models.Transaction
	mutex        *sync.RWMutex
}

func NewTransactionsRepository() (repository.TransactionsRepository, error) {
	return &TransactionsRepository{
		transactions: []*models.Transaction{},
		mutex:        &sync.RWMutex{},
	}, nil
}

func (r *TransactionsRepository) GetTransactionsByTimeRange(from time.Time, to time.Time) ([]*models.Transaction, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	transactions := make([]*models.Transaction, 0, len(r.transactions))
	for _, transaction := range r.transactions {
		if transaction.Time.Equal(from) || transaction.Time.Equal(to) || transaction.Time.After(from) && transaction.Time.Before(to) {
			var t *models.Transaction
			*t = *transaction
			transactions = append(transactions, t)
		}
	}

	return transactions, nil
}

func (r *TransactionsRepository) CreateTransaction(transaction *models.Transaction) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.transactions = append(r.transactions, transaction)

	return nil
}

func (r *TransactionsRepository) UpdateTransaction(transaction *models.Transaction) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	for i := range r.transactions {
		if r.transactions[i].ID == transaction.ID {
			r.transactions[i] = transaction
			break
		}
	}

	return nil
}
