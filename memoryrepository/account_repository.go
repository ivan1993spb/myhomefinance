package memoryrepository

import (
	"sync"

	"github.com/ivan1993spb/myhomefinance/iso4217"
	"github.com/ivan1993spb/myhomefinance/models"
	"github.com/ivan1993spb/myhomefinance/repository"
)

type accountRepository struct {
	accounts []*models.Account
	cursorID uint64
	mutex    *sync.Mutex
	pool     *sync.Pool
}

func NewAccountRepository() (repository.AccountRepository, error) {
	return newAccountRepository()
}

func newAccountRepository() (*accountRepository, error) {
	return &accountRepository{
		accounts: []*models.Account{},
		mutex:    &sync.Mutex{},
		pool: &sync.Pool{
			New: func() interface{} {
				return new(models.Account)
			},
		},
	}, nil
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

	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.cursorID++
	a.ID = r.cursorID

	account := r.pool.Get().(*models.Account)
	*account = *a
	r.accounts = append(r.accounts, account)

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

	r.mutex.Lock()
	defer r.mutex.Unlock()

	for i := range r.accounts {
		if r.accounts[i].ID == a.ID {
			// ignore user id
			r.accounts[i].Name = a.Name
			r.accounts[i].Currency = a.Currency
			*a = *r.accounts[i]
			break
		}
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

	r.mutex.Lock()
	defer r.mutex.Unlock()

	for i := range r.accounts {
		if r.accounts[i].ID == a.ID {
			r.pool.Put(r.accounts[i])
			r.accounts = append(r.accounts[:i], r.accounts[i+1:]...)
			break
		}
	}

	return nil
}

func (r *accountRepository) GetUserAccounts(userID uint64) ([]*models.Account, error) {
	if userID == 0 {
		// todo return error
		return nil, nil
	}

	r.mutex.Lock()
	defer r.mutex.Unlock()

	accounts := make([]*models.Account, 0)
	for i := range r.accounts {
		if r.accounts[i].UserID == userID {
			accounts = append(accounts, &(*r.accounts[i]))
		}
	}

	return accounts, nil
}
