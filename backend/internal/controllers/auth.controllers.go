package controllers

import (
	"chat-app-backend/internal/helpers"
	"chat-app-backend/internal/models"
	"errors"
	"log"
	"net/http"
	"time"

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
	err := Client.FindOne(models.UsersCollection, bson.M{
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
		c.Error(errors.New(string(models.InternalServerError)))
		return
	}

	newUser := models.User{
		Fullname:   json.Fullname,
		Username:   json.Username,
		Password:   hashPassword,
		Gender:     json.Gender,
		ProfilePic: profilePic,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	result, err := Client.InsertOne(models.UsersCollection, newUser)
	if err != nil {
		log.Printf("Error in insert document: %v", err)
		c.Error(errors.New(string(models.InternalServerError)))
		return
	}

	// Create Token
	err = helpers.CreateToken(c, newUser, conf.JwtSecret)
	if err != nil {
		c.Error(errors.New(string(models.InternalServerError)))
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
	err := Client.FindOne(models.UsersCollection, bson.M{
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
	err = helpers.CreateToken(c, user, conf.JwtSecret)
	if err != nil {
		c.Error(errors.New(string(models.InternalServerError)))
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
