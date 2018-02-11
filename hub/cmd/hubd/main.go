package main

import (
	"flag"
	"strconv"
	"strings"

	"github.com/while-loop/levit/common/log"
	libservice "github.com/while-loop/levit/common/service"
	"github.com/while-loop/levit/hub"
	"github.com/while-loop/levit/hub/event"
	proto "github.com/while-loop/levit/hub/proto"
	"github.com/while-loop/levit/hub/service"
	"github.com/while-loop/levit/hub/version"
)

func init() {
	log.Infof("%s %s %s %s", version.Name, version.Version, version.BuildTime, version.Commit)
}

func main() {
	v := flag.Bool("v", false, version.Name+" version")
	laddr := flag.String("laddr", "0.0.0.0:8080", version.Name+" version")
	flag.Parse()

	if *v {
		// version is printed in init()
		return
	}

	parts := strings.Split(*laddr, ":")
	port, err := strconv.ParseInt(parts[1], 10, 32)
	if err != nil {
		log.Fatal(err)
	}

	rpc := libservice.NewGrpcService(libservice.Options{
		ServiceName:    version.Name,
		ServiceVersion: version.Version,
		MetricsAddr:    "0.0.0.0:8181",
		IP:             parts[0],
		Port:           int(port),
	})

	h := hub.New(event.GetHandler())
	go h.Start()
	proto.RegisterHubServer(rpc.GrpcServer(), service.New(h))
	log.Fatal(rpc.Serve())
}
