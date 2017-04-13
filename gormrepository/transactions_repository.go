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

func (r *TransactionsRepository) CreateTransaction(t *models.Transaction) error {
	newTransaction := &transaction{
		Time:     t.Time,
		Amount:   t.Amount,
		Title:    t.Title,
		Category: t.Category,
	}

	if err := r.db.Create(newTransaction).Error; err != nil {
		return fmt.Errorf("cannot create transaction: %s", err)
	}

	t.ID = t.ID

	return nil
}
