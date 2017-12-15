package repo

import "github.com/while-loop/levit/users/proto"

type UsersRepository interface {
	CreateUser(u *users.User) (*users.User, error)
	GetUser(id uint64) (*users.User, error)
	GetUsers(ids ...uint64) ([]*users.User, error)
	UpdateUser(u *users.User) (*users.User, error)
	DeleteUser(id uint64) error
}
