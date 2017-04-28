package memoryrepository

import (
	"sync"
	"time"

	"github.com/satori/go.uuid"

	"github.com/ivan1993spb/myhomefinance/models"
	"github.com/ivan1993spb/myhomefinance/repository"
)

type transactionRepository struct {
	transactions []*models.Transaction
	mutex        *sync.RWMutex
	pool         *sync.Pool
}

func NewTransactionRepository() (repository.TransactionRepository, error) {
	return newTransactionRepository()
}

func newTransactionRepository() (*transactionRepository, error) {
	return &transactionRepository{
		transactions: []*models.Transaction{},
		mutex:        &sync.RWMutex{},
		pool: &sync.Pool{
			New: func() interface{} {
				return new(models.Transaction)
			},
		},
	}, nil
}

type errCreateTransaction string

func (e errCreateTransaction) Error() string {
	return "cannot create transaction: " + string(e)
}

func (r *transactionRepository) CreateTransaction(t *models.Transaction) error {
	if t == nil {
		return errCreateTransaction("passed nil transaction")
	}

	r.mutex.Lock()
	defer r.mutex.Unlock()

	for _, transaction := range r.transactions {
		if transaction.UUID == t.UUID {
			return errCreateTransaction("uuid of passed transaction is already used")
		}
	}

	transaction := r.pool.Get().(*models.Transaction)
	*transaction = *t
	r.transactions = append(r.transactions, transaction)

	return nil
}

type errUpdateTransaction string

func (e errUpdateTransaction) Error() string {
	return "cannot update transaction: " + string(e)
}

func (r *transactionRepository) UpdateTransaction(t *models.Transaction) error {
	if t == nil {
		return errUpdateTransaction("passed nil transaction")
	}

	r.mutex.Lock()
	defer r.mutex.Unlock()

	for i := range r.transactions {
		if r.transactions[i].UUID == t.UUID {
			// ignore account uuid
			// ignore user uuid
			r.transactions[i].Time = t.Time
			r.transactions[i].Amount = t.Amount
			r.transactions[i].Title = t.Title
			r.transactions[i].Category = t.Category
			*t = *r.transactions[i]
			return nil
		}
	}

	return errUpdateTransaction("not found")
}

type errDeleteTransaction string

func (e errDeleteTransaction) Error() string {
	return "cannot delete transaction: " + string(e)
}

func (r *transactionRepository) DeleteTransaction(t *models.Transaction) error {
	if t == nil {
		return errDeleteTransaction("passed nil transaction")
	}

	r.mutex.Lock()
	defer r.mutex.Unlock()

	for i := range r.transactions {
		if r.transactions[i].UUID == t.UUID {
			r.pool.Put(r.transactions[i])
			r.transactions = append(r.transactions[:i], r.transactions[i+1:]...)
			return nil
		}
	}

	return errDeleteTransaction("not found")
}

type errGetUserAccountTransaction string

func (e errGetUserAccountTransaction) Error() string {
	return "cannot get user account transaction: " + string(e)
}

func (r *transactionRepository) GetUserAccountTransaction(userUUID, accountUUID, transactionUUID uuid.UUID) (*models.Transaction, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	for _, t := range r.transactions {
		if t.UserUUID == userUUID && t.AccountUUID == accountUUID && t.UUID == transactionUUID {
			var transaction models.Transaction = *t
			return &transaction, nil
		}
	}

	return nil, errGetUserAccountTransaction("not found")
}

func (r *transactionRepository) GetUserAccountTransactionsByTimeRange(userUUID, accountUUID uuid.UUID, from, to time.Time) ([]*models.Transaction, error) {
	if !from.Before(to) {
		return []*models.Transaction{}, nil
	}

	r.mutex.RLock()
	defer r.mutex.RUnlock()

	transactions := make([]*models.Transaction, 0)
	for _, t := range r.transactions {
		if t.UserUUID == userUUID && accountUUID == t.AccountUUID && between(from, to, t.Time) {
			var transaction models.Transaction = *t
			transactions = append(transactions, &transaction)
		}
	}

	return transactions, nil
}

func between(from, to, t time.Time) bool {
	return t.Equal(from) || t.Equal(to) || t.After(from) && t.Before(to)
}

func (r *transactionRepository) GetUserAccountTransactionsByTimeRangeCategories(userUUID, accountUUID uuid.UUID, from, to time.Time, categories []string) ([]*models.Transaction, error) {
	if !from.Before(to) {
		return []*models.Transaction{}, nil
	}

	r.mutex.RLock()
	defer r.mutex.RUnlock()

	transactions := make([]*models.Transaction, 0)
	for _, t := range r.transactions {
		if t.UserUUID == userUUID && accountUUID == t.AccountUUID && contains(t.Category, categories) && between(from, to, t.Time) {
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

func (r *transactionRepository) GetUserAccountStatsByTimeRange(userUUID, accountUUID uuid.UUID, from, to time.Time) (*models.StatsTimeRange, error) {
	if !from.Before(to) {
		return &models.StatsTimeRange{
			AccountUUID: accountUUID,
			UserUUID:    userUUID,
			From:        from,
			To:          to,
		}, nil
	}

	r.mutex.RLock()
	defer r.mutex.RUnlock()

	if len(r.transactions) == 0 {
		return &models.StatsTimeRange{
			AccountUUID: accountUUID,
			UserUUID:    userUUID,
			From:        from,
			To:          to,
		}, nil
	}

	stats := &models.StatsTimeRange{
		AccountUUID: accountUUID,
		UserUUID:    userUUID,
		From:        from,
		To:          to,
	}

	for _, t := range r.transactions {
		if t.UserUUID == userUUID && accountUUID == t.AccountUUID && between(from, to, t.Time) {
			stats.Count += 1

			if t.Amount > 0 {
				stats.Inflow += t.Amount
			} else if t.Amount < 0 {
				stats.Outflow += t.Amount
			}

			stats.Profit += t.Amount
		}
	}

	if stats.Outflow < 0 {
		stats.Outflow *= -1
	}

	return stats, nil
}

func (r *transactionRepository) GetUserAccountStatsByTimeRangeCategories(userUUID, accountUUID uuid.UUID, from, to time.Time, categories []string) (*models.StatsTimeRangeCategories, error) {
	if !from.Before(to) {
		return &models.StatsTimeRangeCategories{
			AccountUUID: accountUUID,
			UserUUID:    userUUID,
			From:        from,
			To:          to,
			Categories:  categories,
		}, nil
	}

	r.mutex.RLock()
	defer r.mutex.RUnlock()

	if len(r.transactions) == 0 {
		return &models.StatsTimeRangeCategories{
			AccountUUID: accountUUID,
			UserUUID:    userUUID,
			From:        from,
			To:          to,
			Categories:  categories,
		}, nil
	}

	stats := &models.StatsTimeRangeCategories{
		AccountUUID: accountUUID,
		UserUUID:    userUUID,
		From:        from,
		To:          to,
		Categories:  categories,
	}

	for _, t := range r.transactions {
		if t.UserUUID == userUUID && accountUUID == t.AccountUUID && contains(t.Category, categories) && between(from, to, t.Time) {
			stats.Count += 1

			if t.Amount > 0 {
				stats.Inflow += t.Amount
			} else if t.Amount < 0 {
				stats.Outflow += t.Amount
			}

			stats.Profit += t.Amount
		}
	}

	if stats.Outflow < 0 {
		stats.Outflow *= -1
	}

	return stats, nil
}

func (r *transactionRepository) CountUserAccountCategorySumsByTimeRange(userUUID, accountUUID uuid.UUID, from, to time.Time) ([]*models.CategorySum, error) {
	if !from.Before(to) {
		return []*models.CategorySum{}, nil
	}

	r.mutex.RLock()
	defer r.mutex.RUnlock()

	categorySums := make([]*models.CategorySum, 0)

	for _, t := range r.transactions {
		if t.UserUUID == userUUID && accountUUID == t.AccountUUID && between(from, to, t.Time) {
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
		AccountUUID: t.AccountUUID,
		UserUUID:    t.UserUUID,
		From:        from,
		To:          to,
		Category:    t.Category,
		Sum:         t.Amount,
		Count:       1,
	})
}
