package controllers

import (
	"chat-app-backend/internal/config"
	"chat-app-backend/internal/mongodb"
)

var Client mongodb.Client

var conf config.Config

func init() {
	config.LoadEnv()

	conf = config.Conf

	Client.ConnectToDB(conf.MongodbUri, conf.Database)
}
