package middleware

import (
	"log"
	"net/http"
	"os"
)

var allowedOrigins map[string]bool

func ConfigureAllowedOrigins(env string) {
	var origin string

	if env == "production" {
		origin = os.Getenv("PROD_ORIGIN")
		if origin == "" {
			log.Fatal("Missing required env variable: PROD_ORIGIN")
		}
	} else {
		origin = os.Getenv("DEV_ORIGIN")
		if origin == "" {
			log.Fatal("Missing required env variable: DEV_ORIGIN")
		}
	}

	allowedOrigins = map[string]bool{
		origin: true,
	}
}

func EnableCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")

		if allowedOrigins[origin] {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		}

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		h.ServeHTTP(w, r)
	})
}
