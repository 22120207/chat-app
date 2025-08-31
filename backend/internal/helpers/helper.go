package helpers

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func ToPrefixName(fullname string) string {
	fullname = strings.TrimSpace(fullname)
	return strings.ReplaceAll(fullname, " ", "+")
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CreateToken(c *gin.Context, username, secretKey string) error {
	// Create a new JWT token with claims
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": username,                             // Subject (user identifier)
		"iss": "chat-app",                           // Issuer
		"exp": time.Now().Add(8 * time.Hour).Unix(), // Expiration time
		"iat": time.Now().Unix(),                    // Issued at
	})

	tokenString, err := claims.SignedString([]byte(secretKey))
	if err != nil {
		log.Printf("Error in create token string: %v", err)
		return err
	}

	c.SetCookie("token", tokenString, 8*60*60, "/", "localhost", false, true)

	return nil
}

// Function to verify JWT tokens
func VerifyToken(tokenString, secretKey string) error {
	// Parse the token with the secret key
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		log.Printf("Error in parse jwt token: %v", err)
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
