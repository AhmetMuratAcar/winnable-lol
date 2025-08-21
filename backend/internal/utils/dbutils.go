package utils

import (
	"context"
	"errors"
	"fmt"
	"time"
	"winnable/internal/types"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// GetPUUID queries summoners table for given user's PUUID
// Example:
//
//	cacheCheck, err := utils.GetPUUID(ctx, h.pool, req)
func GetPUUID(ctx context.Context, pool *pgxpool.Pool, userInfo types.RequestBody) (types.PUUIDCacheCheck, error) {
	var puuid string
	var updatedAt time.Time

	query := `
        SELECT puuid, updated_at
        FROM summoners
        WHERE region = $1
          AND lower(game_name) = lower($2)
          AND lower(tag_line) = lower($3)
        LIMIT 1;
    `

	err := pool.QueryRow(
		ctx, 
		query, 
		userInfo.Region, userInfo.GameName, userInfo.TagLine,
		).Scan(
			&puuid, 
			&updatedAt,
		)
	
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return types.PUUIDCacheCheck{
				Found: false,
			}, nil
		}

		return types.PUUIDCacheCheck{}, fmt.Errorf("getPUUID query failed: %w", err)
	}

	stale := time.Since(updatedAt) > 24 * time.Hour

	return types.PUUIDCacheCheck{
		Found: true,
		PUUID: puuid,
		Stale: stale,
	}, nil
}

// GetMasteries queries masteries table for given user's champion masteries
// Example:
//
//	championMasteries, err := utils.GetMasteries(ctx, h.pool, PUUID)
func GetMasteries(ctx context.Context, pool *pgxpool.Pool, PUUID string) ([]types.ChampionMastery, error) {
	// TODO: actually query DB for user masteries
	var championMasteries []types.ChampionMastery
	return championMasteries, nil
}

// GetMatchIDs queries matches table for given user's past Match IDs
// Example:
//
//	matchIDs, err := utils.GetMatchIDs(ctx, h.pool, PUUID)
func GetMatchIDs(ctx context.Context, pool *pgxpool.Pool, PUUID string) ([]string, error) {
	// TODO: actually query DB for user matches
	var matchIDs []string
	return matchIDs, nil
}

// GetMatchDataByIDs populates a map of matchID: types.LeagueMatch based on
// if the matchID is cached in the matches table
// Example:
//
//	err := utils.GetMatchDataByIDs(ctx, h.pool, matchIDs, &matchDataMap)
func GetMatchDataByIDs(ctx context.Context, pool *pgxpool.Pool, matchIDs []string, matchMap *map[string]*types.LeagueMatch) error {
	// TODO actually loop over matchMap and check cache status
	// Remember if the matchID isnt cached, the value in the map should be set to nil
	return nil
}

// AddMatchData updates the matches table with newly fetched matchData
// Example:
//
//	err := utils.AddMatchData(ctx, h.pool, toAdd)
func AddMatchData(ctx context.Context, pool *pgxpool.Pool, matchData []types.LeagueMatch) error {
	return nil
}