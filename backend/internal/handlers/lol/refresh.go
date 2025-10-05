package lol

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
	"winnable/internal/lolprofilesvc"
	"winnable/internal/riot"
	"winnable/internal/types"
	"winnable/internal/utils"

	"github.com/jackc/pgx/v5/pgxpool"
)

type LolRefreshHandler struct {
	pool *pgxpool.Pool
}

func NewLoLRefreshHandler(pool *pgxpool.Pool) *LolRefreshHandler {
	return &LolRefreshHandler{pool: pool}
}

func (h *LolRefreshHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "405 Method Not Allowed: only GET is supported", http.StatusMethodNotAllowed)
		return
	}

	log.Println("----------------------------")
	log.Println("Received LoL profile refresh request")

	defer r.Body.Close()
	req, ok := utils.ValidateLeagueProfileReq(w, r)
	if !ok {
		log.Print("Bad request, missing query params")
		return
	}

	log.Printf(
		"Received GameName: %s Tagline: %s Region: %s",
		req.GameName,
		req.TagLine,
		req.Region,
	)

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	cacheCheck, err := utils.SummonerCacheCheck(ctx, h.pool, req)
	if err != nil {
		log.Printf("Error calling SummonerCacheCheck in /lol/refresh: %v", err)
		http.Error(
			w,
			"internal server error",
			http.StatusInternalServerError,
		)
		return
	}

	if !cacheCheck.Found {
		// In the future this case can redirect to /lol/profile for the data maybe?
		log.Printf("User not cached")
		http.Error(
			w,
			"summoner not found",
			http.StatusNotFound,
		)
	}

	// The rest of this endpoint fetches all profile data from Riot
	client := riot.NewClient()
	userProfile := types.LeagueProfilePage{
		PUUID:       cacheCheck.PUUID,
		GameName:    req.GameName,
		TagLine:     req.TagLine,
		Region:      req.Region,
		LastUpdated: time.Now(),
	}

	userProfile.Ranks, err = client.GetSummonerRanks(userProfile.PUUID, userProfile.Region)
	if err != nil {
		log.Printf("Error fetching summoner ranks: %v", err)
		userProfile.Ranks = []types.LeagueRank{}
	} else {
		log.Print("Ranks successfully fetched in /lol/refresh")
	}

	startIndex := 0
	matchCount := 20
	matchIDs, err := client.GetSummonerMatchIDs(userProfile.PUUID, startIndex, matchCount)
	if err != nil {
		log.Printf("Error requesting matchIDs from Riot in /lol/refresh: %v", err)
	} else {
		log.Printf("MatchIDs successfully fetched in /lol/refresh")
	}

	matchDataMap, err := lolprofilesvc.ConstructMatchDataMap(ctx, h.pool, matchIDs)
	if err != nil {
		log.Printf("Error constructing matchDataMap in /lol/refresh: %v", err)
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
				"Error fetching matchID %s in /lol/refresh: %v",
				id,
				err,
			)
			continue
		}

		matchDataMap[id] = &matchData
		toAdd = append(toAdd, matchData)
	}

	userProfile.MatchData = make([]types.LeagueMatch, 0, len(matchIDs))
	for _, id := range matchIDs {
		if m := matchDataMap[id]; m != nil {
			userProfile.MatchData = append(userProfile.MatchData, *m)
		}
	}
	log.Print("Match data successfully added in /lol/refresh")

	userProfile.ProfileIconID, userProfile.Level, err = client.GetSummonerIconAndLevel(
		userProfile.PUUID,
		userProfile.Region,
	)
	if err != nil {
		log.Printf("Error fetching summoner icon and level in /lol/refresh: %v", err)
	} else {
		log.Print("Summoner icon and level successfully fetched in /lol/refresh")
	}

	userProfile.PlayedWith, userProfile.PlayedAgainst, err = lolprofilesvc.ConstructRecentlyPlayedWithAndAgainst(userProfile.MatchData, userProfile.PUUID)
	if err != nil {
		log.Printf("error constructing recently played with in /lol/refresh: %v", err)
	} else {
		log.Print("recently played with constructed in /lol/refresh")
	}

	userProfile.RecentGames, err = lolprofilesvc.ConstructGamesSummary(userProfile.MatchData, userProfile.PUUID)
	if err != nil {
		log.Printf("error constructing games summary in /lol/refresh: %v", err)
	} else {
		log.Print("games summary constructed in /lol/refresh")
	}

	if os.Getenv("ENV") == "development" {
		riotID := userProfile.GameName + "#" + userProfile.TagLine
		err := utils.WriteRefreshToFile(userProfile, riotID)
		if err != nil {
			log.Printf("Failed to write profile to JSON. Error: %v", err)
		}
	}

	// Updating DB
	// TODO: add separate Go routine for updating DB

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(userProfile); err != nil {
		log.Printf("Failed to encode user's profile data in /lol/refresh: %v", err)
		http.Error(
			w,
			"internal server error",
			http.StatusInternalServerError,
		)
		return
	}
}
