package main

import (
	"flag"

	"github.com/while-loop/levit/common/log"
	"github.com/while-loop/levit/common/registry"
	libservice "github.com/while-loop/levit/common/service"
	proto "github.com/while-loop/levit/users/proto"
	"github.com/while-loop/levit/users/repo"
	"github.com/while-loop/levit/users/service"
	"github.com/while-loop/levit/users/version"
)

func init() {
	log.Infof("%s %s %s %s", version.Name, version.Version, version.BuildTime, version.Commit)
}

func main() {
	v := flag.Bool("v", false, version.Name+" version")
	flag.Parse()

	if *v {
		// version is printed in init()
		return
	}

	registry.Use(registry.NewStub())

	rpc := libservice.NewGrpcService(libservice.Options{
		ServiceName:    version.Name,
		ServiceVersion: version.Version,
		MetricsAddr:    ":8181",
	})

	proto.RegisterUsersServer(rpc.GrpcServer(), service.New(repo.NewMockRepo()))
	log.Fatal(rpc.Serve())
}
