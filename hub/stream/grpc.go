package stream

import "github.com/while-loop/levit/hub/proto"

type grpcStream struct {
	proto.Hub_ConnectServer
}

func NewGrpc(s proto.Hub_ConnectServer) Stream {
	return &grpcStream{s}
}

func (g *grpcStream) Close() error {
	return nil
}
