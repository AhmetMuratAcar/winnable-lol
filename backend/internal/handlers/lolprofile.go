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
	log.Print("PUUID fetch successful")

	userProfile := types.LeagueProfilePage{
		GameName: req.GameName,
		TagLine:  req.TagLine,
		Region:   req.Region,
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
	log.Print("Mastery fetch successful")
	userProfile.MasteryData.ChampionMasteries = championMasteries

	// Past matches calls
	startIndex := 0
	var matchIDs []string
	matchIDs, err = client.GetSummonerMatchIDs(PUUID, startIndex)
	if err != nil {
		log.Printf(
			"Error requesting past match IDs: \nPUUID%s\nError: %v",
			PUUID,
			err,
		)
	}
	log.Print("MatchIDs fetch successful")
	
	matchDataMap := make(map[string]*types.LeagueMatch)
	if len(matchIDs) != 0 {
		err = utils.GetMatchDataByIDs(ctx, h.pool, matchIDs, &matchDataMap)
		if err != nil {
			log.Printf(
				"Error populating matchDataMap PUUID: %s\nmatchIDs: %s",
				PUUID,
				matchIDs,
			)
			userProfile.MatchData = nil
		}
	} else {
		userProfile.MatchData = nil
	}

	missing := make([]string, 0, len(matchIDs))
	for _, id := range matchIDs {
		if m, ok := matchDataMap[id]; !ok || m == nil {
			missing = append(missing, id)
		}
	}

	toAdd := make([]types.LeagueMatch, 0, len(missing))
	for _, id := range missing {
		matchData, err := client.GetMatchData(id)
		if err != nil {
			log.Printf(
				"Error fetching matchID %s\nError: %v",
				id,
				err,
			)
			continue
		}

		matchDataMap[id] = &matchData
		toAdd = append(toAdd, matchData)
	}

	// updating matches table
	if len(toAdd) > 0 {
		detachedCtx, cancel := context.WithTimeout(
			context.WithoutCancel(ctx), 
			5*time.Second,
		)
		defer cancel()

		go func (batch []types.LeagueMatch)  {
			if err := utils.AddMatchData(detachedCtx, h.pool, batch); err != nil {
				log.Printf("async AddMatchData error: %v", err)
			} 
		}(toAdd)
	}

	userProfile.MatchData = make([]types.LeagueMatch, 0, len(matchIDs))
	for _, id := range matchIDs {
		if m := matchDataMap[id]; m != nil {
			userProfile.MatchData = append(userProfile.MatchData, *m)
		}
	}
	log.Print("Match data successfully added")
	
	// Remember to set userProfile's icon and level from the data of the games
	// If there are no games default to a riot API call for that data.

	// Rank Calls

	// Writing to file for dev
	riotID := userProfile.GameName + userProfile.TagLine
	err = utils.WriteProfileToFile(userProfile, riotID)
	if err != nil {
		log.Printf("Failed to write profile to JSON. Error: %v", err)
	}

	// Update the user's information in the databse if it was stale or maybe
	// even if it was not

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
