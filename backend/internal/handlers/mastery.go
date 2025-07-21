package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"winnable/internal/riot"
	"winnable/internal/types"
	"winnable/internal/utils"
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
	log.Printf("\nReceived GameName: %s Tagline: %s Region: %s", req.GameName, req.TagLine, req.Region)
	
	// PUUID calls
	client := riot.NewClient()
	isUserCached, puuid, err := utils.GetPUIID(req)
	if err != nil {
		log.Printf(
			"Error querying DB for user's PUUID for GameName: %s Tagline: %s\nerr: %v\n", 
			req.GameName, 
			req.TagLine,
			err,
		)
		// don't return and default back to riot API call for PUUID
		isUserCached = false
	}

	if !isUserCached {
		tmp, err := client.GetSummonerPUUID(req)
		if err != nil {
			http.Error(
				w, 
				"could not fetch summoner: "+err.Error(), 
				http.StatusNotFound,
			)
			return
		}
		puuid = tmp
	}

	log.Printf("PUUID: %s", puuid)
	// Mastery calls
	// DB Check for user mastery info
	// Riot mastery API
	// Call mastery processor
}
