package lol

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"winnable/internal/riot"
	"winnable/internal/types"
	"winnable/internal/utils"

	"github.com/jackc/pgx/v5/pgxpool"
)

type LolLiveHandler struct {
	pool *pgxpool.Pool
}

func NewLolLiveHandler(pool *pgxpool.Pool) *LolLiveHandler {
	return &LolLiveHandler{pool: pool}
}

func (h *LolLiveHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "405 Method Not Allowed: only GET is supported", http.StatusMethodNotAllowed)
		return
	}

	log.Println("----------------------------")
	log.Println("Received LoL live game request")

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
		log.Printf("Error calling SummonerCacheCheck in /lol/live: %v", err)
	}
	var userPUUID string

	client := riot.NewClient()
	if !cacheCheck.Found {
		log.Print("User not in summoners, fetching PUUID from Riot")
		userAccount, err := client.GetSummonerPUUID(req)
		if err != nil {
			log.Printf("Error calling GetSummonerPUUID in /lol/live: %v", err)
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

	out, err := client.GetLiveGame(userPUUID, req.Region)
	if err != nil {
		var RiotAPIEror *types.RiotAPIError
		if errors.As(err, &RiotAPIEror) && RiotAPIEror.StatusCode == http.StatusNotFound {
			log.Printf("%s", RiotAPIEror.Error())
			http.Error(
				w,
				"summoner is not in a game",
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

	// Writing to file for dev
	if os.Getenv("ENV") == "development" {
		riotID := req.GameName + "#" + req.TagLine
		err := utils.WriteLiveGameToFile(out, riotID)
		if err != nil {
			log.Printf("Failed to write mastery to JSON: %v", err)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(out); err != nil {
		log.Printf("Failed to encode user's live game data: %v", err)
		http.Error(
			w,
			"internal server error",
			http.StatusInternalServerError,
		)
		return
	}
}
