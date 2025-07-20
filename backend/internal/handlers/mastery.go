package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/AhmetMuratAcar/winnable-lol/backend/internal/riot"
	"github.com/AhmetMuratAcar/winnable-lol/backend/internal/types"
)

type MasteryHandler struct{}

func (h *MasteryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "405 Method Not Allowed: only POST is supported", http.StatusMethodNotAllowed)
		return
	}

	log.Println("Received mastery request")
	var req types.RequestBody
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if req.GameName == "" {
		http.Error(w, "GameName is required", http.StatusBadRequest)
		return
	}
	
	if req.TagLine == "" {
		http.Error(w, "Tagline is required", http.StatusBadRequest)
		return
	}

	if req.Region == "" {
		http.Error(w, "Region is required", http.StatusBadRequest)
		return
	}

	log.Printf("\nReceived GameName: %s Tagline: %s Region: %s", req.GameName, req.TagLine, req.Region)
	
	client := riot.NewClient()
	puuid, err := client.GetSummonerPUUID(req)
	if err != nil {
		http.Error(
			w, 
			"could not fetch summoner: "+err.Error(), 
			http.StatusNotFound,
		)
		return
	}
	log.Printf("PUUID: %s", puuid)
	// Call riot mastery handler

	// Call mastery processor
}
