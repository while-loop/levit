package hub

import (
	"github.com/while-loop/levit/common/log"
	"github.com/while-loop/levit/hub/proto"
)

const (
	DefaultBufferedChannelSize = 100
)

type Hub struct {
	conns        map[uint64]*Conn
	register     chan *Conn
	deregister   chan *Conn
	broadcast    chan *proto.HubMessage
	stop         chan struct{}
	closed       chan struct{}
	EventsRouter Router
}

func New(events Router) *Hub {
	return &Hub{
		register:     make(chan *Conn, 1),
		deregister:   make(chan *Conn, 1),
		conns:        map[uint64]*Conn{},
		broadcast:    make(chan *proto.HubMessage, DefaultBufferedChannelSize),
		stop:         make(chan struct{}),
		closed:       make(chan struct{}),
		EventsRouter: events,
	}
}

func (h *Hub) Register(conn *Conn) {
	h.register <- conn
}

func (h *Hub) Deregister(conn *Conn) {
	h.deregister <- conn
}

func (h *Hub) Broadcast(message *proto.HubMessage) {
	h.broadcast <- message
}

func (h *Hub) Start() {
	defer func() {
		h.closed <- struct{}{}
	}()

	log.Info("Starting Hub loop")
	for {
		select {
		case conn := <-h.register:
			h.conns[conn.UserId] = conn
			// TODO set status to online (online = true)
		case conn := <-h.deregister:
			if _, exists := h.conns[conn.UserId]; !exists {
				continue
			}

			delete(h.conns, conn.UserId)
			// TODO set status to offline (online = false, last_seen = time.Now())

		case message := <-h.broadcast:
			for uid, conn := range h.conns {
				if !conn.Contains(message) {
					continue
				}

				select {
				case conn.Send() <- message:
				default:
					log.Errorf("Failed to send message to user %d: %s", uid, message)
				}
			}
		case <-h.stop:
			log.Warn("Hub recvd stop from channel")
			uids := make([]uint64, 0)

			// close conns fast so user doesn't miss out on unsent messages
			for _, c := range h.conns {
				uids = append(uids, c.UserId)
				c.Close()
			}

			// close conns fast so user doesn't miss out on unsent messages
			for uid := range uids {
				log.Debug(uid)
				// TODO set status to offline (online = false, last_seen = time.Now())
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
