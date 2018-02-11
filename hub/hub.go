package hub

import (
	"github.com/while-loop/levit/common/log"
	"github.com/while-loop/levit/hub/proto"
)

const (
	DefaultBufferedChannelSize = 100
)

type connFunc func(map[uint64]*Conn)

type Hub struct {
	stop         chan struct{}
	closed       chan struct{}
	connChan     chan connFunc
	EventsRouter Router
}

func New(events Router) *Hub {
	return &Hub{
		connChan:     make(chan connFunc, DefaultBufferedChannelSize),
		stop:         make(chan struct{}),
		closed:       make(chan struct{}),
		EventsRouter: events,
	}
}

func (h *Hub) Register(conn *Conn) {
	h.connChan <- func(conns map[uint64]*Conn) {
		conns[conn.UserId] = conn
		// TODO set status to online (online = true)
	}
}

func (h *Hub) Deregister(conn *Conn) {
	h.connChan <- func(conns map[uint64]*Conn) {
		if _, exists := conns[conn.UserId]; !exists {
			return
		}

		delete(conns, conn.UserId)
		// TODO set status to offline (online = false, last_seen = time.Now())
	}
}

func (h *Hub) Broadcast(message *proto.HubMessage) {
	h.connChan <- func(conns map[uint64]*Conn) {
		for _, conn := range conns {
			if !conn.Contains(message) {
				continue
			}

			go func(conn *Conn) {
				if err := conn.Send(message); err != nil {
					log.Error(err)
				}
			}(conn)
		}
	}
}

func (h *Hub) Start() {
	conns := map[uint64]*Conn{}
	defer func() {
		h.closed <- struct{}{}
	}()
	log.Info("Starting Hub loop")
	for {
		select {
		case fn := <-h.connChan:
			fn(conns)
		case <-h.stop:
			log.Warn("Hub recvd stop from channel")
			uids := make([]uint64, 0)

			// close conns fast so user doesn't miss out on unsent messages
			for _, c := range conns {
				uids = append(uids, c.UserId)
				go c.Close()
				// go setOffline(c.UserId) // TODO set status to offline (online = false, last_seen = time.Now())
			}

			return
		}
	}
}

func (h *Hub) Stop() {
	h.stop <- struct{}{}
	// wait for all connections to shutdown
	<-h.closed
}
