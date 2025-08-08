package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"winnable/internal/handlers"
	"winnable/internal/middleware"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func init() {
	// Load env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or failed to load; using system environment")
	} else {
		log.Println(".env file successfully loaded")
	}
}

func main() {
	if err := run(); err != nil {
		log.Fatalf("Application error: %v", err)
	}
}

func run() error {
	// DB pool setup
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		return fmt.Errorf("DATABASE_URL not set")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return fmt.Errorf("pxpool.New: %w", err)
	}
	defer pool.Close()

	// Config tuning
	pool.Config().MaxConns = 10
	pool.Config().MaxConnLifetime = 30 * time.Minute

	if err := pool.Ping(ctx); err != nil {
		return fmt.Errorf("db ping failed: %w", err)
	}

	// Configure CORS
	middleware.ConfigureAllowedOrigins(os.Getenv("ENV"))

	// Handlers
	mux := http.NewServeMux()
	mux.Handle("/health", middleware.EnableCORS(&handlers.HealthHandler{}))
	mux.Handle("/lol/profile", middleware.EnableCORS(handlers.NewLoLProfileHandler(pool)))

	// Starting server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default fallback
	}

	log.Printf("Server is running on port %s...", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		return err
	}

	return nil
}
