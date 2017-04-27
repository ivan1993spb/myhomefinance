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

type errCreateUser string

func (e errCreateUser) Error() string {
	return "cannot create user: " + string(e)
}

func (r *userRepository) CreateUser(u *models.User) error {
	if u == nil {
		return errCreateUser("passed nil user")
	}

	if u.UUID != 0 {
		return errCreateUser("passed user has an identifier")
	}

	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.cursorID++
	u.UUID = r.cursorID

	user := r.pool.Get().(*models.User)
	*user = *u
	r.users = append(r.users, user)

	return nil
}

type errUpdateUser string

func (e errUpdateUser) Error() string {
	return "cannot update user: " + string(e)
}

func (r *userRepository) UpdateUser(u *models.User) error {
	if u == nil {
		return errUpdateUser("passed nil user")
	}

	if u.UUID == 0 {
		return errUpdateUser("passed user has zero identifier")
	}

	r.mutex.Lock()
	defer r.mutex.Unlock()

	for i := range r.users {
		if r.users[i].UUID == u.UUID {
			*r.users[i] = *u
			break
		}
	}

	return nil
}

type errDeleteUser string

func (e errDeleteUser) Error() string {
	return "cannot delete user: " + string(e)
}

func (r *userRepository) DeleteUser(u *models.User) error {
	if u == nil {
		return errDeleteUser("passed nil user")
	}

	if u.UUID == 0 {
		return errDeleteUser("passed user does not have identifier")
	}

	r.mutex.Lock()
	defer r.mutex.Unlock()

	for i := range r.users {
		if r.users[i].UUID == u.UUID {
			r.pool.Put(r.users[i])
			r.users = append(r.users[:i], r.users[i+1:]...)
			break
		}
	}

	return nil
}
