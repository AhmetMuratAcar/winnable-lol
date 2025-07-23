package handlers

import (
	"encoding/json"
	"fmt"
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
		puuid, err = client.GetSummonerPUUID(req)
		if err != nil {
			http.Error(
				w,
				"could not fetch summoner: "+err.Error(),
				http.StatusNotFound,
			)
			return
		}
	}

	log.Printf("PUUID: %s", puuid)

	// Mastery calls
	isMasteryOutdated := true
	var championMasteries []types.ChampionMastery
	if isUserCached {
		// Check DB for last mastery fetch timestamp
		// If under 24 Hours, fetch from DB and set isMasteryOutdated to false
	}

	if isMasteryOutdated {
		championMasteries, err = client.GetSummonerMastery(req.Region, puuid)
		if err != nil {
			http.Error(
				w,
				"could not fetch summoner's masteries: "+err.Error(),
				http.StatusNotFound,
			)
			return
		}
	}

	if len(championMasteries) == 0 {
		// User hasn't played any games
		// Return an empty response
	}

	// Writing to file for now
	err = utils.WriteMasteryToFile(championMasteries, req.GameName+req.TagLine)
	if err != nil {
		fmt.Println("we fucked up gang")
	}

	graphData, err := utils.ProcessMasteryData(championMasteries)
	if err != nil {
		http.Error(
			w,
			"could not generate radar graph: "+err.Error(),
			http.StatusInternalServerError,
		)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(graphData); err != nil {
		log.Printf("failed to encode graph data: %v", err)
		http.Error(
			w,
			"internal server error",
			http.StatusInternalServerError,
		)
		return
	}
}
