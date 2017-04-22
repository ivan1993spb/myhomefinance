package repository

import "github.com/ivan1993spb/myhomefinance/models"

type UserRepository interface {
	CreateUser(u *models.User) error
	UpdateUser(u *models.User) error
	DeleteUser(u *models.User) error
}
