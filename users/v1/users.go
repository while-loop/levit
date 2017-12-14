package users

import (
	"github.com/while-loop/levit/log"
	"github.com/while-loop/levit/service"
	"github.com/while-loop/levit/users"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

//go:generate protoc -I ./ --go_out=plugins=grpc:./ ./users.proto

const (
	name = users.Name
)

func init() {
	service.Register(name, New)
}

type usersService struct {
}

func (c *usersService) Get(context.Context, *GetRequest) (*User, error) {
	panic("implement me")
}

func (c *usersService) Update(context.Context, *UpdateRequest) (*User, error) {
	panic("implement me")
}

func New(config interface{}) service.Service {
	return &usersService{}
}

func (c *usersService) Register(rpc *grpc.Server) {
	RegisterUsersServer(rpc, c)
	log.Info("UsersService started")
}
