package mongodb

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Collection string

type Client struct {
	Context        context.Context
	Database       *mongo.Database
	UserCollection *mongo.Collection
	client         *mongo.Client
}

func (c *Client) ConnectToDB(uri string) {
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable.")
	}

	// Uses the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	// Defines the options for the MongoDB client
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	// Creates a new client and connects to the server
	client, err := mongo.Connect(opts)
	if err != nil {
		panic(err)
	}

	// Sends a ping to confirm a successful connection
	var result bson.M
	if err := client.Database("chat_app").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	c.Context = context.Background()
	c.client = client
	c.Database = client.Database("chat_app")
	c.UserCollection = c.Database.Collection("users")
}

func (c *Client) FindOne(collection string, filter any) *mongo.SingleResult {
	switch collection {
	case "users":
		return c.UserCollection.FindOne(c.Context, filter)
	default:
		return nil
	}
}

func (c *Client) InsertOne(collection string, document any) (*mongo.InsertOneResult, error) {
	switch collection {
	case "users":
		return c.UserCollection.InsertOne(c.Context, document)
	default:
		return nil, fmt.Errorf("unknown collection: %s", collection)
	}
}

func (c *Client) Disconnect() {
	if c.client != nil {
		if err := c.client.Disconnect(c.Context); err != nil {
			log.Printf("Error disconnecting MongoDB: %v", err)
		} else {
			fmt.Println("MongoDB disconnected")
		}
	}
}
