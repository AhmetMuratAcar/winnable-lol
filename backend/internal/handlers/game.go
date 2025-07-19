package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	// "github.com/AhmetMuratAcar/winnable-lol/backend/internal/riot"
	"github.com/AhmetMuratAcar/winnable-lol/backend/internal/types"

)

type GameHandler struct{}

func (g *GameHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "405 Method Not Allowed: only POST is supported", http.StatusMethodNotAllowed)
		return
	}

	log.Print("Received a request")
	// Validate incoming JSON
	var req types.RequestBody
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

	// Riot API calls
	// client := riot.NewClient()
	// client.GetMatchData()
	// user, err := client.GetSummoner(req.Region, req.Username)
	// if err != nil {
	// 	http.Error(w, "could not fetch summoner: "+err.Error(), http.StatusInternalServerError)
	// 	return
	// }
}
