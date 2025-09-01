package middleware

import (
	"chat-app-backend/internal/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorHandlerMiddleware handles errors added to the Gin context.
func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Check if any errors were added to the context
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err

			log.Printf("Request error: %v", err)

			// Don't send response if already sent
			if c.Writer.Written() {
				return
			}

			statusCode := http.StatusInternalServerError
			message := err.Error()

			switch message {
			case string(models.UnauthorizedError):
				statusCode = http.StatusUnauthorized
			case string(models.InternalServerError):
				statusCode = http.StatusInternalServerError
			}

			c.JSON(statusCode, gin.H{
				"error": message,
			})
		}
	}
}
