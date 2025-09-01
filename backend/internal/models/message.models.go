package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Message struct {
	ID         bson.ObjectID `bson:"_id,omitempty"`
	SenderID   bson.ObjectID `bson:"senderId" binding:"required"`
	ReceiverID bson.ObjectID `bson:"receiverId" binding:"required"`
	Message    string        `bson:"message" binding:"required"`
	CreatedAt  time.Time     `bson:"createdAt"`
	UpdatedAt  time.Time     `bson:"updatedAt"`
}

type MessageRequest struct {
	Message string `json:"message" binding:"required"`
}
