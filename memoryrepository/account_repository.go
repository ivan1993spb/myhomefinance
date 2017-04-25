package memoryrepository

import (
	"sync"

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

type errCreateAccount string

func (e errCreateAccount) Error() string {
	return "cannot create account: " + string(e)
}

func (r *accountRepository) CreateAccount(a *models.Account) error {
	if a == nil {
		return errCreateAccount("passed nil account")
	}

	if a.ID != 0 {
		return errCreateAccount("passed account with non zero identifier")
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

type errUpdateAccount string

func (e errUpdateAccount) Error() string {
	return "cannot update account: " + string(e)
}

func (r *accountRepository) UpdateAccount(a *models.Account) error {
	if a == nil {
		return errCreateAccount("passed nil account")
	}

	if a.ID == 0 {
		return errCreateAccount("passed account with zero identifier")
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

type errDeleteAccount string

func (e errDeleteAccount) Error() string {
	return "cannot delete account: " + string(e)
}

func (r *accountRepository) DeleteAccount(a *models.Account) error {
	if a == nil {
		return errDeleteAccount("passed nil account")
	}

	if a.ID == 0 {
		return errDeleteAccount("passed account with zero identifier")
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

type errGetAccountsByUserID string

func (e errGetAccountsByUserID) Error() string {
	return "cannot get accounts by user identifier: " + string(e)
}

func (r *accountRepository) GetAccountsByUserID(userID uint64) ([]*models.Account, error) {
	if userID == 0 {
		return nil, errGetAccountsByUserID("passed zero user identier")
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
