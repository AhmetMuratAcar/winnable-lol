package lolprofilesvc

import (
	"context"
	"log"
	"time"

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

func CachedProfileConstructor(ctx context.Context, pool *pgxpool.Pool, profile *types.LeagueProfilePage) (types.CachedProfileCheckList, error) {
	return types.CachedProfileCheckList{}, nil
}

func ConstructMatchDataMap(ctx context.Context, pool *pgxpool.Pool, matchIDs []string) (map[string]*types.LeagueMatch, error) {
	matchDataMap := make(map[string]*types.LeagueMatch)
	return matchDataMap, nil
}

func FillLoLProfileCacheGaps(
	checklist types.CachedProfileCheckList,
	profile *types.LeagueProfilePage,
	client RiotAPI,
	ctx context.Context,
	pool *pgxpool.Pool,
) error {
	detachedCtx, cancel := context.WithTimeout(
		context.WithoutCancel(ctx),
		5*time.Second,
	)
	defer cancel()

	var err error
	if !checklist.Masteries {
		// checklist.Masteries is only false if NO masteries are present
		championMasteries, err := client.GetSummonerMastery(profile.Region, profile.PUUID)
		if err != nil {
			log.Printf(
				"Error requesting masteries in FillLoLProfileCacheGaps:\nPUUID:%s\nError: %v",
				profile.PUUID,
				err,
			)
		}

		for _, c := range championMasteries {
			profile.MasteryData.TotalMastery += c.ChampionLevel
			profile.MasteryData.TotalMasteryPoints += c.ChampionPoints
		}
		profile.MasteryData.ChampionsPlayed = len(championMasteries)
		profile.MasteryData.ChampionMasteries = championMasteries

		// TODO: async update DB with this info
	}

	if !checklist.Matches {
		numCachedMatches := len(profile.MatchData)
		matchIdIndexMap := make(map[string]int)
		if numCachedMatches > 0 {
			for i, m := range profile.MatchData {
				matchIdIndexMap[m.MatchID] = i
			}
		}

		startIndex := 0
		count := 20
		matchIDs, err := client.GetSummonerMatchIDs(profile.PUUID, startIndex, count)
		if err != nil {
			log.Printf(
				"Error requesting past match IDs in FillLoLProfileCacheGaps: \nPUUID%s\nError: %v",
				profile.PUUID,
				err,
			)
		}

		res := make([]types.LeagueMatch, 0, 20)
		toAdd := make([]types.LeagueMatch, 0, 20)
		for _, id := range matchIDs {
			if index, ok := matchIdIndexMap[id]; ok {
				res = append(res, profile.MatchData[index])
			} else {
				matchData, err := client.GetMatchData(id)
				if err != nil {
					log.Printf(
						"Error fetching matchID %s in FillLoLProfileCacheGaps\nError: %v",
						id,
						err,
					)
					continue
				}

				res = append(res, matchData)
				toAdd = append(toAdd, matchData)
			}
		}
		profile.MatchData = res

		go func(batch []types.LeagueMatch) {
			if err := utils.AddMatchData(detachedCtx, pool, batch); err != nil {
				log.Printf("async AddMatchData error in FillLoLProfileCacheGaps: %v", err)
			}
		}(toAdd)
	}

	if !checklist.Ranks {
		profile.Ranks, err = client.GetSummonerRanks(profile.PUUID, profile.Region)
		if err != nil {
			log.Printf("error fetching summoner ranks in FillLoLProfileCacheGaps. Error: %v", err)
		}

		// TODO: async update DB with this info
	}

	if !checklist.ProfileIcon || !checklist.Level {
		profile.ProfileIconID, profile.Level, err = client.GetSummonerIconAndLevel(
			profile.PUUID,
			profile.Region,
		)
		if err != nil {
			log.Printf("error fetching summoner icon and level in FillLoLProfileCacheGaps. Error: %v", err)
		}

		// TODO: async update DB with this info
	}

	return nil
}
