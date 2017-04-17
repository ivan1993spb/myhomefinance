package memoryrepository

import (
	"sync"
	"time"

	"github.com/ivan1993spb/myhomefinance/models"
	"github.com/ivan1993spb/myhomefinance/repository"
)

type transactionsRepository struct {
	transactions []*models.Transaction
	cursorID     uint64
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
	if t == nil {
		// todo return error
		return nil
	}

	if t.ID != 0 {
		// todo return error
		return nil
	}

	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.cursorID++
	t.ID = r.cursorID

	transaction := r.pool.Get().(*models.Transaction)
	*transaction = *t
	r.transactions = append(r.transactions, transaction)

	return nil
}

func (r *transactionsRepository) UpdateTransaction(t *models.Transaction) error {
	if t == nil {
		// todo return error
		return nil
	}

	if t.ID == 0 {
		// todo return error
		return nil
	}

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
	if t == nil {
		// todo return error
		return nil
	}

	if t.ID == 0 {
		// todo return error
		return nil
	}

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

func (r *transactionsRepository) GetAccountTransactionsByTimeRange(accountID uint64, from time.Time, to time.Time) ([]*models.Transaction, error) {

	r.mutex.RLock()
	defer r.mutex.RUnlock()

	transactions := make([]*models.Transaction, 0)
	for _, t := range r.transactions {
		if accountID == t.AccountID && between(from, to, t.Time) {
			var transaction models.Transaction = *t
			transactions = append(transactions, &transaction)
		}
	}

	return transactions, nil
}

func between(from, to, t time.Time) bool {
	return t.Equal(from) || t.Equal(to) || t.After(from) && t.Before(to)
}

func (r *transactionsRepository) GetAccountTransactionsByTimeRangeCategories(accountID uint64, from time.Time, to time.Time, categories []string) ([]*models.Transaction, error) {

	r.mutex.RLock()
	defer r.mutex.RUnlock()

	transactions := make([]*models.Transaction, 0)
	for _, t := range r.transactions {
		if accountID == t.AccountID && contains(t.Category, categories) && between(from, to, t.Time) {
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

func (r *transactionsRepository) GetAccountStatsByTimeRange(accountID uint64, from time.Time, to time.Time) (float64, float64, float64, uint64) {

	r.mutex.RLock()
	defer r.mutex.RUnlock()

	if len(r.transactions) == 0 {
		return 0, 0, 0, 0
	}

	var inflow, outflow, profit float64
	var count uint64

	for _, t := range r.transactions {
		if accountID == t.AccountID && between(from, to, t.Time) {
			count += 1

			if t.Amount > 0 {
				inflow += t.Amount
			} else if t.Amount < 0 {
				outflow += t.Amount
			}

			profit += t.Amount
		}
	}

	if outflow < 0 {
		outflow *= -1
	}

	return inflow, outflow, profit, count
}

func (r *transactionsRepository) GetAccountStatsByTimeRangeCategories(accountID uint64, from time.Time, to time.Time, categories []string) (float64, float64, float64, uint64) {

	r.mutex.RLock()
	defer r.mutex.RUnlock()

	if len(r.transactions) == 0 {
		return 0, 0, 0, 0
	}

	var inflow, outflow, profit float64
	var count uint64

	for _, t := range r.transactions {
		if accountID == t.AccountID && contains(t.Category, categories) && between(from, to, t.Time) {
			count += 1

			if t.Amount > 0 {
				inflow += t.Amount
			} else if t.Amount < 0 {
				outflow += t.Amount
			}

			profit += t.Amount
		}
	}

	if outflow < 0 {
		outflow *= -1
	}

	return inflow, outflow, profit, count
}
