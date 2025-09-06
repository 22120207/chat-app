package websocket

import (
	"chat-app-backend/internal/models"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool {
        origin := r.Header.Get("Origin")
        return origin == "http://localhost:5173"
    },
}

type ClientWS struct {
	id string
	conn *websocket.Conn
	send chan []byte
}

func (c *ClientWS) readPump(hub *Hub) {
	defer func() {
		hub.unregister <- c
		c.conn.Close()
	}()

	for {
		_, rawMsg, err := c.conn.ReadMessage()
		if err != nil {
			break
		}

		var incoming models.Message
		if err := json.Unmarshal(rawMsg, &incoming); err != nil {
			log.Printf("Error in unmarshal message: %v", err)
			continue
		}

		jsonMsg, err := json.Marshal(incoming)
		if err != nil {
			log.Printf("Error in marshal incoming websocket message: %v", err)
		}

		// Send to receiver if online
		hub.mu.RLock()
		receiver, ok := hub.clients[incoming.ReceiverID.Hex()]
		hub.mu.RUnlock()
		if ok {
			receiver.send <- jsonMsg
		}
	}
}

func (c *ClientWS) writePump() {
	for msg := range c.send {
		err := c.conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			break
		}
	}
}

func ServeWs(hub *Hub, c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.Error(errors.New(string(models.UnauthorizedError)))
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Error in ServeWs: %v", err)
		return
	}

	sender := user.(models.User)

	client := &ClientWS{
		id:   sender.ID.Hex(),
		conn: conn,
		send: make(chan []byte, 256),
	}

	hub.register <- client

	go client.writePump()
	go client.readPump(hub)
}