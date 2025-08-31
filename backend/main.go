package main

import (
	"chat-app-backend/internal/routes"
	"log"
)

func main() {
	// Setup Router
	r := routes.SetupRouter()

	if err := r.Run(":5000"); err != nil {
		log.Printf("Chat-app Backend server failed: %v", err)
		return
	}
}
