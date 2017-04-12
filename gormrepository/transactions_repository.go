package gormrepository

import (
	"time"

	"github.com/jinzhu/gorm"

	"github.com/ivan1993spb/myhomefinance/models"
)

type TransactionsRepository struct {
	db *gorm.DB
}

func (r *TransactionsRepository) GetTransactionsByTimeRange(from time.Time, to time.Time) ([]*models.Transaction, error) {
	return []*models.Transaction{}, nil
}
