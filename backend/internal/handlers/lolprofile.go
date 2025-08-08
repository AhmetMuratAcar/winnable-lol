package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"winnable/internal/riot"
	"winnable/internal/types"
	"winnable/internal/utils"

	"github.com/jackc/pgx/v5/pgxpool"
)

type LoLProfileHandler struct {
	pool *pgxpool.Pool
}

func NewLoLProfileHandler(pool *pgxpool.Pool) *LoLProfileHandler {
	return &LoLProfileHandler{pool: pool}
}

func (h *LoLProfileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "405 Method Not Allowed: only POST is supported", http.StatusMethodNotAllowed)
		return
	}

	log.Println("Received LoL profile request")
	var req types.RequestBody
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	log.Printf("\nReceived GameName: %s Tagline: %s Region: %s", req.GameName, req.TagLine, req.Region)

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	// PUUID calls
	var PUUID string
	cacheCheck, err := utils.GetPUUID(ctx, h.pool, req)
	if err != nil {
		log.Printf(
			"Error querying DB for user's PUUID for GameName: %s Tagline: %s\nError: %v\n",
			req.GameName,
			req.TagLine,
			err,
		)
		// don't return and default back to riot API call for PUUID
		cacheCheck.Found = false
	}

	client := riot.NewClient()
	if !cacheCheck.Found {
		PUUID, err = client.GetSummonerPUUID(req)
		if err != nil {
			http.Error(
				w,
				"could not fetch summoner: "+err.Error(),
				http.StatusNotFound,
			)
			return
		}
	} else {
		PUUID = cacheCheck.PUUID
	}

	userProfile := types.LeagueProfilePage{
		GameName: req.GameName,
		TagLine: req.TagLine,
		Region: req.Region,
	}

	// Mastery Calls
	var championMasteries []types.ChampionMastery
	if cacheCheck.Stale {
		championMasteries, err = client.GetSummonerMastery(req.Region, PUUID)
		if err != nil {
			log.Printf(
				"Error requesting masteries:\nPUUID:%s\nError: %v",
				PUUID,
				err,
			)
		}
	} else {
		championMasteries, err = utils.GetMasteries(ctx, h.pool, PUUID)
		if err != nil {
			log.Printf(
				"Error querying DB for user's masteries PUUID: %s\nError: %v",
				PUUID,
				err,
			)
		}
	}
	userProfile.MasteryData.ChampionMasteries = championMasteries

	// Past matches calls

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(userProfile); err != nil {
		log.Printf("failed to encode user's profile data: %v", err)
		http.Error(
			w,
			"internal server error",
			http.StatusInternalServerError,
		)
		return
	}
}
