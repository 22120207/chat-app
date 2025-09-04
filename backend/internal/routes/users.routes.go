package routes

import (
	"chat-app-backend/internal/controllers"
	"chat-app-backend/internal/middleware"
)

func (api *Routes) setUpUsersRoutes() {
	apiAuth := api.routerGroup.Group("/users")
	apiAuth.GET("", middleware.AuthMiddleware, controllers.GetUsers)
}
