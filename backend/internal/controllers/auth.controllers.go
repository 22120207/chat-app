package controllers

import (
	"chat-app-backend/internal/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthSignup(c *gin.Context) {
	var json models.SingupRequest

	if err := c.ShouldBindBodyWithJSON(&json); err != nil {
		log.Printf("Error in bind Signup Request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Signup Request is missing fields"})
		return
	}

	if json.Password != json.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password don't match"})
		return
	}
}
