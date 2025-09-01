package mongodb

import (
	"chat-app-backend/internal/models"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Client struct {
	Context  context.Context
	Database *mongo.Database
	client   *mongo.Client
}

func (c *Client) ConnectToDB(uri, database string) {
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable.")
	}

	// Set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	// Defines the options for the MongoDB client
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	// Connects to the server
	client, err := mongo.Connect(opts)
	if err != nil {
		panic(err)
	}

	// Sends a ping to confirm a successful connection
	var result bson.M
	if err := client.Database(database).RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	c.Context = context.Background()
	c.client = client
	c.Database = client.Database(database)
}

func (c *Client) FindOne(collection models.Collection, filter any) *mongo.SingleResult {
	coll := c.Database.Collection(string(collection))
	return coll.FindOne(c.Context, filter)
}

func (c *Client) Find(collection models.Collection, filter any) (*mongo.Cursor, error) {
	coll := c.Database.Collection(string(collection))
	return coll.Find(c.Context, filter)
}

func (c *Client) InsertOne(collection models.Collection, document any) (*mongo.InsertOneResult, error) {
	coll := c.Database.Collection(string(collection))
	return coll.InsertOne(c.Context, document)
}

func (c *Client) UpdateOne(collection models.Collection, filter any, update any) (*mongo.UpdateResult, error) {
	coll := c.Database.Collection(string(collection))
	return coll.UpdateOne(c.Context, filter, update)
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
