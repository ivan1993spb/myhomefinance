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

func NewTransactionsRepository(db *gorm.DB) (repository.TransactionsRepository, error) {
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

	newTransaction := r.pool.Get().(*transaction)
	newTransaction.ID = t.ID
	newTransaction.Time = t.Time
	newTransaction.Amount = t.Amount
	newTransaction.Title = t.Title
	newTransaction.Category = t.Category
	defer r.pool.Put(newTransaction)

	if err := r.db.Create(newTransaction).Error; err != nil {
		return fmt.Errorf("cannot create transaction: %s", err)
	}

	t.ID = t.ID

	return nil
}

func (r *transactionsRepository) UpdateTransaction(t *models.Transaction) error {
	if t == nil {
		// todo return error
		return nil
	}

	updatedTransaction := r.pool.Get().(*transaction)
	updatedTransaction.ID = t.ID
	updatedTransaction.Time = t.Time
	updatedTransaction.Amount = t.Amount
	updatedTransaction.Title = t.Title
	updatedTransaction.Category = t.Category
	defer r.pool.Put(updatedTransaction)

	if err := r.db.Save(updatedTransaction).Error; err != nil {
		return fmt.Errorf("cannot update transaction: %s", err)
	}

	return nil
}

func (r transactionsRepository) DeleteTransaction(t *models.Transaction) error {
	if t == nil {
		// todo return error
		return nil
	}

	return nil
}

func (r *transactionsRepository) GetTransactionsByTimeRange(from time.Time, to time.Time) ([]*models.Transaction, error) {
	transactions := []*transaction{}

	if err := r.db.Where("time BETWEEN ? AND ?", from, to).Find(&transactions).Error; err != nil {
		return []*models.Transaction{}, fmt.Errorf("cannot get transactions by time range: %s", err)
	}

	out := make([]*models.Transaction, len(transactions))

	for i := range transactions {
		out[i] = &models.Transaction{
			ID:       transactions[i].ID,
			Time:     transactions[i].Time,
			Amount:   transactions[i].Amount,
			Title:    transactions[i].Title,
			Category: transactions[i].Category,
		}
	}

	return out, nil
}

func (r *transactionsRepository) GetTransactionsByTimeRangeCategories(from time.Time, to time.Time, categories []string) ([]*models.Transaction, error) {
	return nil, nil
}

func (r *transactionsRepository) GetStatsByTimeRange(from time.Time, to time.Time) (float64, float64, float64, uint64) {
	return 0, 0, 0, 0
}

func (r *transactionsRepository) GetStatsByTimeRangeCategories(from time.Time, to time.Time, categories []string) (float64, float64, float64, uint64) {
	return 0, 0, 0, 0
}
