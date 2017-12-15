package service

import "google.golang.org/grpc"

type Service interface{
	Serve() error
	GracefulStop() error
	Options() Options
	GrpcServer() *grpc.Server
}
