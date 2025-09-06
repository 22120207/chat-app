package routes

import (
	"chat-app-backend/internal/middleware"
	"chat-app-backend/internal/websocket"

	"github.com/gin-gonic/gin"
)

func (api *Routes) setUpWebsocketRoutes() {
	hub := websocket.NewHub()
	go hub.Run()

	apiAuth := api.routerGroup.Group("/ws")
	apiAuth.GET("", middleware.AuthMiddleware, func(c *gin.Context) {
		websocket.ServeWs(hub, c)
	})
}
