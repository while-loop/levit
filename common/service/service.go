package service

import (
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
)

type Service interface {
	Serve() error
	GracefulStop() error
	Options() Options
	GrpcServer() *grpc.Server
	Register()
	Deregister()
}

func CtrlCSig() chan os.Signal {
	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGABRT)
	return sigs
}
