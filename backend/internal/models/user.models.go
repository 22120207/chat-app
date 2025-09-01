package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	ID         bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Fullname   string        `bson:"fullname" json:"fullname" binding:"required"`
	Username   string        `bson:"username" json:"username" binding:"required"`
	Password   string        `bson:"password" json:"-" binding:"required,min=6"`
	Gender     string        `bson:"gender" json:"gender" binding:"required,oneof=male female"`
	ProfilePic string        `bson:"profilePic" json:"profilePic"`
	CreatedAt  time.Time     `bson:"createdAt" json:"createdAt"`
	UpdatedAt  time.Time     `bson:"updatedAt" json:"updatedAt"`
}
