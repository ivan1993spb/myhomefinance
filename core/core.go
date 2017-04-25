package core

import (
	"time"

	"github.com/ivan1993spb/myhomefinance/iso4217"
	"github.com/ivan1993spb/myhomefinance/models"
	"github.com/ivan1993spb/myhomefinance/repository"
)

type Core struct {
	transactionsRepository repository.TransactionsRepository
	accountRepository      repository.AccountRepository
}

func New(transactionsRepository repository.TransactionsRepository, accountRepository repository.AccountRepository) *Core {
	return &Core{
		transactionsRepository: transactionsRepository,
		accountRepository:      accountRepository,
	}
}

func (c *Core) CreateTransaction(accountID uint64, t time.Time, amount float64, title, category string) (*models.Transaction, error) {
	tr := &models.Transaction{
		AccountID: accountID,
		Time:      t,
		Amount:    amount,
		Title:     title,
		Category:  category,
	}
	if err := c.transactionsRepository.CreateTransaction(tr); err != nil {
		return nil, err
	}
	return tr, nil
}

func (c *Core) UpdateTransaction(ID uint64, t time.Time, amount float64, title, category string) (*models.Transaction, error) {
	tr := &models.Transaction{
		ID: ID,
		// ignore AccountID
		Time:     t,
		Amount:   amount,
		Title:    title,
		Category: category,
	}
	if err := c.transactionsRepository.UpdateTransaction(tr); err != nil {
		return nil, err
	}
	return tr, nil
}

func (c *Core) DeleteTransaction(ID uint64) error {
	return c.transactionsRepository.DeleteTransaction(&models.Transaction{
		ID: ID,
	})
}

func (c *Core) GetAccountTransactionsByTimeRange(accountID uint64, from, to time.Time) ([]*models.Transaction, error) {
	return c.transactionsRepository.GetAccountTransactionsByTimeRange(accountID, from, to)
}

func (c *Core) GetAccountTransactionsByTimeRangeCategories(accountID uint64, from, to time.Time, categories []string) ([]*models.Transaction, error) {
	return c.transactionsRepository.GetAccountTransactionsByTimeRangeCategories(accountID, from, to, categories)
}

func (c *Core) GetAccountStatsByTimeRange(accountID uint64, from, to time.Time) (*models.StatsTimeRange, error) {
	return c.transactionsRepository.GetAccountStatsByTimeRange(accountID, from, to)
}

func (c *Core) GetAccountStatsByTimeRangeCategories(accountID uint64, from, to time.Time, categories []string) (*models.StatsTimeRangeCategories, error) {
	return c.transactionsRepository.GetAccountStatsByTimeRangeCategories(accountID, from, to, categories)
}

func (c *Core) CountAccountCategoriesSumsByTimeRange(accountID uint64, from, to time.Time) ([]*models.CategorySum, error) {
	return c.transactionsRepository.CountAccountCategoriesSumsByTimeRange(accountID, from, to)
}

func (c *Core) CreateAccount(userID uint64, name string, currency iso4217.Currency) (*models.Account, error) {
	a := &models.Account{
		UserID:   userID,
		Name:     name,
		Currency: currency,
	}
	if err := c.accountRepository.CreateAccount(a); err != nil {
		return nil, err
	}
	return a, nil
}

func (c *Core) UpdateAccount(ID uint64, name string, currency iso4217.Currency) (*models.Account, error) {
	a := &models.Account{
		ID: ID,
		// ignore UserID
		Name:     name,
		Currency: currency,
	}
	if err := c.accountRepository.UpdateAccount(a); err != nil {
		return nil, err
	}
	return a, nil
}

func (c *Core) DeleteAccount(ID uint64) error {
	return c.accountRepository.DeleteAccount(&models.Account{
		ID: ID,
	})
}
