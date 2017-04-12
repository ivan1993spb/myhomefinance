package repository

import (
	"time"

	"github.com/ivan1993spb/myhomefinance/models"
)

type TransactionsRepository interface {
	GetTransactionsByTimeRange(from time.Time, to time.Time) ([]*models.Transaction, error)
}
