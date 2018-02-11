package repo

import "github.com/while-loop/levit/users/proto"

type UsersRepository interface {
	Create(u *users.User) (*users.User, error)
	Get(id uint64) (*users.User, error)
	GetAll(ids ...uint64) ([]*users.User, error)
	Update(u *users.User) (*users.User, error)
	Delete(id uint64) error
}
