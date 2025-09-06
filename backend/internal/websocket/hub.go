package websocket

import (
	"log"
	"sync"
)

type Hub struct {
	clients    map[string]*ClientWS
	register   chan *ClientWS
	unregister chan *ClientWS
	broadcast  chan []byte
	mu         sync.RWMutex
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *ClientWS),
		unregister: make(chan *ClientWS),
		clients:    make(map[string]*ClientWS),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.mu.Lock()
			h.clients[client.id] = client
			h.mu.Unlock()
			log.Printf("Client %s connected", client.id)

		case client := <-h.unregister:
			h.mu.Lock()
			if _, ok := h.clients[client.id]; ok {
				delete(h.clients, client.id)
				close(client.send)
			}
			h.mu.Unlock()
			log.Printf("Client %s disconnected", client.id)

		case message := <-h.broadcast:
			h.mu.RLock()
			for _, client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client.id)
				}
			}
			h.mu.RUnlock()
		}
	}
}