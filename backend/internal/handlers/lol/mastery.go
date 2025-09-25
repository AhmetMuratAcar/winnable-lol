package lol

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

type LolMasteryHandler struct {
	pool *pgxpool.Pool
}

func NewLolMasteryHandler(pool *pgxpool.Pool) *LolMasteryHandler {
	return &LolMasteryHandler{pool: pool}
}

func (h *LolMasteryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
