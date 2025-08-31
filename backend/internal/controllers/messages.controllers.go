package controllers

import (
	"chat-app-backend/internal/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendMessage(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.Error(errors.New(string(models.UnauthorizedError)))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Send messages successfully",
		"user":    user.(models.User),
	})
}
