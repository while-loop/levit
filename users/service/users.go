package service

import (
	"fmt"

	"github.com/while-loop/levit/common/log"
	proto "github.com/while-loop/levit/users/proto"
	"github.com/while-loop/levit/users/repo"
	"golang.org/x/net/context"
)

//go:generate protoc -I ../proto --go_out=plugins=grpc:../proto ../proto/users.proto

type UsersService struct {
	repo repo.UsersRepository
}

// TODO input pre-processing

func (c *UsersService) Create(ctx context.Context, u *proto.User) (*proto.User, error) {
	return c.Update(ctx, u)
}

func (c *UsersService) GetAll(ctx context.Context, req *proto.GetRequest) (*proto.UsersArr, error) {
	us, err := c.repo.GetUsers(req.Ids...)
	if err != nil {
		return nil, err
	}

	return &proto.UsersArr{Users: us}, nil
}

func (c *UsersService) Get(ctx context.Context, req *proto.GetRequest) (*proto.User, error) {
	log.Debugf("UsersService:Get %#v", req)
	if len(req.Ids) == 1 {
		return c.repo.GetUser(req.Ids[0])
	}

	return nil, fmt.Errorf("bad request")
}

func (c *UsersService) Update(ctx context.Context, req *proto.User) (*proto.User, error) {
	return c.repo.UpdateUser(req)
}

func New(repository repo.UsersRepository) proto.UsersServer {
	return &UsersService{repository}
}
