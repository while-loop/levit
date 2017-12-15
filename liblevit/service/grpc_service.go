package service

import (
	"google.golang.org/grpc"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/while-loop/levit/liblevit/log"
	"net"
	"time"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"github.com/while-loop/levit/liblevit/registry"
	"github.com/google/uuid"
)

var _ Service = &grpcService{}

type grpcService struct {
	*grpc.Server
	options Options
}

func (s *grpcService) GrpcServer() *grpc.Server {
	return s.Server
}

func (s *grpcService) Options() Options {
	return s.options
}

func NewGrpcService(options Options) Service {
	rpc := grpc.NewServer(
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
		grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
	)

	if options.MetricsAddr != "" {
		metrics(rpc, options.MetricsAddr)
	}

	if options.Port <= 0 {
		options.Port = 8080
	}
	options.Uuid = uuid.New().String()

	s := &grpcService{rpc,
		options,
	}
	return s
}

func (s *grpcService) Serve() error {
	laddr := s.options.laddr()
	lis, err := net.Listen("tcp", laddr)
	for err != nil {
		log.Error("failed to get service listener", err)
		time.Sleep(3 * time.Second)
		lis, err = net.Listen("tcp", laddr)
	}

	log.Info("Running UsersService gRPC Server", laddr)
	srv := s.regService()
	err = registry.Register(srv)
	for err != nil {
		log.Errorf("failed to register %s: %v", s.options.ServiceName, err)
		time.Sleep(3 * time.Second)
		err = registry.Register(srv)
	}

	return s.Server.Serve(lis)
}

func (s *grpcService) GracefulStop() error {
	log.Infof("Gracefully stopping %s...", s.options.ServiceName)
	if err := registry.Deregister(s.regService()); err != nil {
		log.Errorf("Failed to deregister %s: %v", s.options.ServiceName, err)
	}
	s.Server.GracefulStop()

	return nil
}

func metrics(rpc *grpc.Server, httpAddr string) {
	l, err := net.Listen("tcp", httpAddr)
	if err != nil {
		log.Fatal("no metrics for you")
		return
	}

	grpc_prometheus.Register(rpc)
	m := http.NewServeMux()
	m.Handle("/metrics", promhttp.Handler())
	go http.Serve(l, m)
	log.Infof("Running metrics %s/metrics", l.Addr())
}

func (s *grpcService) regService() registry.Service {
	return registry.Service{
		Version: s.options.ServiceVersion,
		Name:    s.options.ServiceName,
		Instances: map[string]registry.Instance{
			s.options.Uuid: {
				Port: s.options.Port,
				IP:   s.options.IP,
				UUID: s.options.Uuid,
			},
		},
	}
}
