package lol

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
	"winnable/internal/types"
	"winnable/internal/utils"

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

	// User's puuid and mastery data should always be in the DB table due to
	// structuring of frontend logic and /lol/profile but checks in place just
	// in case.
	var out types.MasteryData
	cacheCheck, err := utils.SummonerCacheCheck(ctx, h.pool, req)
	if err != nil {
		log.Printf("Error calling SummonerCacheCheck in /lol/mastery: %v", err)
		http.Error(
			w,
			"internal server error",
			http.StatusInternalServerError,
		)
		return
	}

	if !cacheCheck.Found {
		log.Print("User not in summoners table for /lol/mastery")
		http.Error(
			w,
			"user not found",
			http.StatusNotFound,
		)
		return
	}

	// Don't care if the data is stale for mastery
	out.ChampionMasteries, err = utils.GetMasteries(ctx, h.pool, cacheCheck.PUUID)
	if err != nil {
		log.Printf("Error calling GetMasteries in /lol/mastery: %v", err)
		return
	}

	for _, c := range out.ChampionMasteries {
		out.TotalMastery += c.ChampionLevel
		out.TotalMasteryPoints += c.ChampionPoints
	}
	out.ChampionsPlayed = len(out.ChampionMasteries)

	// Writing to file for dev
	if os.Getenv("ENV") == "development" {
		riotID := cacheCheck.GameName + "#" + cacheCheck.TagLine
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
