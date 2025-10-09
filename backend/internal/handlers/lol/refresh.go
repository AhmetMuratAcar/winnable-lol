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
	matchIDs, err := client.GetSummonerMatchIDs(userProfile.PUUID, userProfile.Region, startIndex, matchCount)
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
		matchData, err := client.GetMatchData(id, userProfile.Region)
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
			log.Printf("Failed to write profile refresh to JSON. Error: %v", err)
		}
	}

	// Updating DB
	pool := h.pool
	checkCopy := cacheCheck
	profileCopy := userProfile
	toAddCopy := append([]types.LeagueMatch(nil), toAdd...)

	go func(
		parent context.Context,
		pool *pgxpool.Pool,
		check types.PUUIDCacheCheck,
		profile types.LeagueProfilePage,
		newMatches []types.LeagueMatch,
	) {
		log.Print("---START REFRESH POST-PROCESSING---")
		ctx, cancel := context.WithTimeout(context.WithoutCancel(parent), 30*time.Second)
		defer cancel()

		start := time.Now()
		var err error

		// mastery calls
		if !check.Found || !check.IsPopulated || check.Stale {
			safeClient := riot.NewClient()
			profile.MasteryData.ChampionMasteries, err = safeClient.GetSummonerMastery(
				profile.Region,
				profile.PUUID,
			)
			if err != nil {
				log.Printf("Error calling GetSummonerMastery in /lol/refresh: %v", err)
			} else {
				log.Print("Masteries successfully fetched from Riot")
			}

			for _, c := range profile.MasteryData.ChampionMasteries {
				profile.MasteryData.TotalMastery += c.ChampionPoints
				profile.MasteryData.TotalMasteryPoints += c.ChampionPoints
			}
			profile.MasteryData.ChampionsPlayed = len(profile.MasteryData.ChampionMasteries)
		}

		// update/add current user's data
		if err := utils.SyncSummonerProfileData(ctx, pool, check, profile); err != nil {
			log.Printf("SyncProfileData error in /lol/refresh: %v", err)
		} else {
			log.Print("Profile data synced")
		}

		// add all new PUUIDs found for matches to summoners
		idMap := make(map[string]bool)
		newRows := make([]types.SummonerRow, 0, len(profile.MatchData)*9)
		for _, m := range profile.MatchData {
			for _, p := range m.Participants {
				id := p.PUUID
				if _, ok := idMap[id]; !ok && id != profile.PUUID {
					row := types.SummonerRow{
						PUUID:         id,
						Region:        profile.Region,
						GameName:      p.RiotIDGameName,
						TagLine:       p.RiotIDTagline,
						ProfileIconID: p.ProfileIconID,
						SummonerLevel: p.SummonerLevel,
						CreatedAt:     time.Now(),
						UpdatedAt:     time.Now(),
						IsPopulated:   false,
					}
					newRows = append(newRows, row)
					idMap[id] = true
				}
			}
		}

		if err := utils.AddNewSummoners(ctx, pool, newRows); err != nil {
			log.Printf("AddNewSummoners error in /lol/refresh: %v", err)
			return
			// Early returning here because matches and match_participants tables rely on
			// PUUIDs in the summoners table as foreign keys for their entries.
		} else {
			log.Print("New summoners added")
		}

		// update matches tables
		if len(newMatches) > 0 {
			if err := utils.AddMatchData(ctx, pool, newMatches); err != nil {
				log.Printf("AddMatchData error in /lol/refresh: %v", err)
			} else {
				log.Printf("DB successfully updated (matches added: %d)", len(newMatches))
			}
		}

		log.Printf("Refresh post-processing completed in %s", time.Since(start))
	}(ctx, pool, checkCopy, profileCopy, toAddCopy)

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
