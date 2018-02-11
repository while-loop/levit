package repo

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	pb "github.com/while-loop/levit/users/proto"
)

func testImpl(t *testing.T, r UsersRepository) {
	a := require.New(t)

	u := createUser(1)
	updated, err := r.Create(u)
	a.NoError(err)
	a.Equal(u, updated)

	u2 := createUser(2)
	updated2, err := r.Create(u2)
	a.NoError(err)
	a.Equal(u2, updated2)

	us, err := r.GetAll(1, 2)
	a.NoError(err)
	a.Len(us, 2)

	u2.First = "first3"
	updated3, err := r.Update(u2)
	a.NoError(err)
	a.Equal(u2, updated3)

	a.NoError(r.Delete(1))
	us, err = r.GetAll(1, 2)
	a.NoError(err)
	a.Len(us, 1)
	a.Equal(us[0].First, "first3")

	newUser, err := r.Get(2)
	a.NoError(err)
	a.NotEqual(newUser.Id, 0)

	a.NoError(r.Delete(2))
	us, err = r.GetAll(2)
	a.NoError(err)
	a.Len(us, 0)

	_, err = r.Get(2)
	a.Error(err)
}

func createUser(idx uint64) *pb.User {
	index := fmt.Sprintf("%d", idx)
	return &pb.User{
		Id:         uint64(idx),
		CreatedAt:  time.Now().Unix(),
		Deleted:    idx%2 == 0,
		FacebookId: "facebook" + index,
		First:      "first" + index,
		GoogleId:   "google" + index,
		Last:       "last" + index,
	}
}
