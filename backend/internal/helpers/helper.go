package helpers

import (
	"chat-app-backend/internal/config"
	"chat-app-backend/internal/models"
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

func CreateToken(c *gin.Context, user models.User, secretKey string) error {
	log.Println(user.ID.Hex())

	// Create a new JWT token with claims
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID.Hex(),                        // Subject (user identifier)
		"iss": config.Conf.Database,                 // Issuer
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
func VerifyToken(tokenString string, secretKey string) (*jwt.Token, jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		return nil, nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, nil, fmt.Errorf("cannot parse claims")
	}

	return token, claims, nil
}