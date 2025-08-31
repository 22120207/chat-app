package routes

import (
	"chat-app-backend/internal/controllers"
)

func (api *Routes) setUpMessagesRoutes() {
	apiAuth := api.routerGroup.Group("/messages")
	apiAuth.POST("/send/:id", controllers.AuthMiddleware, controllers.SendMessage)
}
