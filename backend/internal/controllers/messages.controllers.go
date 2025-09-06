package controllers

import (
	"chat-app-backend/internal/models"
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func CreateMessage(senderID, receiverID bson.ObjectID, text string) (*models.Message, error) {
	newMessage := models.Message{
		ID:         bson.NewObjectID(),
		SenderID:   senderID,
		ReceiverID: receiverID,
		Message:    text,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	// Check if conversation exists
	var conversation models.Conversation
	err := Client.FindOne(
		models.ConversationsCollection,
		bson.M{"participants": bson.M{"$all": []bson.ObjectID{senderID, receiverID}}},
	).Decode(&conversation)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			// Create new conversation
			conversation = models.Conversation{
				ID:           bson.NewObjectID(),
				Participants: []bson.ObjectID{senderID, receiverID},
				Messages:     []bson.ObjectID{newMessage.ID},
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
			}
			if _, err := Client.InsertOne(models.ConversationsCollection, conversation); err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	} else {
		// Update existing conversation
		_, err = Client.UpdateOne(
			models.ConversationsCollection,
			bson.M{"_id": conversation.ID},
			bson.M{
				"$push": bson.M{"messages": newMessage.ID},
				"$set":  bson.M{"updatedAt": time.Now()},
			},
		)
		if err != nil {
			return nil, err
		}
	}

	// Insert new message
	if _, err := Client.InsertOne(models.MessagesCollection, newMessage); err != nil {
		return nil, err
	}

	return &newMessage, nil
}

func SendMessage(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.Error(errors.New(string(models.UnauthorizedError)))
		return
	}
	sender := user.(models.User)

	var json models.MessageRequest
	if err := c.ShouldBindBodyWithJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Message Request is not correct"})
		return
	}

	receiverId := c.Param("id")
	receiverObjID, err := bson.ObjectIDFromHex(receiverId)
	if err != nil {
		c.Error(errors.New(string(models.InternalServerError)))
		return
	}

	msg, err := CreateMessage(sender.ID, receiverObjID, json.Message)
	if err != nil {
		c.Error(errors.New(string(models.InternalServerError)))
		return
	}

	c.JSON(http.StatusCreated, msg)
}

// Get all messages of the person that the user chat with
func GetMessage(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.Error(errors.New(string(models.UnauthorizedError)))
		return
	}

	sender := user.(models.User)

	receiverId := c.Param("id")
	receiverObjID, err := bson.ObjectIDFromHex(receiverId)
	if err != nil {
		c.Error(errors.New(string(models.InternalServerError)))
		return
	}

	var conversation models.Conversation
	err = Client.FindOne(models.ConversationsCollection, bson.M{
		"participants": bson.M{"$all": []bson.ObjectID{sender.ID, receiverObjID}},
	}).Decode(&conversation)

	if err == mongo.ErrNoDocuments {
		c.JSON(http.StatusOK, gin.H{})
		return
	}

	if err != nil {
		c.Error(errors.New(string(models.InternalServerError)))
		return
	}

	var messages []models.Message
	cursor, err := Client.Find(models.MessagesCollection, bson.M{
		"_id": bson.M{"$in": conversation.Messages},
	})
	if err != nil {
		c.Error(errors.New(string(models.InternalServerError)))
		return
	}
	defer cursor.Close(context.Background())

	if err = cursor.All(context.Background(), &messages); err != nil {
		c.Error(errors.New(string(models.InternalServerError)))
		return
	}

	c.JSON(http.StatusOK, messages)
}
