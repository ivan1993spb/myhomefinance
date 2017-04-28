package core

import (
	"time"

	"github.com/satori/go.uuid"

	"github.com/ivan1993spb/myhomefinance/iso4217"
	"github.com/ivan1993spb/myhomefinance/models"
	"github.com/ivan1993spb/myhomefinance/repository"
)

type Core struct {
	userRepository        repository.UserRepository
	accountRepository     repository.AccountRepository
	transactionRepository repository.TransactionRepository
}

func New(userRepository repository.UserRepository, accountRepository repository.AccountRepository, transactionRepository repository.TransactionRepository) *Core {
	return &Core{
		userRepository:        userRepository,
		accountRepository:     accountRepository,
		transactionRepository: transactionRepository,
	}
}

func (c *Core) CreateTransaction(userUUID, accountUUID uuid.UUID, t time.Time, amount float64, title, category string) (*models.Transaction, error) {
	tr := &models.Transaction{
		UUID:        uuid.NewV4(),
		AccountUUID: accountUUID,
		UserUUID:    userUUID,
		Time:        t,
		Amount:      amount,
		Title:       title,
		Category:    category,
	}
	if err := c.transactionRepository.CreateTransaction(tr); err != nil {
		return nil, err
	}
	return tr, nil
}

func (c *Core) UpdateTransaction(userUUID, accountUUID, transactionUUID uuid.UUID, t time.Time, amount float64, title, category string) (*models.Transaction, error) {
	tr := &models.Transaction{
		UUID:        transactionUUID,
		AccountUUID: accountUUID,
		UserUUID:    userUUID,
		Time:        t,
		Amount:      amount,
		Title:       title,
		Category:    category,
	}
	if err := c.transactionRepository.UpdateTransaction(tr); err != nil {
		return nil, err
	}
	return tr, nil
}

func (c *Core) DeleteTransaction(userUUID, accountUUID, transactionUUID uuid.UUID) error {
	return c.transactionRepository.DeleteTransaction(&models.Transaction{
		UUID:        transactionUUID,
		AccountUUID: accountUUID,
		UserUUID:    userUUID,
	})
}

func (c *Core) GetUserAccountTransaction(userUUID, accountUUID, transactionUUID uuid.UUID) (*models.Transaction, error) {
	return c.transactionRepository.GetUserAccountTransaction(userUUID, accountUUID, transactionUUID)
}

func (c *Core) GetUserAccountTransactionsByTimeRange(userUUID, accountUUID uuid.UUID, from, to time.Time) ([]*models.Transaction, error) {
	return c.transactionRepository.GetUserAccountTransactionsByTimeRange(userUUID, accountUUID, from, to)
}

func (c *Core) GetUserAccountTransactionsByTimeRangeCategories(userUUID, accountUUID uuid.UUID, from, to time.Time, categories []string) ([]*models.Transaction, error) {
	return c.transactionRepository.GetUserAccountTransactionsByTimeRangeCategories(userUUID, accountUUID, from, to, categories)
}

func (c *Core) GetUserAccountStatsByTimeRange(userUUID, accountUUID uuid.UUID, from, to time.Time) (*models.StatsTimeRange, error) {
	return c.transactionRepository.GetUserAccountStatsByTimeRange(userUUID, accountUUID, from, to)
}

func (c *Core) GetUserAccountStatsByTimeRangeCategories(userUUID, accountUUID uuid.UUID, from, to time.Time, categories []string) (*models.StatsTimeRangeCategories, error) {
	return c.transactionRepository.GetUserAccountStatsByTimeRangeCategories(userUUID, accountUUID, from, to, categories)
}

func (c *Core) CountUserAccountCategorySumsByTimeRange(userUUID, accountUUID uuid.UUID, from, to time.Time) ([]*models.CategorySum, error) {
	return c.transactionRepository.CountUserAccountCategorySumsByTimeRange(userUUID, accountUUID, from, to)
}

func (c *Core) CreateAccount(userUUID uuid.UUID, name string, currency iso4217.Currency) (*models.Account, error) {
	a := &models.Account{
		UUID:     uuid.NewV4(),
		UserUUID: userUUID,
		Name:     name,
		Currency: currency,
	}
	if err := c.accountRepository.CreateAccount(a); err != nil {
		return nil, err
	}
	return a, nil
}

func (c *Core) UpdateAccount(userUUID, accountUUID uuid.UUID, name string, currency iso4217.Currency) (*models.Account, error) {
	a := &models.Account{
		UUID:     accountUUID,
		UserUUID: userUUID,
		Name:     name,
		Currency: currency,
	}
	if err := c.accountRepository.UpdateAccount(a); err != nil {
		return nil, err
	}
	return a, nil
}

func (c *Core) DeleteAccount(userUUID, accountUUID uuid.UUID) error {
	return c.accountRepository.DeleteAccount(&models.Account{
		UUID:     accountUUID,
		UserUUID: userUUID,
	})
}

func (c *Core) GetUserAccounts(userUUID uuid.UUID) ([]*models.Account, error) {
	return c.accountRepository.GetUserAccounts(userUUID)
}

func (c Core) CreateUser() (*models.User, error) {
	u := &models.User{
		UUID: uuid.NewV4(),
	}
	if err := c.userRepository.CreateUser(u); err != nil {
		return nil, err
	}
	return u, nil
}

func (c Core) UpdateUser(userUUID uuid.UUID) (*models.User, error) {
	// todo create fields
	u := &models.User{
		UUID: userUUID,
	}
	if err := c.userRepository.UpdateUser(u); err != nil {
		return nil, err
	}
	return u, nil
}

func (c Core) DeleteUser(userUUID uuid.UUID) error {
	return c.userRepository.DeleteUser(&models.User{
		UUID: userUUID,
	})
}
