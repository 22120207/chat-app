package controllers

import (
	"chat-app-backend/internal/helpers"
	"chat-app-backend/internal/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func AuthSignup(c *gin.Context) {
	var json models.SingupRequest

	if err := c.ShouldBindBodyWithJSON(&json); err != nil {
		log.Printf("Error in bind Signup Request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Signup Request is not correct"})
		return
	}

	if json.Password != json.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password don't match"})
		return
	}

	var user models.User
	err := Client.FindOne(string(UserCollection), bson.M{
		"username": json.Username,
	}).Decode(&user)

	if err != nil {
		if err != mongo.ErrNoDocuments {
			log.Panic(err)
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username must be unique"})
		return
	}

	profilePic := `https://avatar.iran.liara.run/username?username=` + helpers.ToPrefixName(json.Fullname)

	hashPassword, err := helpers.HashPassword(json.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	newUser := models.User{
		Fullname:   json.Fullname,
		Username:   json.Username,
		Password:   hashPassword,
		Gender:     json.Gender,
		ProfilePic: profilePic,
	}

	result, err := Client.InsertOne(string(UserCollection), newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Create Token
	err = helpers.CreateToken(c, json.Username, conf.JwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"id":      result.InsertedID,
	})
}

func AuthLogin(c *gin.Context) {
	var json models.LoginRequest

	if err := c.ShouldBindBodyWithJSON(&json); err != nil {
		log.Printf("Error in bind Login Request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Login Request is not correct"})
		return
	}

	var user models.User
	err := Client.FindOne(string(UserCollection), bson.M{
		"username": json.Username,
	}).Decode(&user)

	if err != nil {
		log.Printf("Error in find user %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Username/Password is not correct",
		})
		return
	}

	// If password is not correct
	if !helpers.CheckPasswordHash(json.Password, user.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Username/Password is not correct",
		})
		return
	}

	// Create Token
	err = helpers.CreateToken(c, json.Username, conf.JwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successfully",
		"id":      user.ID,
	})
}

func AuthLogout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Logout successfully",
	})
}

func AuthMiddleware(c *gin.Context) {
	// Retrieve the token from the cookie
	tokenString, err := c.Cookie("token")
	if err != nil {
		log.Printf("Token missing in cookie: %v", err)
		c.Abort()
		return
	}

	// Verify the token
	err = helpers.VerifyToken(tokenString, conf.JwtSecret)
	if err != nil {
		log.Printf("Token verification failed: %v\\n", err)
		c.Abort()
		return
	}

	// Continue with the next middleware or route handler
	c.Next()
}
