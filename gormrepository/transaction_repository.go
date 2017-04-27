package gormrepository

import (
	"fmt"
	"sync"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/ivan1993spb/myhomefinance/models"
	"github.com/ivan1993spb/myhomefinance/repository"
)

type transactionsRepository struct {
	db   *gorm.DB
	pool *sync.Pool
}

func NewTransactionsRepository(db *gorm.DB) (repository.TransactionRepository, error) {
	return newTransactionsRepository(db)
}

func newTransactionsRepository(db *gorm.DB) (*transactionsRepository, error) {
	r := &transactionsRepository{
		db: db,
		pool: &sync.Pool{
			New: func() interface{} {
				return new(transaction)
			},
		},
	}

	if err := r.init(); err != nil {
		return nil, fmt.Errorf("cannot create transaction repository: %s", err)
	}

	return r, nil
}

func (r *transactionsRepository) init() error {
	if err := r.db.AutoMigrate(&transaction{}).Error; err != nil {
		return fmt.Errorf("cannot initialize table: %s", err)
	}

	return nil
}

func (r *transactionsRepository) CreateTransaction(t *models.Transaction) error {
	if t == nil {
		// todo return error
		return nil
	}

	if t.UUID != 0 {
		// todo return error
		return nil
	}

	newTransaction := r.pool.Get().(*transaction)
	newTransaction.ID = t.UUID
	newTransaction.AccountID = t.AccountUUID
	newTransaction.Time = t.Time
	newTransaction.Amount = t.Amount
	newTransaction.Title = t.Title
	newTransaction.Category = t.Category
	defer r.pool.Put(newTransaction)

	if err := r.db.Create(newTransaction).Error; err != nil {
		return fmt.Errorf("cannot create transaction: %s", err)
	}

	t.UUID = newTransaction.ID

	return nil
}

func (r *transactionsRepository) UpdateTransaction(t *models.Transaction) error {
	if t == nil {
		// todo return error
		return nil
	}

	if t.UUID == 0 {
		// todo return error
		return nil
	}

	updatedTransaction := r.pool.Get().(*transaction)
	updatedTransaction.ID = t.UUID
	// ignore AccountUUID
	updatedTransaction.Time = t.Time
	updatedTransaction.Amount = t.Amount
	updatedTransaction.Title = t.Title
	updatedTransaction.Category = t.Category
	defer r.pool.Put(updatedTransaction)

	if err := r.db.Model(updatedTransaction).Update("time", "amount", "title", "category").Error; err != nil {
		return fmt.Errorf("cannot update transaction: %s", err)
	}

	return nil
}

func (r transactionsRepository) DeleteTransaction(t *models.Transaction) error {
	if t == nil {
		// todo return error
		return nil
	}

	if t.UUID == 0 {
		// todo return error
		return nil
	}

	deleteTransaction := r.pool.Get().(*transaction)
	deleteTransaction.ID = t.UUID
	defer r.pool.Put(deleteTransaction)

	if err := r.db.Save(deleteTransaction).Error; err != nil {
		return fmt.Errorf("cannot delete transaction: %s", err)
	}

	return nil
}

func (r *transactionsRepository) GetAccountTransactionsByTimeRange(accountID uint64, from, to time.Time) ([]*models.Transaction, error) {
	if !from.Before(to) {
		return []*models.Transaction{}, nil
	}

	transactions := []*transaction{}

	if err := r.db.Where("account_id = ? AND time BETWEEN ? AND ?", accountID, from, to).Find(&transactions).Error; err != nil {
		return []*models.Transaction{}, fmt.Errorf("cannot get transactions by time range: %s", err)
	}

	out := make([]*models.Transaction, len(transactions))

	for i := range transactions {
		out[i] = &models.Transaction{
			UUID:        transactions[i].ID,
			AccountUUID: transactions[i].AccountID,
			Time:        transactions[i].Time,
			Amount:      transactions[i].Amount,
			Title:       transactions[i].Title,
			Category:    transactions[i].Category,
		}
	}

	return out, nil
}

func (r *transactionsRepository) GetAccountTransactionsByTimeRangeCategories(accountID uint64, from, to time.Time, categories []string) ([]*models.Transaction, error) {
	if !from.Before(to) {
		return []*models.Transaction{}, nil
	}

	transactions := []*transaction{}

	if err := r.db.Where("account_id = ? AND time BETWEEN ? AND ? AND category IN (?)", accountID, from, to, categories).Find(&transactions).Error; err != nil {
		return []*models.Transaction{}, fmt.Errorf("cannot get transactions by time range and categories: %s", err)
	}

	out := make([]*models.Transaction, len(transactions))

	for i := range transactions {
		out[i] = &models.Transaction{
			UUID:        transactions[i].ID,
			AccountUUID: transactions[i].AccountID,
			Time:        transactions[i].Time,
			Amount:      transactions[i].Amount,
			Title:       transactions[i].Title,
			Category:    transactions[i].Category,
		}
	}

	return out, nil
}

func (r *transactionsRepository) GetAccountStatsByTimeRange(accountID uint64, from, to time.Time) *models.StatsTimeRange {
	if !from.Before(to) {
		return &models.StatsTimeRange{
			AccountID: accountID,
			From:      from,
			To:        to,
		}
	}

	stats := &models.StatsTimeRange{
		AccountID: accountID,
		From:      from,
		To:        to,
	}

	err := r.db.Model(&transaction{}).Where("account_id = ? AND time BETWEEN ? AND ?", accountID, from, to).Count(&stats.Count).Error
	if err != nil {
		return &models.StatsTimeRange{
			AccountID: accountID,
			From:      from,
			To:        to,
		}
	}

	return stats
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

	return &models.StatsTimeRangeCategories{
		AccountID:  accountID,
		From:       from,
		To:         to,
		Categories: categories,
	}
}

func (r *transactionsRepository) CountAccountCategoriesSumsByTimeRange(accountID uint64, from, to time.Time) ([]*models.CategorySum, error) {
	if !from.Before(to) {
		return []*models.CategorySum{}, nil
	}

	return []*models.CategorySum{}, nil
}
