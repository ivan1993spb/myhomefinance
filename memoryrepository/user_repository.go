package memoryrepository

import (
	"sync"

	"github.com/ivan1993spb/myhomefinance/models"
	"github.com/ivan1993spb/myhomefinance/repository"
)

type userRepository struct {
	users    []*models.User
	cursorID uint64
	mutex    *sync.Mutex
	pool     *sync.Pool
}

func NewUserRepository() (repository.UserRepository, error) {
	return newUserRepository()
}

func newUserRepository() (*userRepository, error) {
	return &userRepository{
		users: []*models.User{},
		mutex: &sync.Mutex{},
		pool: &sync.Pool{
			New: func() interface{} {
				return new(models.User)
			},
		},
	}, nil
}

func (r *userRepository) CreateUser(u *models.User) error {
	if u == nil {
		// todo return error
		return nil
	}

	if u.ID != 0 {
		// todo return error
		return nil
	}

	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.cursorID++
	u.ID = r.cursorID

	user := r.pool.Get().(*models.User)
	*user = *u
	r.users = append(r.users, user)

	return nil
}

func (r *userRepository) UpdateUser(u *models.User) error {
	if u == nil {
		// todo return error
		return nil
	}

	if u.ID == 0 {
		// todo return error
		return nil
	}

	r.mutex.Lock()
	defer r.mutex.Unlock()

	for i := range r.users {
		if r.users[i].ID == u.ID {
			*r.users[i] = *u
			break
		}
	}

	return nil
}

func (r *userRepository) DeleteUser(u *models.User) error {
	if u == nil {
		// todo return error
		return nil
	}

	if u.ID == 0 {
		// todo return error
		return nil
	}

	r.mutex.Lock()
	defer r.mutex.Unlock()

	for i := range r.users {
		if r.users[i].ID == u.ID {
			r.pool.Put(r.users[i])
			r.users = append(r.users[:i], r.users[i+1:]...)
			break
		}
	}

	return nil
}
