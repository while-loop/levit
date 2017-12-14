package main

import (
	"flag"
	"fmt"
	"net"

	"github.com/while-loop/levit/log"
	"github.com/while-loop/levit/service"
	"google.golang.org/grpc/reflection"
)

const (
	Version = "1"
	Name    = "levit.users"
)

func main() {
	version := flag.Bool("v", false, Name+" version")
	laddr := *flag.String("laddr", ":8080", "binding address")
	flag.Parse()

	if *version {
		fmt.Printf("%s version %s\n", Name, Version)
		return
	}

	lis, err := net.Listen("tcp", laddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	rpc := service.NewRpcServer()
	service.Start(nil, rpc)
	// Register reflection service on gRPC server.
	reflection.Register(rpc)

	log.Info("Running gRPC Server", lis.Addr())
	if err := rpc.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
