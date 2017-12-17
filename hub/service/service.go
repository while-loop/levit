package service

import (
	"github.com/while-loop/levit/hub"
	"github.com/while-loop/levit/hub/proto"
	"github.com/while-loop/levit/hub/stream"
)

//go:generate protoc -I ../proto --go_out=plugins=grpc:../proto ../proto/hub.proto

type HubService struct {
	hub *hub.Hub
}

func (h *HubService) Connect(s proto.Hub_ConnectServer) error {
	conn := hub.NewConn(h.hub, stream.NewGrpc(s))

	h.hub.Register(conn)
	conn.Loop()
	h.hub.Deregister(conn)
	return nil
}

func New(hub *hub.Hub) proto.HubServer {
	return &HubService{
		hub: hub,
	}
}
