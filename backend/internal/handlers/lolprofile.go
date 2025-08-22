package handlers

import (
	"context"
	"encoding/json"
	"errors"
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

	log.Println("----------------------------")
	log.Println("Received LoL profile request")
	var req types.RequestBody
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	log.Printf(
		"Received GameName: %s Tagline: %s Region: %s",
		req.GameName,
		req.TagLine,
		req.Region,
	)

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	// PUUID calls
	client := riot.NewClient()
	userProfile := types.LeagueProfilePage{
		GameName: req.GameName,
		TagLine:  req.TagLine,
		Region:   req.Region,
	}

	cacheCheck, err := utils.GetPUUID(ctx, h.pool, req)
	if err != nil {
		log.Printf(
			"Error querying DB for user's PUUID for GameName: %s Tagline: %s\nError: %v\n",
			req.GameName,
			req.TagLine,
			err,
		)
		// don't return and default back to riot API call for PUUID
	}

	if cacheCheck.Found {
		log.Print("user is cached, fetching data from DB")
		userProfile.PUUID = cacheCheck.PUUID

		if !cacheCheck.Stale {
			log.Print("data not stale, fetching from DB")
			// Fetch everything from DB and write to ResponseWriter if all data is present
			checklist, err := lolprofilesvc.CachedProfileConstructor(ctx, h.pool, &userProfile)
			if err != nil {
				log.Printf("Error calling CachedProfileConstructor: %v", err)
				// marking data as stale so consecutive calls dont lead to the same error
				// and instead default to Riot API calls
				err = utils.MarkSummonerStale(ctx, h.pool, userProfile.PUUID)
				if err != nil {
					log.Printf("Failed to mark summoner stale: %v", err)
				}

				http.Error(
					w,
					"internal server error",
					http.StatusInternalServerError,
				)
				return
			}

			err = lolprofilesvc.FillLoLProfileCacheGaps(
				checklist,
				&userProfile,
				client,
				ctx,
				h.pool,
			)
			if err != nil {
				log.Printf("Error calling FillLoLProfileCacheGaps: %v", err)
				// marking data as stale so consecutive calls dont lead to the same error
				// and instead default to Riot API calls
				err = utils.MarkSummonerStale(ctx, h.pool, userProfile.PUUID)
				if err != nil {
					log.Printf("Failed to mark summoner stale: %v", err)
				}

				http.Error(
					w,
					"internal server error",
					http.StatusInternalServerError,
				)
				return
			}

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
	} else {
		userProfile.PUUID, err = client.GetSummonerPUUID(req)
		if err != nil {
			var httpErr *types.HTTPError
			if errors.As(err, &httpErr) {
				log.Printf("GetSummonerPUUID HTTPError: status=%d err=%v", httpErr.StatusCode, err)
				http.Error(
					w,
					"could not fetch summoner",
					httpErr.StatusCode,
				)
			} else {
				log.Printf("GetSummonerPUUID internal error: %v", err)
				http.Error(
					w,
					"internal server error",
					http.StatusInternalServerError,
				)
			}

			return
		}
		// Only early terminating if PUUID fetch fails. If other client requests fail the userProfile
		// is constructed with any successfully received data.
	}
	log.Print("PUUID fetch successful")

	// Mastery Calls
	championMasteries, err := client.GetSummonerMastery(req.Region, userProfile.PUUID)
	if err != nil {
		log.Printf(
			"Error requesting masteries:\nPUUID:%s\nError: %v",
			userProfile.PUUID,
			err,
		)
	} else {
		log.Print("Mastery fetch successful")
	}

	for _, c := range championMasteries {
		userProfile.MasteryData.TotalMastery += c.ChampionLevel
		userProfile.MasteryData.TotalMasteryPoints += c.ChampionPoints
	}
	userProfile.MasteryData.ChampionsPlayed = len(championMasteries)
	userProfile.MasteryData.ChampionMasteries = championMasteries

	// Past matches calls
	startIndex := 0
	matchCount := 20
	matchIDs, err := client.GetSummonerMatchIDs(userProfile.PUUID, startIndex, matchCount)
	if err != nil {
		log.Printf(
			"Error requesting past match IDs: \nPUUID%s\nError: %v",
			userProfile.PUUID,
			err,
		)
	} else {
		log.Print("MatchIDs fetch successful")
	}

	matchDataMap := make(map[string]*types.LeagueMatch)
	if len(matchIDs) != 0 {
		matchDataMap, err = lolprofilesvc.ConstructMatchDataMap(ctx, h.pool, matchIDs)
		if err != nil {
			log.Printf(
				"Error constructing matchDataMap PUUID: %s\nmatchIDs: %s",
				userProfile.PUUID,
				matchIDs,
			)
			userProfile.MatchData = nil
		}
	} else {
		userProfile.MatchData = nil
	}

	// finding non-cached matches
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

	userProfile.MatchData = make([]types.LeagueMatch, 0, len(matchIDs))
	for _, id := range matchIDs {
		if m := matchDataMap[id]; m != nil {
			userProfile.MatchData = append(userProfile.MatchData, *m)
		}
	}
	log.Print("Match data successfully added")

	if len(userProfile.MatchData) > 0 {
		match := userProfile.MatchData[0]
		var userIndex int
		for i, v := range match.ParticipantPUUIDs {
			if v == userProfile.PUUID {
				userIndex = i
				break
			}
		}

		log.Println("getting icon and level from last match")
		userProfile.ProfileIconID = match.Participants[userIndex].ProfileIconID
		userProfile.Level = match.Participants[userIndex].SummonerLevel
	} else {
		userProfile.ProfileIconID, userProfile.Level, err = client.GetSummonerIconAndLevel(
			userProfile.PUUID,
			userProfile.Region,
		)

		if err != nil {
			log.Printf("error fetching summoner icon and level. Error: %v", err)
		} else {
			log.Print("Summoner icon and level successfully addded")
		}
	}

	// Rank Calls
	userProfile.Ranks, err = client.GetSummonerRanks(userProfile.PUUID, userProfile.Region)
	if err != nil {
		log.Printf("error fetching summoner ranks. Error: %v", err)
	} else {
		log.Print("Summoner ranks successfully added")
	}

	// Writing to file for dev
	if os.Getenv("ENV") == "development" {
		riotID := userProfile.GameName + "#" + userProfile.TagLine
		err := utils.WriteProfileToFile(userProfile, riotID)
		if err != nil {
			log.Printf("Failed to write profile to JSON. Error: %v", err)
		}
	}

	// Update DB with new data
	// updating matches table
	if len(toAdd) > 0 {
		detachedCtx, cancel := context.WithTimeout(
			context.WithoutCancel(ctx),
			5*time.Second,
		)
		defer cancel()

		go func(batch []types.LeagueMatch) {
			if err := utils.AddMatchData(detachedCtx, h.pool, batch); err != nil {
				log.Printf("async AddMatchData error: %v", err)
			}
		}(toAdd)
	}

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
