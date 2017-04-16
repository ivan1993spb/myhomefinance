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

func (c *Core) UpdateTransaction(t *models.Transaction) error {
	return c.transactionsRepository.UpdateTransaction(t)
}

func (c *Core) DeleteTransaction(t *models.Transaction) error {
	return c.transactionsRepository.DeleteTransaction(t)
}

func (c *Core) GetTransactionsByTimeRange(from time.Time, to time.Time) ([]*models.Transaction, error) {
	return c.transactionsRepository.GetTransactionsByTimeRange(from, to)
}

func (c *Core) GetTransactionsByTimeRangeCategories(from time.Time, to time.Time, categories []string) ([]*models.Transaction, error) {
	return c.transactionsRepository.GetTransactionsByTimeRangeCategories(from, to, categories)
}

func (c *Core) GetStats() (float64, float64, float64) {
	transactions, err := c.GetTransactionsByTimeRange(time.Unix(0, 0), time.Now())
	if err != nil {
		return 0, 0, 0
	}
	var inflow, outlow, balance float64
	for _, t := range transactions {
		if t.Amount > 0 {
			inflow += t.Amount
		} else if t.Amount < 0 {
			outlow += t.Amount
		}

		balance += t.Amount
	}

	if outlow < 0 {
		outlow *= -1
	}

	return inflow, outlow, balance
}
