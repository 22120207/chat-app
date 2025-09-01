package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Conversation struct {
	ID           bson.ObjectID   `bson:"_id,omitempty"`
	Participants []bson.ObjectID `bson:"participants" binding:"required"`
	Messages     []bson.ObjectID `bson:"messages" binding:"required"`
	CreatedAt    time.Time       `bson:"createdAt"`
	UpdatedAt    time.Time       `bson:"updatedAt"`
}
