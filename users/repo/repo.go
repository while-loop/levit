package repo

import "github.com/while-loop/levit/users/proto"

type UsersRepository interface {
	CreateUser(users.User) (users.User, error)
	GetUser(ids uint64) ([]users.User, error)
	UpdateUser(user users.User) (users.User, error)
	DeleteUser(id uint64) error
}
