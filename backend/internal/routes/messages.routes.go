package routes

import (
	"chat-app-backend/internal/controllers"
	"chat-app-backend/internal/middleware"
)

func (api *Routes) setUpMessagesRoutes() {
	apiAuth := api.routerGroup.Group("/messages")
	apiAuth.POST("/send/:id", middleware.AuthMiddleware, controllers.SendMessage)
	apiAuth.GET("/:id", middleware.AuthMiddleware, controllers.GetMessage)
}
