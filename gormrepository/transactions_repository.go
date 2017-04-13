package gormrepository

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/ivan1993spb/myhomefinance/models"
)

type TransactionsRepository struct {
	db *gorm.DB
}

func NewTransactionsRepository(db *gorm.DB) (*TransactionsRepository, error) {
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
	return []*models.Transaction{}, nil
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
