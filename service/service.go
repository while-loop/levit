package service

import (
	"net/http"

	"net"

	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/while-loop/levit/log"
	"google.golang.org/grpc"
)

type RegisterFunc func(config interface{}) Service
type Service interface {
	Register(rpc *grpc.Server)
}

var services = map[string]RegisterFunc{}

func Register(name string, construc RegisterFunc) {
	services[name] = construc
}

func Start(app interface{}, rpc *grpc.Server) {
	for name, construct := range services {
		log.Infof("Starting %s...\n", name)
		construct(app).Register(rpc)
	}
}

func NewRpcServer() *grpc.Server {
	rpc := grpc.NewServer(
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
		grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
	)
	Metrics(rpc)
	return rpc
}

func Metrics(rpc *grpc.Server) {
	l, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatal("no metrics for you")
		return
	}

	grpc_prometheus.Register(rpc)
	m := http.NewServeMux()
	m.Handle("/metrics", promhttp.Handler())
	go http.Serve(l, m)
	log.Info("Running metrics", l.Addr())
}
