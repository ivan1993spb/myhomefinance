package memoryrepository

import (
	"sync/atomic"

	"github.com/ivan1993spb/myhomefinance/models"
	"github.com/ivan1993spb/myhomefinance/repository"
)

type accountRepository struct {
	accounts []*models.Account
	cursorID uint64
}

func NewAccountRepository() (repository.AccountRepository, error) {
	return newAccountRepository()
}

func newAccountRepository() (*accountRepository, error) {
	return &accountRepository{}, nil
}

func (r *accountRepository) CreateAccount(a *models.Account) error {
	if a == nil {
		// todo return error
		return nil
	}

	if a.ID != 0 {
		// todo return error
		return nil
	}

	a.ID = atomic.AddUint64(&r.cursorID, 1)

	return nil
}

func (r *accountRepository) UpdateAccount(a *models.Account) error {
	if a == nil {
		// todo return error
		return nil
	}

	if a.ID == 0 {
		// todo return error
		return nil
	}

	return nil
}

func (r *accountRepository) DeleteAccount(a *models.Account) error {
	if a == nil {
		// todo return error
		return nil
	}

	if a.ID == 0 {
		// todo return error
		return nil
	}

	return nil
}
