package repo

import (
	"fmt"

	proto "github.com/while-loop/levit/users/proto"
)

type mockRepo struct {
	users map[uint64]*proto.User
}

func NewMockRepo() UsersRepository {
	return &mockRepo{users: map[uint64]*proto.User{}}
}

func (m *mockRepo) Create(u *proto.User) (*proto.User, error) {
	return m.Update(u)
}

func (m *mockRepo) GetAll(ids ...uint64) ([]*proto.User, error) {
	users := make([]*proto.User, 0)
	for _, id := range ids {
		if user, exists := m.users[id]; exists {
			users = append(users, user)
		}
	}
	return users, nil
}
func (m *mockRepo) Get(id uint64) (*proto.User, error) {
	if user, exists := m.users[id]; exists {
		return user, nil
	}

	return nil, fmt.Errorf("dne")
}

func (m *mockRepo) Update(u *proto.User) (*proto.User, error) {
	m.users[u.Id] = u
	return u, nil
}

func (m *mockRepo) Delete(id uint64) error {
	delete(m.users, id)
	return nil
}
