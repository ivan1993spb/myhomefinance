package memoryrepository

import (
	"sync"

	"github.com/satori/go.uuid"

	"github.com/ivan1993spb/myhomefinance/models"
	"github.com/ivan1993spb/myhomefinance/repository"
)

type accountRepository struct {
	accounts []*models.Account
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
				return &models.Account{}
			},
		},
	}, nil
}

type errCreateAccount string

func (e errCreateAccount) Error() string {
	return "cannot create account: " + string(e)
}

func (r *accountRepository) CreateAccount(a *models.Account) error {
	if a == nil {
		return errCreateAccount("passed nil account")
	}

	r.mutex.Lock()
	defer r.mutex.Unlock()

	for _, account := range r.accounts {
		if account.UUID == a.UUID {
			return errCreateAccount("uuid of passed account is already used")
		}
	}

	account := r.pool.Get().(*models.Account)
	*account = *a
	r.accounts = append(r.accounts, account)

	return nil
}

type errUpdateAccount string

func (e errUpdateAccount) Error() string {
	return "cannot update account: " + string(e)
}

func (r *accountRepository) UpdateAccount(a *models.Account) error {
	if a == nil {
		return errCreateAccount("passed nil account")
	}

	r.mutex.Lock()
	defer r.mutex.Unlock()

	for i := range r.accounts {
		if r.accounts[i].UUID == a.UUID {
			// ignore user uuid
			r.accounts[i].Name = a.Name
			r.accounts[i].Currency = a.Currency
			*a = *r.accounts[i]
			return nil
		}
	}

	return errCreateAccount("not found")
}

type errDeleteAccount string

func (e errDeleteAccount) Error() string {
	return "cannot delete account: " + string(e)
}

func (r *accountRepository) DeleteAccount(a *models.Account) error {
	if a == nil {
		return errDeleteAccount("passed nil account")
	}

	r.mutex.Lock()
	defer r.mutex.Unlock()

	for i := range r.accounts {
		if r.accounts[i].UUID == a.UUID {
			r.pool.Put(r.accounts[i])
			r.accounts = append(r.accounts[:i], r.accounts[i+1:]...)
			return nil
		}
	}

	return errDeleteAccount("not found")
}

type errGetUserAccounts string

func (e errGetUserAccounts) Error() string {
	return "cannot get user accounts: " + string(e)
}

func (r *accountRepository) GetUserAccounts(userUUID uuid.UUID) ([]*models.Account, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	accounts := make([]*models.Account, 0)
	for i := range r.accounts {
		if r.accounts[i].UserUUID == userUUID {
			accounts = append(accounts, &(*r.accounts[i]))
		}
	}

	return accounts, nil
}
