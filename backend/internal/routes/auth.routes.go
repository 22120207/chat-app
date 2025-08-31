package routes

import (
	"chat-app-backend/internal/controllers"
)

func (api *Routes) setUpAuthRoutes() {
	apiAuth := api.routerGroup.Group("/auth")
	apiAuth.POST("/signup", controllers.AuthSignup)
	apiAuth.POST("/login", controllers.AuthLogin)
	apiAuth.POST("/logout", controllers.AuthLogout)
}
