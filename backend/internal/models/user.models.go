package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	ID         bson.ObjectID `bson:"_id,omitempty"`
	Fullname   string        `bson:"fullname" binding:"required"`
	Username   string        `bson:"username" binding:"required"`
	Password   string        `bson:"password" binding:"required,min=6"`
	Gender     string        `bson:"gender" binding:"required,oneof=male female"`
	ProfilePic string        `bson:"profilePic" default:""`
}
