package memoryrepository

import "github.com/ivan1993spb/myhomefinance/models"

type AccountRepository struct {
	accounts []*models.Account
	cursorID uint64
}

func (r *AccountRepository) CreateTransaction(a *models.Account) error {
	if a == nil {
		// todo return error
		return nil
	}

	if a.ID != 0 {
		// todo return error
		return nil
	}

	return nil
}

func (r *AccountRepository) UpdateTransaction(a *models.Account) error {
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

func (r *AccountRepository) DeleteTransaction(a *models.Account) error {
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
