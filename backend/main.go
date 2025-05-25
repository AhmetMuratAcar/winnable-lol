package main

import (
	"log"
	"net/http"
	"os"

	"github.com/AhmetMuratAcar/winnable-lol/internal/handlers"
	"github.com/joho/godotenv"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("Application error: %v", err)
	}
}

func run() error {
	// Load and set env variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or failed to load; using system environment")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default fallback
	}

	// Handlers
	mux := http.NewServeMux()
	mux.Handle("/health", &handlers.HealthHandler{})
	mux.Handle("/game", &handlers.GameHandler{})

	// Starting server
	log.Printf("Server is running on port %s...", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		return err
	}

	return nil
}
