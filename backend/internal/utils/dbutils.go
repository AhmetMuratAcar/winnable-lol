package utils

import (
	"context"
	"errors"
	"fmt"
	"time"
	"winnable/internal/config"
	"winnable/internal/types"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// MarkSummonerStale marks a summoner as stale by changing the updated_at column of the
// summoners table to 24 hours before the current time.
func MarkSummonerStale(ctx context.Context, pool *pgxpool.Pool, puuid string) error {
	const query = `
		UPDATE summoners
		SET updated_at = now() - interval '24 hours'
		WHERE puuid = $1;
	`

	cmdTag, err := pool.Exec(ctx, query, puuid)
	if err != nil {
		return fmt.Errorf("failed to mark summoner stale: %w", err)
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("no summoner found with puuid: %s", puuid)
	}

	return nil
}

// GetPUUID queries summoners table for given user's PUUID
// Example:
//
//	cacheCheck, err := utils.GetPUUID(ctx, h.pool, req)
func GetPUUID(ctx context.Context, pool *pgxpool.Pool, userInfo types.RequestBody) (types.PUUIDCacheCheck, error) {
	var puuid string
	var updatedAt time.Time

	const query = `
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
			return types.PUUIDCacheCheck{Found: false}, nil
		}

		return types.PUUIDCacheCheck{Found: false}, fmt.Errorf("getPUUID query failed: %w", err)
	}

	stale := time.Since(updatedAt) > 24*time.Hour

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
	const query = `
		SELECT champion_id, champion_level, champion_points
		FROM champion_masteries
		WHERE puuid = $1
		ORDER BY champion_points DESC;
	`

	rows, err := pool.Query(ctx, query, PUUID)
	if err != nil {
		return nil, fmt.Errorf("GetMasteries query failed: %w", err)
	}
	defer rows.Close()

	championMasteries := make([]types.ChampionMastery, 0, config.LEAGUE_CHAMP_COUNT)
	for rows.Next() {
		var currMastery types.ChampionMastery
		if err := rows.Scan(
			&currMastery.ChampionID,
			&currMastery.ChampionLevel,
			&currMastery.ChampionPoints,
		); err != nil {
			return nil, fmt.Errorf("GetMasteries scan failed: %w", err)
		}

		championMasteries = append(championMasteries, currMastery)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("GetMasteries rows iteration failed: %w", err)
	}

	return championMasteries, nil
}

// GetMatchIDs queries matches table for given user's past 20 Match IDs
// Example:
//
//	matchIDs, err := utils.GetMatchIDs(ctx, h.pool, PUUID)
func GetMatchIDs(ctx context.Context, pool *pgxpool.Pool, PUUID string) ([]string, error) {
	const query = `
		SELECT match_id
		FROM match_participants
		WHERE puuid = $1
		ORDER BY game_start DESC
		LIMIT 20;
	`

	rows, err := pool.Query(ctx, query, PUUID)
	if err != nil {
		return nil, fmt.Errorf("GetMatchIDs query failed: %w", err)
	}
	defer rows.Close()

	matchIDs := make([]string, 0, 20)
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, fmt.Errorf("GetMatchIDs scan failed: %w", err)
		}
		matchIDs = append(matchIDs, id)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetMatchIDs rows iteration failed: %w", err)
	}

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