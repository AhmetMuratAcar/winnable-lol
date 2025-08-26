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

// SyncProfileData updates all of a league profile's data in the DB.
//
// Updates tables: summoners, champion_masteries, ranks
func SyncProfileData(
	ctx context.Context,
	pool *pgxpool.Pool,
	cacheCheck types.PUUIDCacheCheck,
	userProfile types.LeagueProfilePage,
) error {
	var sr = []types.SummonerRow{{
		PUUID:              userProfile.PUUID,
		Region:             userProfile.Region,
		GameName:           userProfile.GameName,
		TagLine:            userProfile.TagLine,
		ProfileIconID:      userProfile.ProfileIconID,
		SummonerLevel:      userProfile.Level,
		TotalMastery:       userProfile.MasteryData.TotalMastery,
		TotalMasteryPoints: userProfile.MasteryData.TotalMasteryPoints,
		ChampionsPlayed:    userProfile.MasteryData.ChampionsPlayed,
		UpdatedAt:          time.Now(),
		IsPopulated:        true,
	}}

	// summoners table conditional
	if cacheCheck.Found {
		if err := UpdateSummonersAll(ctx, pool, sr); err != nil {
			return fmt.Errorf("error calling UpdateSummonersAll: %w", err)
		}
	} else {
		sr[0].CreatedAt = time.Now()
		if err := AddNewSummoners(ctx, pool, sr); err != nil {
			return fmt.Errorf("error calling AddNewSummoners: %w", err)
		}
	}

	// add everything else
	if err := AddMasteries(ctx, pool, userProfile.PUUID, userProfile.MasteryData.ChampionMasteries); err != nil {
		return fmt.Errorf("error calling AddMasteries: %w", err)
	}

	if err := AddRanks(ctx, pool, userProfile.PUUID, userProfile.Ranks); err != nil {
		return fmt.Errorf("error calling AddRanks: %w", err)
	}

	return nil
}

/* ------------------------ SELECT Queries ------------------------ */

// GetPUUID queries summoners table for given user's PUUID
func GetPUUID(ctx context.Context, pool *pgxpool.Pool, userInfo types.RequestBody) (types.PUUIDCacheCheck, error) {
	var puuid string
	var updatedAt time.Time
	var isPopulated bool
	var iconID int
	var level int

	const query = `
        SELECT puuid, updated_at, is_populated, profile_icon_id, summoner_level
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
		&isPopulated,
		&iconID,
		&level,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return types.PUUIDCacheCheck{Found: false}, nil
		}

		return types.PUUIDCacheCheck{Found: false}, fmt.Errorf("getPUUID query failed: %w", err)
	}

	stale := time.Since(updatedAt) > 24*time.Hour

	return types.PUUIDCacheCheck{
		Found:         true,
		PUUID:         puuid,
		Stale:         stale,
		IsPopulated:   isPopulated,
		LastUpdated:   updatedAt,
		ProfileIconID: iconID,
		Level:         level,
	}, nil
}

func GetSummonerRanks(ctx context.Context, pool *pgxpool.Pool, PUUID string) ([]types.LeagueRank, error) {
	if PUUID == "" {
		return []types.LeagueRank{}, nil
	}

	const q = `
		SELECT
			queue_type,
			tier,
			rank,
			league_points,
			wins,
			losses
		FROM ranks
		WHERE puuid = $1
		ORDER BY queue_type
	`

	rows, err := pool.Query(ctx, q, PUUID)
	if err != nil {
		return nil, fmt.Errorf("query ranks: %w", err)
	}
	defer rows.Close()

	out := make([]types.LeagueRank, 0, 2) // most players have 0–2 queues
	for rows.Next() {
		var r types.LeagueRank
		if err := rows.Scan(
			&r.QueueType,
			&r.Tier,
			&r.Rank,
			&r.LeaguePoints,
			&r.Wins,
			&r.Losses,
		); err != nil {
			return nil, fmt.Errorf("scan rank row: %w", err)
		}
		out = append(out, r)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate ranks: %w", err)
	}
	return out, nil
}

// GetMasteries queries champion_masteries table for given user's champion masteries
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

// GetMatchIDs queries match_participants table for given user's past 20 Match IDs.
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

func GetMatchesByIDs(ctx context.Context, pool *pgxpool.Pool, ids []string) (map[string]types.MatchRow, error) {
	out := make(map[string]types.MatchRow, len(ids))
	return out, nil
}

func GetParticipantsForMatches(ctx context.Context, pool *pgxpool.Pool, ids []string) (map[string][]types.MatchParticipantRow, error) {
	out := make(map[string][]types.MatchParticipantRow, len(ids))
	return out, nil
}

/* ------------------------ INSERT Queries ------------------------ */

// AddMatchData updates the matches and match_participants tables with newly fetched matchData.
func AddMatchData(ctx context.Context, pool *pgxpool.Pool, matchData []types.LeagueMatch) error {
	return nil
}

// AddRanks updates the ranks table for a given PUUID. If a rank row for that (PUUID, queueType) exists, it updates that CHANGED data for that row.
func AddRanks(ctx context.Context, pool *pgxpool.Pool, puuid string, ranks []types.LeagueRank) error {
	if puuid == "" || len(ranks) == 0 {
		return nil
	}

	query := `
		INSERT INTO ranks (
			puuid,
			queue_type,
			tier,
			rank,
			league_points,
			wins,
			losses
		)
		VALUES ($1,$2,$3,$4,$5,$6,$7)
		ON CONFLICT (puuid, queue_type)
		DO UPDATE SET
			tier          = EXCLUDED.tier,
			rank          = EXCLUDED.rank,
			league_points = EXCLUDED.league_points,
			wins          = EXCLUDED.wins,
			losses        = EXCLUDED.losses
		WHERE ranks.tier       IS DISTINCT FROM EXCLUDED.tier
		OR ranks.rank          IS DISTINCT FROM EXCLUDED.rank
		OR ranks.league_points IS DISTINCT FROM EXCLUDED.league_points
		OR ranks.wins          IS DISTINCT FROM EXCLUDED.wins
		OR ranks.losses        IS DISTINCT FROM EXCLUDED.losses
	`

	tx, err := pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}
	defer tx.Rollback(ctx) // no-call on successful tx.Commit

	batch := &pgx.Batch{}
	for _, r := range ranks {
		batch.Queue(
			query,
			puuid,
			r.QueueType,
			r.Tier,
			r.Rank,
			r.LeaguePoints,
			r.Wins,
			r.Losses,
		)
	}

	br := tx.SendBatch(ctx, batch)

	for range ranks {
		if _, err := br.Exec(); err != nil {
			return fmt.Errorf("upsert rank (puuid=%s) failed: %w", puuid, err)
		}
	}

	if err := br.Close(); err != nil {
		return fmt.Errorf("close batch: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("commit tx: %w", err)
	}
	return nil
}

// AddMasteries adds all given masteries for a given PUUID. If a mastery row for that (PUUID, championID) exists, it updates the data for that row.
func AddMasteries(ctx context.Context, pool *pgxpool.Pool, puuid string, masteries []types.ChampionMastery) error {
	if puuid == "" || len(masteries) == 0 {
		return nil
	}

	const query = `
		INSERT INTO champion_masteries (
			puuid,
			champion_id,
			champion_level,
			champion_points
		)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (puuid, champion_id)
		DO UPDATE SET
			champion_level  = EXCLUDED.champion_level,
			champion_points = EXCLUDED.champion_points
		WHERE champion_masteries.champion_level IS DISTINCT FROM EXCLUDED.champion_level
   		OR champion_masteries.champion_points IS DISTINCT FROM EXCLUDED.champion_points
	`

	tx, err := pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}
	defer tx.Rollback(ctx) // no-call on successful tx.Commit

	batch := &pgx.Batch{}
	for _, m := range masteries {
		batch.Queue(query, puuid, m.ChampionID, m.ChampionLevel, m.ChampionPoints)
	}

	br := tx.SendBatch(ctx, batch)

	for range masteries {
		if _, err := br.Exec(); err != nil {
			return fmt.Errorf("upsert mastery (puuid=%s) failed: %w", puuid, err)
		}
	}

	if err := br.Close(); err != nil {
		return fmt.Errorf("close batch: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("commit tx: %w", err)
	}
	return nil
}

// AddNewSummoners inserts new PUUIDs into the summoners table
//
// Only the PUUID, CreatedAt(now), UpdatedAt(now), and IsPopulated(false) columns are populated
func AddNewSummoners(ctx context.Context, pool *pgxpool.Pool, rows []types.SummonerRow) error {
	if len(rows) == 0 {
		return nil
	}

	const query = `
		INSERT INTO summoners (
			puuid,
			region,
			game_name,
			tag_line,
			profile_icon_id,
			summoner_level,
			total_mastery,
			total_mastery_points,
			champions_played,
			created_at,
			updated_at,
			is_populated
		)
		VALUES (
			$1, $2, $3, $4, $5,
			$6, $7, $8, $9, $10,
			$11, $12
		)
		ON CONFLICT (puuid) DO NOTHING
	`

	batch := &pgx.Batch{}
	for _, r := range rows {
		if r.CreatedAt.IsZero() {
			r.CreatedAt = time.Now()
		}

		if r.UpdatedAt.IsZero() {
			r.UpdatedAt = time.Now()
		}

		batch.Queue(
			query,
			r.PUUID,
			r.Region,
			r.GameName,
			r.TagLine,
			r.ProfileIconID,
			r.SummonerLevel,
			r.TotalMastery,
			r.TotalMasteryPoints,
			r.ChampionsPlayed,
			r.CreatedAt,
			r.UpdatedAt,
			r.IsPopulated,
		)
	}

	br := pool.SendBatch(ctx, batch)
	defer br.Close()

	for range rows {
		if _, err := br.Exec(); err != nil {
			return fmt.Errorf("insert summoiner failed: %w", err)
		}
	}

	return nil
}

/* ------------------------ UPDATE Queries ------------------------ */

// UpdateSummonersAll updates all contents of a summoners' table row for each row given in rows
func UpdateSummonersAll(ctx context.Context, pool *pgxpool.Pool, rows []types.SummonerRow) error {
	if len(rows) == 0 {
		return nil
	}

	const query = `
		UPDATE summoners
		SET
			region               = $2,
			game_name            = $3,
			tag_line             = $4,
			profile_icon_id      = $5,
			summoner_level       = $6,
			total_mastery        = $7,
			total_mastery_points = $8,
			champions_played     = $9,
			updated_at           = $10,
			is_populated         = $11
		WHERE puuid = $1
	`

	batch := &pgx.Batch{}
	for _, r := range rows {
		if r.UpdatedAt.IsZero() {
			r.UpdatedAt = time.Now()
		}

		batch.Queue(
			query,
			r.PUUID,
			r.Region,
			r.GameName,
			r.TagLine,
			r.ProfileIconID,
			r.SummonerLevel,
			r.TotalMastery,
			r.TotalMasteryPoints,
			r.ChampionsPlayed,
			r.UpdatedAt,
			r.IsPopulated,
		)
	}

	br := pool.SendBatch(ctx, batch)
	defer br.Close()

	for range rows {
		if _, err := br.Exec(); err != nil {
			return fmt.Errorf("update summoner failed: %w", err)
		}
	}

	return nil
}

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
