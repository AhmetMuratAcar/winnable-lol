package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type RequestBody struct {
	Username string `json:"username"`
}

func CurrentGameHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	var req RequestBody
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON request body", http.StatusBadRequest)
		return
	}

	// TODO: move username formatting to frontend.
	// Username#Tag validation within text box.
	if req.Username == "" {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}

	log.Printf("Received username: %s", req.Username)
	
	// TODO: call Riot API client in internal/riotapi/client.go by passing username.
}
