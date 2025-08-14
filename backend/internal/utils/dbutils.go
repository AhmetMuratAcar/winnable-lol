package utils

import (
	"context"
	"winnable/internal/types"

	"github.com/jackc/pgx/v5/pgxpool"
)

// GetPUUID queries users table for given user's PUUID
// Example:
//
//	cacheCheck, err := utils.GetPUUID(ctx, h.pool, req)
func GetPUUID(ctx context.Context, pool *pgxpool.Pool, userInfo types.RequestBody) (types.PUUIDCacheCheck, error) {
	// TODO: actually query DB for user PUUID
	res := types.PUUIDCacheCheck{
		Stale: true,
	}
	return res, nil
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