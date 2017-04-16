package core

import (
	"time"

	"fmt"
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
	var inflow, outflow, balance float64
	for _, t := range transactions {
		if t.Amount > 0 {
			inflow += t.Amount
		} else if t.Amount < 0 {
			outflow += t.Amount
		}

		balance += t.Amount
	}

	if outflow < 0 {
		outflow *= -1
	}

	return inflow, outflow, balance
}

func (c *Core) GetStatsMonth() (float64, float64, float64) {
	now := time.Now()
	monthStart := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.Local)
	fmt.Println(monthStart)
	transactions, err := c.GetTransactionsByTimeRange(monthStart, now)
	if err != nil {
		return 0, 0, 0
	}

	var inflow, outflow, profit float64
	for _, t := range transactions {
		if t.Amount > 0 {
			inflow += t.Amount
		} else if t.Amount < 0 {
			outflow += t.Amount
		}

		profit += t.Amount
	}

	if outflow < 0 {
		outflow *= -1
	}

	return inflow, outflow, profit
}
