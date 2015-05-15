package rooms

import (
	conn "namelychat/connections"

	"github.com/satori/go.uuid"
)

type Room struct {
	Id          string
	Name        string
	Private     bool
	Members     map[*User]bool
	connections map[*conn.connection]bool
	broadcast   chan []byte
	register    chan *conn.connection
	unregister  chan *conn.connection
}

func newRoom(name string, private bool) *hub {
	id := uuid.NewV4().String()
	if name == "" {
		name = id
	}
	return &hub{
		Name:        name,
		Id:          id,
		Private:     private,
		Members:     make(map[*User]bool),
		broadcast:   make(chan []byte),
		register:    make(chan *conn.connection),
		unregister:  make(chan *conn.connection),
		connections: make(map[*conn.connection]bool),
	}
}

func (h *Room) run() {
	for {
		select {
		case c := <-h.register:
			h.connections[c] = true

		case c := <-h.unregister:
			if _, ok := h.connections[c]; ok {
				delete(h.connections, c)
				close(c.send)
			}

		case m := <-h.broadcast:
			for c := range h.connections {
				select {
				case c.send <- m:
				default:
					delete(h.connections, c)
					close(c.send)
				}
			}
		}
	}
}
