package main

import (
	"chat-app-backend/internal/controllers"
	"chat-app-backend/internal/routes"
	"log"
)

func main() {
	// Setup Router
	r := routes.SetupRouter()

	// defer disconnect connection to MongoDB
	defer controllers.Client.Disconnect()

	if err := r.Run("localhost:5000"); err != nil {
		log.Printf("Chat-app Backend server failed: %v", err)
		return
	}
}
