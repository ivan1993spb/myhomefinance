package core

import (
	"time"

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

func (c *Core) UpdateTransaction(t *models.Transaction) error {
	return c.transactionsRepository.UpdateTransaction(t)
}

func (c *Core) DeleteTransaction(ID uint64) error {
	return c.transactionsRepository.DeleteTransaction(&models.Transaction{
		ID: ID,
	})
}

func (c *Core) GetAccountTransactionsByTimeRange(accountID uint64, from time.Time, to time.Time) ([]*models.Transaction, error) {
	return c.transactionsRepository.GetAccountTransactionsByTimeRange(accountID, from, to)
}

func (c *Core) GetAccountTransactionsByTimeRangeCategories(accountID uint64, from time.Time, to time.Time, categories []string) ([]*models.Transaction, error) {
	return c.transactionsRepository.GetAccountTransactionsByTimeRangeCategories(accountID, from, to, categories)
}

func (c *Core) GetAccountStatsByTimeRange(accountID uint64, from time.Time, to time.Time) (float64, float64, float64, uint64) {
	return c.transactionsRepository.GetAccountStatsByTimeRange(accountID, from, to)
}

func (c *Core) CountAccountCategoriesSumsByTimeRange(accountID uint64, from time.Time, to time.Time) ([]*models.CategorySum, error) {
	return c.transactionsRepository.CountAccountCategoriesSumsByTimeRange(accountID, from, to)
}

func (c *Core) CreateAccount() (*models.Account, error) {
	a := &models.Account{}
	if err := c.accountRepository.CreateAccount(a); err != nil {
		return nil, err
	}
	return a, nil
}

func (c *Core) UpdateAccount(a *models.Account) error {
	return c.accountRepository.UpdateAccount(a)
}

func (c *Core) DeleteAccount(ID uint64) error {
	return c.accountRepository.DeleteAccount(&models.Account{
		ID: ID,
	})
}
