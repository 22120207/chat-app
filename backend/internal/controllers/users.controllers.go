package controllers

import (
	"chat-app-backend/internal/models"
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func GetUsers(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.Error(errors.New(string(models.UnauthorizedError)))
		return
	}

	loggedInUser := user.(models.User)

	var otherUsers []models.User
	cursor, err := Client.Find(models.UsersCollection, bson.M{
		"_id": bson.M{"$ne": loggedInUser.ID},
	})
	if err != nil {
		c.Error(errors.New(string(models.InternalServerError)))
		return
	}
	defer cursor.Close(context.Background())

	if err = cursor.All(context.Background(), &otherUsers); err != nil {
		c.Error(errors.New(string(models.InternalServerError)))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"loggedInUserId": loggedInUser.ID.Hex(),
		"otherUsers":     otherUsers,
	})
}
