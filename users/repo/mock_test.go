package repo

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	proto "github.com/while-loop/levit/users/proto"
)

func TestMockRepo_CRUD(t *testing.T) {
	a := require.New(t)
	m := NewMockRepo()

	u := createUser(1)
	updated, err := m.CreateUser(u)
	a.NoError(err)
	a.Equal(u, updated)

	u2 := createUser(2)
	updated2, err := m.CreateUser(u2)
	a.NoError(err)
	a.Equal(u2, updated2)

	us, err := m.GetUsers(1, 2)
	a.NoError(err)
	a.Len(us, 2)

	u2.First = "first3"
	updated3, err := m.UpdateUser(u2)
	a.NoError(err)
	a.Equal(u2, updated3)

	a.NoError(m.DeleteUser(1))
	us, err = m.GetUsers(1, 2)
	a.NoError(err)
	a.Len(us, 1)
	a.Equal(us[0].First, "first3")

	newUser, err := m.GetUser(2)
	a.NoError(err)
	a.NotEqual(newUser.Id, 0)

	a.NoError(m.DeleteUser(2))
	us, err = m.GetUsers(2)
	a.NoError(err)
	a.Len(us, 0)

	_, err = m.GetUser(2)
	a.Error(err)
}

func createUser(idx uint64) *proto.User {
	index := fmt.Sprintf("%d", idx)
	return &proto.User{
		Id:         uint64(idx),
		CreatedAt:  time.Now().Unix(),
		Deleted:    idx%2 == 0,
		FacebookId: "facebook" + index,
		First:      "first" + index,
		GoogleId:   "google" + index,
		Last:       "last" + index,
	}
}
