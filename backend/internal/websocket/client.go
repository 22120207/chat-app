package websocket

import (
	"chat-app-backend/internal/controllers"
	"chat-app-backend/internal/models"
	"encoding/json"
	"errors"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/v2/bson"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
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

	senderId, err := bson.ObjectIDFromHex(c.id)
	if err != nil {
		log.Printf("Invalid receiver ID: %v", err)
		return
	}

	for {
		_, rawMsg, err := c.conn.ReadMessage()
		if err != nil {
			break
		}

		var incoming struct {
			ReceiverID string `json:"receiverId"`
			Message    string `json:"message"`
		}
		if err := json.Unmarshal(rawMsg, &incoming); err != nil {
			log.Printf("Error in unmarshal message: %v", err)
			continue
		}

		receiverObjID, err := bson.ObjectIDFromHex(incoming.ReceiverID)
		if err != nil {
			log.Printf("Invalid receiver ID: %v", err)
			continue
		}

		// Save to MongoDB
		msg, err := controllers.CreateMessage(senderId, receiverObjID, incoming.Message)
		if err != nil {
			log.Printf("Error saving message: %v", err)
			continue
		}

		jsonMsg, err := json.Marshal(msg)
		if err != nil {
			log.Printf("Error marshaling message: %v", err)
			continue
		}

		// Echo back to sender
		c.send <- jsonMsg

		// Send to receiver if online
		hub.mu.RLock()
		receiver, ok := hub.clients[receiverObjID.Hex()]
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