package core

import (
	"time"

	"github.com/ivan1993spb/myhomefinance/models"
	"github.com/ivan1993spb/myhomefinance/repository"
)

type Core struct {
	transactionsRepository repository.TransactionsRepository
}

func New(transactionsRepository repository.TransactionsRepository) *Core {
	return &Core{transactionsRepository: transactionsRepository}
}

func (c *Core) CreateTransaction(t *models.Transaction) error {
	return c.transactionsRepository.CreateTransaction(t)
}

func (c *Core) GetTransactionsByTimeRange(from time.Time, to time.Time) ([]*models.Transaction, error) {
	return c.transactionsRepository.GetTransactionsByTimeRange(from, to)
}

func (c *Core) UpdateTransaction(t *models.Transaction) error {
	return c.transactionsRepository.UpdateTransaction(t)
}
