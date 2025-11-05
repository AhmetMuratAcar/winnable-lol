package lolprofilesvc

import (
	"context"
	"fmt"

	"winnable/internal/types"
	"winnable/internal/utils"

	"github.com/jackc/pgx/v5/pgxpool"
)

type RiotAPI interface {
	GetSummonerPUUID(reqBody types.RequestBody) (puuid string, err error)
	GetSummonerMastery(region, puuid string) ([]types.ChampionMastery, error)
	GetSummonerMatchIDs(puuid string, start int, count int) ([]string, error)
	GetMatchData(matchID string) (types.LeagueMatch, error)
	GetSummonerIconAndLevel(puuid, region string) (int, int, error)
	GetSummonerRanks(puuid, region string) ([]types.LeagueRank, error)
}

// CachedProfileConstructor fills out the MasteryData, and Ranks fields of a cached LeagueProfilePage.
func CachedProfileConstructor(ctx context.Context, pool *pgxpool.Pool, profile types.LeagueProfilePage) (types.LeagueProfilePage, error) {
	var err error
	out := profile

	out.MasteryData.ChampionMasteries, err = utils.GetMasteries(ctx, pool, out.PUUID)
	if err != nil {
		return profile, fmt.Errorf("error calling GetMasteries in CachedProfileConstructor: %w", err)
	}

	for _, c := range out.MasteryData.ChampionMasteries {
		out.MasteryData.TotalMastery += c.ChampionLevel
		out.MasteryData.TotalMasteryPoints += c.ChampionPoints
	}
	out.MasteryData.ChampionsPlayed = len(out.MasteryData.ChampionMasteries)

	out.Ranks, err = utils.GetSummonerRanks(ctx, pool, out.PUUID)
	if err != nil {
		return profile, fmt.Errorf("error calling GetSummonerRanks in CachedProfileConstructor: %w", err)
	}

	return out, nil
}

func ConstructMatchDataMap(ctx context.Context, pool *pgxpool.Pool, matchIDs []string) (map[string]*types.LeagueMatch, error) {
	out := make(map[string]*types.LeagueMatch, len(matchIDs))
	if len(matchIDs) == 0 {
		return out, nil
	}

	for _, id := range matchIDs {
		out[id] = nil
	}

	matches, err := utils.GetMatchesByIDs(ctx, pool, matchIDs)
	if err != nil {
		return nil, err
	}

	participantsByMatch, err := utils.GetParticipantsForMatches(ctx, pool, matchIDs)
	if err != nil {
		return nil, err
	}

	for mID, m := range matches {
		ps := participantsByMatch[mID]
		lm := AssembleLeagueMatch(m, ps)
		// copy to new var to safely take address in loop
		v := lm
		out[mID] = &v
	}

	return out, nil
}

// AssembleLeagueMatch converts DB rows -> API struct
func AssembleLeagueMatch(m types.MatchRow, ps []types.MatchParticipantRow) types.LeagueMatch {
	lm := types.LeagueMatch{
		EndOfGameResult:           m.EndOfGameResult,
		GameDuration:              m.GameDurationSec,
		GameEndedInEarlySurrender: m.GameEndedInEarlySurrender,
		GameStartTimestamp:        int(m.GameStart.UnixMilli()),
		GameVersion:               m.GameVersion,
		MatchID:                   m.MatchID,
		ParticipantPUUIDs:         make([]string, 0, len(ps)),
		Participants:              make([]types.LeagueMatchParticipant, 0, len(ps)),
		QueueId:                   m.QueueID,
		Bans:                      [][]int{m.BansBlue, m.BansRed},
		WinningTeam:               m.WinningTeam,
		MostDamageDone:            0,
		MostDamageTaken:           0,
	}

	for _, p := range ps {
		lm.Participants = append(lm.Participants, types.LeagueMatchParticipant{
			Assists:                     p.Assists,
			ChampionID:                  p.ChampionID,
			ChampLevel:                  p.ChampLevel,
			Deaths:                      p.Deaths,
			GoldEarned:                  p.GoldEarned,
			Items:                       p.Items,
			Kills:                       p.Kills,
			ParticipantIndex:            p.ParticipantIndex,
			ProfileIconID:               p.ProfileIconAtMatch,
			PUUID:                       p.PUUID,
			RiotIDGameName:              p.RiotIDGameName,
			RiotIDTagline:               p.RiotIDTagLine,
			Summoner1ID:                 p.Summoner1ID,
			Summoner2ID:                 p.Summoner2ID,
			SummonerLevel:               p.SummonerLevelAtMatch,
			Team:                        p.Team,
			TeamPosition:                p.TeamPosition,
			TotalDamageTaken:            p.TotalDamageTaken,
			TotalDamageDealtToChampions: p.TotalDamageToChamps,
			TotalMinionsKilled:          p.TotalMinionsKilled,
			Runes:                       p.Runes,
			ControlWardsPlaced:          p.ControlWardsPlaced,
			WardsPlaced:                 p.WardsPlaced,
			WardsKilled:                 p.WardsKilled,
		})

		if p.TotalDamageToChamps > lm.MostDamageDone {
			lm.MostDamageDone = p.TotalDamageToChamps
		}

		if p.TotalDamageTaken > lm.MostDamageTaken {
			lm.MostDamageTaken = p.TotalDamageTaken
		}

		lm.ParticipantPUUIDs = append(lm.ParticipantPUUIDs, p.PUUID)
	}

	return lm
}
