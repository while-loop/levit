package service

import (
	"fmt"
	"testing"
	"time"

	"context"

	"github.com/stretchr/testify/require"
	proto "github.com/while-loop/levit/users/proto"
	"github.com/while-loop/levit/users/repo"
)

var c = context.Background()

func TestUsersService(t *testing.T) {
	a := require.New(t)
	srvc := New(repo.NewMockRepo(), nil)

	u := createUser(1)
	updated, err := srvc.Create(c, u)
	a.NoError(err)
	a.Equal(u, updated.User)

	u2 := createUser(2)
	updated2, err := srvc.Create(c, u2)
	a.NoError(err)
	a.Equal(u2, updated2.User)

	us, err := srvc.GetAll(c, &proto.GetRequest{Ids: []uint64{1, 2}})
	a.NoError(err)
	a.Len(us.Users, 2)

	u2.First = "first3"
	updated3, err := srvc.Update(c, u2)
	a.NoError(err)
	a.Equal(u2, updated3.User)

	newUser, err := srvc.Get(c, &proto.User{Id: 2})
	a.NoError(err)
	a.NotEqual(newUser.User.Id, 0)

	get2, err := srvc.Get(c, &proto.User{Id: 2})
	a.NoError(err)
	a.Equal(u2, get2.User)
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
