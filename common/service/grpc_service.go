package service

import (
	"net"
	"net/http"
	"time"

	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/while-loop/levit/common/log"
	"google.golang.org/grpc"
)

var _ Service = &grpcService{}

type grpcService struct {
	*grpc.Server
	options Options
}

func NewGrpcService(options Options) Service {
	rpc := grpc.NewServer(
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
		grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
	)

	if options.MetricsAddr != "" {
		metrics(rpc, options.MetricsAddr)
	}

	options.applyDefaults()

	s := &grpcService{rpc,
		options,
	}
	return s
}

func (s *grpcService) GrpcServer() *grpc.Server {
	return s.Server
}

func (s *grpcService) Options() Options {
	return s.options
}

func (s *grpcService) Serve() error {
	laddr := s.options.laddr()
	lis, err := net.Listen("tcp", laddr)
	for err != nil {
		log.Error("failed to get service listener", err)
		time.Sleep(3 * time.Second)
		lis, err = net.Listen("tcp", laddr)
	}

	log.Info("Running UsersService gRPC Server ", laddr)

	exitCh := make(chan error)
	go func() {
		exitCh <- s.Server.Serve(lis)
	}()

	err = nil
	running := true
	sig := CtrlCSig()
	for running {
		select {
		case err = <-exitCh:
			log.Info("Server has stopped serving ", err)
		case <-sig:
			log.Warn("Ctrl+C captured")
			s.GracefulStop()
			running = false
		}
	}

	return err
}

func (s *grpcService) GracefulStop() error {
	log.Infof("Gracefully stopping %s...", s.options.ServiceName)
	s.Server.GracefulStop()

	return nil
}

func metrics(rpc *grpc.Server, httpAddr string) {
	l, err := net.Listen("tcp", httpAddr)
	if err != nil {
		log.Fatal("no metrics for you", err)
		return
	}

	grpc_prometheus.Register(rpc)
	m := http.NewServeMux()
	m.Handle("/metrics", promhttp.Handler())
	go http.Serve(l, m)
	log.Infof("Running metrics %s/metrics", l.Addr())
}
