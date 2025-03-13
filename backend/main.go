package main

import (
	"log"
	"net/http"

	"github.com/AhmetMuratAcar/winnable-lol/internal/handlers"
)

func main() {
	const PORT = "8080"

	// Handlers
	http.HandleFunc("/current-game", handlers.CurrentGameHandler)

	// Starting server
	log.Printf("Server is running on port %s...", PORT)
	if err := http.ListenAndServe(":"+PORT, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
