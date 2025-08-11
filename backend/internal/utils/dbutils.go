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
