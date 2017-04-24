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
			// todo ignore account id
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

func (r *transactionsRepository) GetAccountTransactionsByTimeRange(accountID uint64, from, to time.Time) ([]*models.Transaction, error) {
	if !from.Before(to) {
		return []*models.Transaction{}, nil
	}

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

func (r *transactionsRepository) GetAccountTransactionsByTimeRangeCategories(accountID uint64, from, to time.Time, categories []string) ([]*models.Transaction, error) {
	if !from.Before(to) {
		return []*models.Transaction{}, nil
	}

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

func (r *transactionsRepository) GetAccountStatsByTimeRange(accountID uint64, from, to time.Time) *models.StatsTimeRange {
	if !from.Before(to) {
		return &models.StatsTimeRange{
			AccountID: accountID,
			From:      from,
			To:        to,
		}
	}

	r.mutex.RLock()
	defer r.mutex.RUnlock()

	if len(r.transactions) == 0 {
		return &models.StatsTimeRange{
			AccountID: accountID,
			From:      from,
			To:        to,
		}
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

	return &models.StatsTimeRange{
		AccountID: accountID,
		From:      from,
		To:        to,
		Inflow:    inflow,
		Outflow:   outflow,
		Profit:    profit,
		Count:     count,
	}
}

func (r *transactionsRepository) GetAccountStatsByTimeRangeCategories(accountID uint64, from, to time.Time, categories []string) *models.StatsTimeRangeCategories {
	if !from.Before(to) {
		return &models.StatsTimeRangeCategories{
			AccountID:  accountID,
			From:       from,
			To:         to,
			Categories: categories,
		}
	}

	r.mutex.RLock()
	defer r.mutex.RUnlock()

	if len(r.transactions) == 0 {
		return &models.StatsTimeRangeCategories{
			AccountID:  accountID,
			From:       from,
			To:         to,
			Categories: categories,
		}
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

	return &models.StatsTimeRangeCategories{
		AccountID:  accountID,
		From:       from,
		To:         to,
		Inflow:     inflow,
		Outflow:    outflow,
		Profit:     profit,
		Count:      count,
		Categories: categories,
	}
}

func (r *transactionsRepository) CountAccountCategoriesSumsByTimeRange(accountID uint64, from, to time.Time) ([]*models.CategorySum, error) {
	if !from.Before(to) {
		return []*models.CategorySum{}, nil
	}

	r.mutex.RLock()
	defer r.mutex.RUnlock()

	categorySums := make([]*models.CategorySum, 0)

	for _, t := range r.transactions {
		if accountID == t.AccountID && between(from, to, t.Time) {
			categorySums = sumCategoryTransaction(categorySums, t, from, to)
		}
	}

	return categorySums, nil
}

func sumCategoryTransaction(categorySums []*models.CategorySum, t *models.Transaction, from, to time.Time) []*models.CategorySum {
	if t == nil {
		return categorySums
	}

	for _, categorySum := range categorySums {
		if categorySum.Category == t.Category {
			categorySum.Count += 1
			categorySum.Sum += t.Amount
			return categorySums
		}
	}

	return append(categorySums, &models.CategorySum{
		AccountID: t.AccountID,
		From:      from,
		To:        to,
		Category:  t.Category,
		Sum:       t.Amount,
		Count:     1,
	})
}
