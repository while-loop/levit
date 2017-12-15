package main

import (
	"flag"
	"fmt"

	proto "github.com/while-loop/levit/users/proto"
	"github.com/while-loop/levit/users/service"
	"github.com/while-loop/levit/users/config"
	"github.com/while-loop/levit/liblevit/registry"
	"github.com/while-loop/levit/liblevit/registry/stub"
	libservice "github.com/while-loop/levit/liblevit/service"
	"github.com/while-loop/levit/liblevit/log"
)

func main() {
	version := flag.Bool("v", false, config.Name+" version")
	flag.Parse()

	if *version {
		fmt.Printf("%s version %s\n", config.Name, config.Version)
		return
	}

	registry.Use(stub.New())
	rpc := libservice.NewGrpcService(libservice.Options{
		ServiceName:    config.Name,
		ServiceVersion: config.Version,
		MetricsAddr:    ":8181",
	})

	proto.RegisterUsersServer(rpc.GrpcServer(), service.New())
	log.Fatal(rpc.Serve())
}
