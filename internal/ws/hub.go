package ws

import (
	"log"
)

type Hub struct {
	clients   map[*Client]bool
	Broadcast chan []byte
}

func MakeHub() *Hub {
	return &Hub{
		clients:   make(map[*Client]bool),
		Broadcast: make(chan []byte),
	}
}

func (h *Hub) AddClient(c *Client) {
	h.clients[c] = true
}

func (h *Hub) removeClient(c *Client) {
	delete(h.clients, c)
}

func (h *Hub) Run() {
	for {
		message := <-h.Broadcast

		for c := range h.clients {
			select {
			case c.Send <- message:
				log.Println("broadcast to client")
			default:
				close(c.Send)
				delete(h.clients, c)
			}
		}
	}
}
