package controllers

import (
	"chat-app-backend/internal/models"
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func SendMessage(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.Error(errors.New(string(models.UnauthorizedError)))
		return
	}

	sender := user.(models.User)

	var json models.MessageRequest

	if err := c.ShouldBindBodyWithJSON(&json); err != nil {
		log.Printf("Error in bind Message Request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Message Request is not correct"})
		return
	}

	receiverId := c.Param("id")
	receiverObjID, err := bson.ObjectIDFromHex(receiverId)
	if err != nil {
		c.Error(errors.New(string(models.InternalServerError)))
		return
	}

	newMessage := models.Message{
		ID:         bson.NewObjectID(),
		SenderID:   sender.ID,
		ReceiverID: receiverObjID,
		Message:    json.Message,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	// Check if the conversation is exists
	var conversation models.Conversation
	err = Client.FindOne(models.ConversationsCollection, bson.M{
		"participants": bson.M{"$all": []bson.ObjectID{sender.ID, receiverObjID}},
	}).Decode(&conversation)

	if err != nil {
		// If there is not conversation match --> Create a new conversation
		if err == mongo.ErrNoDocuments {
			conversation = models.Conversation{
				ID:           bson.NewObjectID(),
				Participants: []bson.ObjectID{sender.ID, receiverObjID},
				Messages:     []bson.ObjectID{newMessage.ID},
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
			}

			_, err = Client.InsertOne(models.ConversationsCollection, conversation)
			if err != nil {
				c.Error(errors.New(string(models.InternalServerError)))
				return
			}
		} else {
			c.Error(errors.New(string(models.InternalServerError)))
		}
	} else {
		// If conversation exists --> update it
		_, err = Client.UpdateOne(
			models.ConversationsCollection,
			bson.M{"_id": conversation.ID},
			bson.M{
				"$push": bson.M{"messages": newMessage.ID},
				"$set":  bson.M{"updatedAt": time.Now()},
			},
		)
		if err != nil {
			c.Error(errors.New(string(models.InternalServerError)))
			return
		}
	}

	// Insert message into messages collection (optional if you keep it embedded too)
	_, err = Client.InsertOne(models.MessagesCollection, newMessage)
	if err != nil {
		c.Error(errors.New(string(models.InternalServerError)))
		return
	}

	c.JSON(http.StatusCreated, newMessage)
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
