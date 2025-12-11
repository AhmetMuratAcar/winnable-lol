package handlers

import (
	"net/http"
)

type HealthHandler struct{}

func (h *HealthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "405 Method Not Allowed: only GET is supported", http.StatusMethodNotAllowed)
		return
	}

	// return ok
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("lookin good\n"))
}
