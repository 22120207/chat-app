package controllers

import (
	"chat-app-backend/internal/config"
	"chat-app-backend/internal/mongodb"
)

type Collection string

const (
	UserCollection Collection = "users"
)

var Client mongodb.Client

var conf config.Config

func init() {
	conf = config.LoadEnv()

	Client.ConnectToDB(conf.MongodbUri)
}
