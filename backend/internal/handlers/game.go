package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/AhmetMuratAcar/winnable-lol/backend/internal/riot"
)

type RequestBody struct {
	Region   string `json:"region"`
	Username string `json:"ign"`
}

type GameHandler struct{}

func (g *GameHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "405 Method Not Allowed: only POST is supported", http.StatusMethodNotAllowed)
		return
	}

	// Validate incoming JSON
	var req RequestBody
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if req.Username == "" {
		http.Error(w, "IGN is required", http.StatusBadRequest)
		return
	}

	if req.Region == "" {
		http.Error(w, "Region is required", http.StatusBadRequest)
		return
	}

	log.Printf("\nReceived IGN: %s\nReceived region: %s", req.Username, req.Region)

	// TODO: call Riot API client in internal/riot/client.go by passing req.
}
