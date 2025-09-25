package lol

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

type LolLiveHandler struct {
	pool *pgxpool.Pool
}

func NewLolLiveHandler(pool *pgxpool.Pool) *LolLiveHandler {
	return &LolLiveHandler{pool: pool}
}

func (h *LolLiveHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
