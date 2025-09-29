package lol

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
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "405 Method Not Allowed: only GET is supported", http.StatusMethodNotAllowed)
		return
	}

	log.Println("----------------------------")
	log.Println("Received LoL profile request")

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

	// PUUID calls
	client := riot.NewClient()
	userProfile := types.LeagueProfilePage{Region: req.Region}

	cacheCheck, err := utils.SummonerCacheCheck(ctx, h.pool, req)
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
		log.Print("user is cached, fetching PUUID from DB")
		userProfile.PUUID = cacheCheck.PUUID
		userProfile.ProfileIconID = cacheCheck.ProfileIconID
		userProfile.Level = cacheCheck.Level
		userProfile.GameName = cacheCheck.GameName
		userProfile.TagLine = cacheCheck.TagLine
		userProfile.LastUpdated = cacheCheck.LastUpdated

		if !cacheCheck.Stale && cacheCheck.IsPopulated {
			// guaranteed to have good data to return
			log.Print("data not stale and populated, fetching from DB")

			userProfile, err = lolprofilesvc.CachedProfileConstructor(ctx, h.pool, userProfile)
			if err != nil {
				// default back to Riot API calls
				log.Printf("Error calling CachedProfileConstructor: %v", err)
				err = utils.MarkSummonerStale(ctx, h.pool, userProfile.PUUID)
				if err != nil {
					log.Printf("Error marking summoner as stale: %v", err)
				}

				cacheCheck.IsPopulated = false
				cacheCheck.Found = false
				cacheCheck.Stale = true
			}
		}
	} else {
		riotAccount, err := client.GetSummonerPUUID(req)

		if err != nil {
			var httpErr *types.HTTPError
			if errors.As(err, &httpErr) {
				log.Printf("GetSummonerPUUID HTTPError: status=%d err=%v", httpErr.StatusCode, err)
				http.Error(
					w,
					"could not fetch summoner",
					httpErr.StatusCode,
				)
				return
			} else {
				log.Printf("GetSummonerPUUID internal error: %v", err)
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
		}

		userProfile.PUUID = riotAccount.Puuid
		userProfile.GameName = riotAccount.GameName
		userProfile.TagLine = riotAccount.TagLine
		userProfile.LastUpdated = time.Now()
		// Only early terminating if PUUID fetch fails. If other client requests fail
		// the userProfile is constructed with any successfully received data.
	}
	log.Print("PUUID fetch successful")

	// Rank Calls
	if !cacheCheck.Found || !cacheCheck.IsPopulated || cacheCheck.Stale {
		userProfile.Ranks, err = client.GetSummonerRanks(userProfile.PUUID, userProfile.Region)
		if err != nil {
			log.Printf("error fetching summoner ranks. Error: %v", err)
		} else {
			log.Print("Summoner ranks successfully added")
		}
	}

	// Past matches calls
	// No matter the cached status all matches are added here
	var matchIDs []string
	startIndex := 0
	matchCount := 20

	if cacheCheck.Found && cacheCheck.IsPopulated && !cacheCheck.Stale {
		matchIDs, err = utils.GetMatchIDs(ctx, h.pool, userProfile.PUUID)
		if err != nil {
			log.Printf("Error querying DB for past match IDs: %v", err)
		}
	}

	if len(matchIDs) < 20 {
		matchIDs, err = client.GetSummonerMatchIDs(userProfile.PUUID, startIndex, matchCount)
		if err != nil {
			log.Printf("Error requesting past match IDs from Riot: %v", err)
		}
	}
	log.Print("MatchIDs fetch successful")

	matchDataMap, err := lolprofilesvc.ConstructMatchDataMap(ctx, h.pool, matchIDs)
	if err != nil {
		log.Printf(
			"Error constructing matchDataMap PUUID: %s\nmatchIDs: %s\nerr: %v",
			userProfile.PUUID,
			matchIDs,
			err,
		)
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

	// Level and icon ID calls
	var match types.LeagueMatch
	if len(userProfile.MatchData) > 0 {
		match = userProfile.MatchData[0]
	}
	endTime := time.UnixMilli(int64(match.GameStartTimestamp)).
		Add(time.Duration(match.GameDuration) * time.Second)

	if endTime.After(cacheCheck.LastUpdated) && len(userProfile.MatchData) > 0 {
		var userIndex int
		for _, v := range match.Participants {
			if v.PUUID == userProfile.PUUID {
				userIndex = v.ParticipantIndex
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

	// Summary components' data
	userProfile.PlayedWith, userProfile.PlayedAgainst, err = lolprofilesvc.ConstructRecentlyPlayedWithAndAgainst(userProfile.MatchData, userProfile.PUUID)
	if err != nil {
		log.Printf("error constructing recently played with: %v", err)
	} else {
		log.Print("recently played with constructed")
	}

	userProfile.RecentGames, err = lolprofilesvc.ConstructGamesSummary(userProfile.MatchData, userProfile.PUUID)
	if err != nil {
		log.Printf("error constructing games summary: %v", err)
	} else {
		log.Print("games summary constructed")
	}

	// Writing to file for dev
	if os.Getenv("ENV") == "development" {
		riotID := userProfile.GameName + "#" + userProfile.TagLine
		err := utils.WriteProfileToFile(userProfile, riotID)
		if err != nil {
			log.Printf("Failed to write profile to JSON. Error: %v", err)
		}
	}

	// Updating DB with new data
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
		log.Print("---START POST-PROCESSING---")
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
				log.Printf("Error calling GetSummonerMastery in /lol/profile: %v", err)
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
			log.Printf("SyncProfileData error: %v", err)
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
			log.Printf("AddNewSummoners error: %v", err)
			return
			// Early returning here because matches and match_participants tables rely on
			// PUUIDs in the summoners table as foreign keys for their entries.
		} else {
			log.Print("New summoners added")
		}

		// update matches tables
		if len(newMatches) > 0 {
			if err := utils.AddMatchData(ctx, pool, newMatches); err != nil {
				log.Printf("AddMatchData error: %v", err)
			} else {
				log.Printf("DB successfully updated (matches added: %d)", len(newMatches))
			}
		}

		log.Printf("Post-processing completed in %s", time.Since(start))
	}(ctx, pool, checkCopy, profileCopy, toAddCopy)

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
