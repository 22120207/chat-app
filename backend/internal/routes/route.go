package routes

import (
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

	r.Use(func(c *gin.Context) {

		// Set HTTP response headers for all requests
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	})

	api := Routes{
		routerGroup: r.Group("/api/"),
	}

	api.setUpAuthRoutes()

	return r
}
