package memoryrepository

import (
	"sync"
	"time"

	"github.com/ivan1993spb/myhomefinance/models"
	"github.com/ivan1993spb/myhomefinance/repository"
)

type transactionsRepository struct {
	transactions []*models.Transaction
	mutex        *sync.RWMutex
	pool         *sync.Pool
}

func NewTransactionsRepository() (repository.TransactionsRepository, error) {
	return newTransactionsRepository()
}

func newTransactionsRepository() (*transactionsRepository, error) {
	return &transactionsRepository{
		transactions: []*models.Transaction{},
		mutex:        &sync.RWMutex{},
		pool: &sync.Pool{
			New: func() interface{} {
				return new(models.Transaction)
			},
		},
	}, nil
}

func (r *transactionsRepository) CreateTransaction(t *models.Transaction) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	transaction := r.pool.Get().(*models.Transaction)
	*transaction = *t
	r.transactions = append(r.transactions, transaction)

	return nil
}

func (r *transactionsRepository) UpdateTransaction(t *models.Transaction) error {
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

func (r *transactionsRepository) DeleteTransaction(t *models.Transaction) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	for i := range r.transactions {
		if r.transactions[i].ID == t.ID {
			r.pool.Put(r.transactions[i])
			r.transactions = append(r.transactions[:i], r.transactions[i+1:]...)
			break
		}
	}

	return nil
}

func (r *transactionsRepository) GetTransactionsByTimeRange(from time.Time, to time.Time) ([]*models.Transaction, error) {
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

func (r *transactionsRepository) GetTransactionsByTimeRangeCategories(from time.Time, to time.Time, categories []string) ([]*models.Transaction, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	transactions := make([]*models.Transaction, 0)
	for _, t := range r.transactions {
		if contains(t.Category, categories) && between(from, to, t.Time) {
			transaction := *t
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
