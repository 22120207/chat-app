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

func AuthMiddleware(c *gin.Context) {
	config.LoadEnv()
	conf := config.Conf

	// Retrieve the token from the cookie
	tokenString, err := c.Cookie("token")
	if err != nil {
		log.Printf("Token missing in cookie: %v", err)
		c.Error(errors.New(string(models.UnauthorizedError)))
		c.Abort()
		return
	}

	// Verify the token
	_, claims, err := helpers.VerifyToken(tokenString, conf.JwtSecret)
	if err != nil {
		c.Error(errors.New(string(models.UnauthorizedError)))
		return
	}

	userId := claims["sub"]

	objectID, err := bson.ObjectIDFromHex(userId.(string))
	if err != nil {
		c.Error(errors.New(string(models.InternalServerError)))
		return
	}

	var user models.User
	err = controllers.Client.FindOne(models.UsersCollection, bson.M{
		"_id": objectID,
	}).Decode(&user)

	if err != nil {
		c.Error(errors.New(string(models.UnauthorizedError)))
		return
	}

	c.Set("user", user)

	// Continue with the next middleware or route handler
	c.Next()
}
