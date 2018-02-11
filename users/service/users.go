package service

import (
	pb "github.com/while-loop/levit/users/proto"
	"github.com/while-loop/levit/users/repo"
	"golang.org/x/net/context"
)

//go:generate protoc -I ../proto --go_out=plugins=grpc:../proto ../proto/users.proto

type UsersService struct {
	repo    repo.UsersRepository
	tknSrvc *TokenService
}

func New(repository repo.UsersRepository, service *TokenService) pb.UsersServer {
	return &UsersService{repository, service}
}

func (u *UsersService) Create(ctx context.Context, req *pb.User) (*pb.Response, error) {
	return u.Update(ctx, req)
}

func (u *UsersService) Get(ctx context.Context, req *pb.User) (*pb.Response, error) {
	user, err := u.repo.Get(req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.Response{User: user}, nil
}

func (u *UsersService) Update(ctx context.Context, req *pb.User) (*pb.Response, error) {
	user, err := u.repo.Update(req)
	if err != nil {
		return nil, err
	}

	return &pb.Response{User: user}, nil

}

func (u *UsersService) GetAll(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	users, err := u.repo.GetAll(req.Ids...)
	if err != nil {
		return nil, err
	}

	return &pb.Response{Users: users}, nil
}

func (u *UsersService) Auth(ctx context.Context, ser *pb.User) (*pb.Token, error) {
	panic("implement me")
}

func (u *UsersService) ValidateToken(ctx context.Context, req *pb.Token) (*pb.Token, error) {
	panic("implement me")
}
