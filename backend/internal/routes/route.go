package routes

import (
	"chat-app-backend/internal/middleware"

	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
)

type Routes struct {
	routerGroup *gin.RouterGroup
}

func SetupRouter() *gin.Engine {

	r := gin.Default()

	// Add a logger middleware, which logs all requests, like a combined access and error log
	r.Use(logger.SetLogger())

	r.Use(middleware.ErrorHandlerMiddleware())

	r.Use(func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		if origin == "http://localhost:5173" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		}

		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, UPDATE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	})

	api := Routes{
		routerGroup: r.Group("/api/"),
	}

	api.setUpAuthRoutes()
	api.setUpMessagesRoutes()
	api.setUpUsersRoutes()

	return r
}
