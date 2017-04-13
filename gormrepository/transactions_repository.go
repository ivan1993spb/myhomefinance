package gormrepository

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/ivan1993spb/myhomefinance/models"
	"github.com/ivan1993spb/myhomefinance/repository"
)

type TransactionsRepository struct {
	db *gorm.DB
}

func NewTransactionsRepository(db *gorm.DB) (repository.TransactionsRepository, error) {
	r := &TransactionsRepository{db: db}

	if err := r.init(); err != nil {
		return nil, fmt.Errorf("cannot create transaction repository: %s", err)
	}

	return r, nil
}

func (r *TransactionsRepository) init() error {
	if err := r.db.AutoMigrate(&transaction{}).Error; err != nil {
		return fmt.Errorf("cannot initialize table: %s", err)
	}

	return nil
}

func (r *TransactionsRepository) GetTransactionsByTimeRange(from time.Time, to time.Time) ([]*models.Transaction, error) {
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

func (r *TransactionsRepository) CreateTransaction(transaction *models.Transaction) error {
	newTransaction := &transaction{
		Time:     transaction.Time,
		Amount:   transaction.Amount,
		Title:    transaction.Title,
		Category: transaction.Category,
	}

	if err := r.db.Create(newTransaction).Error; err != nil {
		return fmt.Errorf("cannot create transaction: %s", err)
	}

	transaction.ID = transaction.ID

	return nil
}

func (r *TransactionsRepository) UpdateTransaction(transaction *models.Transaction) error {
	if err := r.db.Save(&transaction{
		ID:       transaction.ID,
		Time:     transaction.Time,
		Amount:   transaction.Amount,
		Title:    transaction.Title,
		Category: transaction.Category,
	}).Error; err != nil {
		return fmt.Errorf("cannot update transaction: %s", err)
	}

	return nil
}
