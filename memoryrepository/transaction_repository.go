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

	transactions := make([]*models.Transaction, 0)
	for _, t := range r.transactions {
		if between(from, to, t.Time) {
			var transaction models.Transaction = *t
			transactions = append(transactions, &transaction)
		}
	}

	return transactions, nil
}

func between(from, to, t time.Time) bool {
	return t.Equal(from) || t.Equal(to) || t.After(from) && t.Before(to)
}

func (r *TransactionsRepository) GetTransactionsByTimeRangeCategories(from time.Time, to time.Time, categories []string) ([]*models.Transaction, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	transactions := make([]*models.Transaction, 0)
	for _, t := range r.transactions {
		if contains(t.Category, categories) && between(from, to, t.Time) {
			var transaction models.Transaction = *t
			transactions = append(transactions, &transaction)
		}
	}

	return transactions, nil
}

func contains(str string, slice []string) bool {
	for i := range slice {
		if slice[i] == str {
			return true
		}
	}
	return false
}

func (r *TransactionsRepository) CreateTransaction(t *models.Transaction) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	var transaction models.Transaction = *t
	r.transactions = append(r.transactions, &transaction)

	return nil
}

func (r *TransactionsRepository) UpdateTransaction(t *models.Transaction) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	for i := range r.transactions {
		if r.transactions[i].ID == t.ID {
			*r.transactions[i] = *t
			break
		}
	}

	return nil
}
