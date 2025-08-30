package routes

import (
	"chat-app-backend/internal/controllers"
)

func (api *Routes) setUpAuthRoutes() {
	api.routerGroup.POST("/auth", controllers.AuthSignup)
}
