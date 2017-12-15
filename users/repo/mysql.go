package repo

import "github.com/while-loop/levit/users/proto"


func NewMySql() {

}

type mySqlRepo struct {

}

func (mySqlRepo) CreateUser(users.User) (users.User, error) {
	panic("implement me")
}

func (mySqlRepo) GetUser(ids uint) ([]users.User, error) {
	panic("implement me")
}

func (mySqlRepo) UpdateUser(user users.User) (users.User, error) {
	panic("implement me")
}

func (mySqlRepo) DeleteUser(id uint) error {
	panic("implement me")
}

