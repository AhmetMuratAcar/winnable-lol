package main

import (
	"log"
	"net/http"
	"os"

	"winnable/internal/handlers"
	"winnable/internal/middleware"
	"github.com/joho/godotenv"
)

func init() {
	// Load env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or failed to load; using system environment")
	} else {
		log.Println(".env file successfully loaded")
	}
}

func main() {
	if err := run(); err != nil {
		log.Fatalf("Application error: %v", err)
	}
}

func run() error {
	// Configure CORS
	middleware.ConfigureAllowedOrigins(os.Getenv("ENV"))

	// Handlers
	mux := http.NewServeMux()
	mux.Handle("/health", middleware.EnableCORS(&handlers.HealthHandler{}))
	mux.Handle("/mastery", middleware.EnableCORS(&handlers.MasteryHandler{}))
	// mux.Handle("/game", middleware.EnableCORS(&handlers.GameHandler{}))

	// Starting server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default fallback
	}

	log.Printf("Server is running on port %s...", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		return err
	}

	return nil
}
