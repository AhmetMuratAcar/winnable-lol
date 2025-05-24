package handlers

import (
	"net/http"
)

type HealthHandler struct{}

func (h *HealthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// return ok
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
