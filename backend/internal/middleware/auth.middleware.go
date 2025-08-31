package middleware

import (
	"chat-app-backend/internal/config"
	"chat-app-backend/internal/controllers"
	"chat-app-backend/internal/helpers"
	"chat-app-backend/internal/models"
	"errors"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

var conf = config.Conf

func AuthMiddleware(c *gin.Context) {
	// Retrieve the token from the cookie
	tokenString, err := c.Cookie("token")
	if err != nil {
		log.Printf("Token missing in cookie: %v", err)
		c.Abort()
		return
	}

	// Verify the token
	_, claims, err := helpers.VerifyToken(tokenString, conf.JwtSecret)
	if err != nil {
		log.Printf("Token verification failed: %v", err)
		c.Error(errors.New(string(models.UnauthorizedError)))
		return
	}

	userId := claims["sub"]

	var user models.User
	err = controllers.Client.FindOne(models.UsersCollection, bson.M{
		"_id": userId,
	}).Decode(&user)

	if err != nil {
		log.Printf("Error in find user %v", err)
		c.Error(errors.New(string(models.UnauthorizedError)))
		return
	}

	c.Set("user", user)

	// Continue with the next middleware or route handler
	c.Next()
}
