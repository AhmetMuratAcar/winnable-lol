package lol

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/AhmetMuratAcar/winnable-lol/apps/api/internal/riot"
	"github.com/AhmetMuratAcar/winnable-lol/apps/api/internal/types"
	"github.com/AhmetMuratAcar/winnable-lol/apps/api/internal/utils"

	"github.com/jackc/pgx/v5/pgxpool"
)

type LolMasteryHandler struct {
	pool *pgxpool.Pool
}

func NewLolMasteryHandler(pool *pgxpool.Pool) *LolMasteryHandler {
	return &LolMasteryHandler{pool: pool}
}

func (h *LolMasteryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "405 Method Not Allowed: only GET is supported", http.StatusMethodNotAllowed)
		return
	}

	log.Println("----------------------------")
	log.Println("Received LoL mastery request")

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
		log.Printf("Error calling SummonerCacheCheck in /lol/mastery: %v", err)
	}
	var userPUUID string

	client := riot.NewClient()
	if !cacheCheck.Found {
		log.Print("User not in summoners, fetching PUUID and mastery data from Riot")
		userAccount, err := client.GetSummonerPUUID(req)
		if err != nil {
			// TODO: Update this logic to be like the block below is the As logic. Have to Update
			// GetSummonerPUUID first.
			log.Printf("Error calling GetSummonerPUUID in /lol/mastery: %v", err)
			http.Error(
				w,
				"internal server error",
				http.StatusInternalServerError,
			)
			return
		}
		userPUUID = userAccount.Puuid
	} else {
		userPUUID = cacheCheck.PUUID
	}

	// Don't care if the data is stale for mastery
	var out types.MasteryData
	if cacheCheck.Found && cacheCheck.IsPopulated {
		out.ChampionMasteries, err = utils.GetMasteries(ctx, h.pool, userPUUID)
		if err != nil {
			log.Printf("Error calling GetMasteries in /lol/mastery: %v", err)
			cacheCheck.Found = false // to default below
		}
	}

	if !cacheCheck.Found {
		out.ChampionMasteries, err = client.GetSummonerMastery(req.Region, userPUUID)
		if err != nil {
			var RiotAPIEror *types.RiotAPIError
			if errors.As(err, &RiotAPIEror) && RiotAPIEror.StatusCode == http.StatusNotFound {
				log.Printf("%s", RiotAPIEror.Error())
				http.Error(
					w,
					"summoner not found",
					http.StatusNotFound,
				)
			} else {
				log.Printf("Riot API error: %v", err)
				http.Error(
					w,
					"internal server error",
					http.StatusInternalServerError,
				)
			}

			return
		}
	}

	for _, c := range out.ChampionMasteries {
		out.TotalMastery += c.ChampionLevel
		out.TotalMasteryPoints += c.ChampionPoints
	}
	out.ChampionsPlayed = len(out.ChampionMasteries)

	// Writing to file for dev
	if os.Getenv("ENV") == "development" {
		riotID := req.GameName + "#" + req.TagLine
		err := utils.WriteMasteryToFile(out, riotID)
		if err != nil {
			log.Printf("Failed to write mastery to JSON: %v", err)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(out); err != nil {
		log.Printf("failed to encode user's mastery data: %v", err)
		http.Error(
			w,
			"internal server error",
			http.StatusInternalServerError,
		)
		return
	}
}
