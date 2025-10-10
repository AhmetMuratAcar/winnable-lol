package lol

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
	"winnable/internal/lolprofilesvc"
	"winnable/internal/riot"
	"winnable/internal/types"
	"winnable/internal/utils"

	"github.com/jackc/pgx/v5/pgxpool"
)

type LolMatchHandler struct {
	pool *pgxpool.Pool
}

func NewLolMatchHandler(pool *pgxpool.Pool) *LolMatchHandler {
	return &LolMatchHandler{pool: pool}
}

func (h *LolMatchHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "405 Method Not Allowed: only GET is supported", http.StatusMethodNotAllowed)
		return
	}

	log.Println("----------------------------")
	log.Println("Received LoL match request")

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	var out types.LeagueMatch
	var err error

	defer r.Body.Close()
	defer cancel()

	matchID := r.URL.Query().Get("matchID")
	if matchID == "" {
		resp := utils.ErrorResponse{
			Error:         "missing required query parameters",
			MissingParams: []string{"matchID"},
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			log.Printf("Failed to encode user's error data in /lol/match: %v", err)
			http.Error(
				w,
				"internal server error",
				http.StatusInternalServerError,
			)
		}
		return
	}

	var matchRow types.MatchRow
	var participantRows []types.MatchParticipantRow
	matchRow, err = utils.GetMatchRowByID(ctx, h.pool, matchID)
	if err != nil {
		log.Printf("Error calling GetMatchRowByID in /lol/match: %v", err)
		out.GameDuration = -1
	} else {
		participantRows, err = utils.GetParticipantRowsByID(ctx, h.pool, matchID)
		if err != nil {
			log.Printf("Error calling GetParticipantRowsByID in /lol/match: %v", err)
			out.GameDuration = -1
		}
	}

	if out.GameDuration == -1 {
		// Shouldn't happen but in case there is some sort of race condition
		// where the match data isn't in the DB or there was an error querying
		client := riot.NewClient()
		region := strings.Split(matchID, "_")[0]
		out, err = client.GetMatchData(matchID, region)
		if err != nil {
			log.Printf("Error fetching match data from Riot in /lol/match: %v", err)
			http.Error(
				w,
				"internal server error",
				http.StatusInternalServerError,
			)
			return
		}
	} else {
		out = lolprofilesvc.AssembleLeagueMatch(matchRow, participantRows)
	}

	// Writing to file for dev
	if os.Getenv("ENV") == "development" {
		err := utils.WriteMatchDataToFile(out, matchID)
		if err != nil {
			log.Printf("Failed to write match data to JSON: %v", err)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(out); err != nil {
		log.Printf("Failed to encode user's match data in /lol/match: %v", err)
		http.Error(
			w,
			"internal server error",
			http.StatusInternalServerError,
		)
		return
	}
}
