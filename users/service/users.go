package service

import (
	proto "github.com/while-loop/levit/users/proto"
	"golang.org/x/net/context"
	"github.com/while-loop/levit/liblevit/log"
)

//go:generate protoc -I ../proto --go_out=plugins=grpc:../proto ../proto/users.proto

type UsersService struct {
}

func (c *UsersService) Get(ctx context.Context, req *proto.GetRequest) (*proto.User, error) {
	log.Debugf("UsersService:Get %#v", req)
	return &proto.User{
		FirstName: "Tobias",
	}, nil
}

func (c *UsersService) Update(ctx context.Context, req *proto.User) (*proto.User, error) {
	log.Debugf("UsersService:Update %#v", req)
	return req, nil
}

func New() *UsersService {
	return &UsersService{}
}
