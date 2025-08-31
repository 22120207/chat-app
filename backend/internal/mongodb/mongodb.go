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
	Context                 context.Context
	Database                *mongo.Database
	UsersCollection         *mongo.Collection
	MessagesCollection      *mongo.Collection
	ConversationsCollection *mongo.Collection
	client                  *mongo.Client
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

	c.UsersCollection = c.Database.Collection(string(models.UsersCollection))
	c.MessagesCollection = c.Database.Collection(string(models.MessagesCollection))
	c.ConversationsCollection = c.Database.Collection(string(models.ConversationsCollection))
}

func (c *Client) FindOne(collection models.Collection, filter any) *mongo.SingleResult {
	switch collection {
	case models.UsersCollection:
		return c.UsersCollection.FindOne(c.Context, filter)
	case models.MessagesCollection:
		return c.MessagesCollection.FindOne(c.Context, filter)
	case models.ConversationsCollection:
		return c.ConversationsCollection.FindOne(c.Context, filter)
	default:
		return nil
	}
}

func (c *Client) InsertOne(collection models.Collection, document any) (*mongo.InsertOneResult, error) {
	switch collection {
	case models.UsersCollection:
		return c.UsersCollection.InsertOne(c.Context, document)
	case models.MessagesCollection:
		return c.MessagesCollection.InsertOne(c.Context, document)
	case models.ConversationsCollection:
		return c.ConversationsCollection.InsertOne(c.Context, document)
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
