package lolprofilesvc

import (
	"fmt"
	"winnable/internal/types"
)

func UpdateGamesSummary() {}

func ConstructGamesSummary(matches []types.LeagueMatch, PUUID string) (types.GamesSummary, error) {
	if len(matches) == 0 {
		return types.GamesSummary{}, fmt.Errorf("no matches provided")
	}

	kpaByRole := make(map[string]int)
	deathsByRole := make(map[string]int)
	totalKPA := 0
	totalDeaths := 0

	summary := types.GamesSummary{KDAsByRole: make(map[string]float64)}

	for _, m := range matches {
		if m.GameEndedInEarlySurrender {
			// remake
			continue
		}

		if m.QueueId != 420 && m.QueueId != 440 {
			// not soloq or flex
			continue
		}

		userIndex := -1
		for i, p := range m.ParticipantPUUIDs {
			if p == PUUID {
				userIndex = i
				break
			}
		}
		if userIndex == -1 {
			// should not happen but just in case
			continue
		}

		user := m.Participants[userIndex]
		currSummary := types.LeagueMatchSummary{
			ChampionID:    user.ChampionID,
			OppChampionID: -1, // defaulting for now
			Role:          user.TeamPosition,
			Kills:         user.Kills,
			Deaths:        user.Deaths,
			Assists:       user.Assists,
			QueueID:       m.QueueId,
		}

		if m.WinningTeam == user.Team {
			currSummary.DidWin = true
			summary.Wins++
		} else {
			currSummary.DidWin = false
			summary.Losses++
		}

		// finding lane opponent
		for _, o := range m.Participants {
			if o.TeamPosition == user.TeamPosition && o.PUUID != user.PUUID && o.Team != user.Team {
				currSummary.OppChampionID = o.ChampionID
				break
			}
		}

		// accumulating aggregates
		kpa := currSummary.Kills + currSummary.Assists
		kpaByRole[currSummary.Role] += kpa
		deathsByRole[currSummary.Role] += currSummary.Deaths

		totalKPA += kpa
		totalDeaths += currSummary.Deaths

		summary.MatchSummaries = append(summary.MatchSummaries, currSummary)
	}

	// calc KDAs from aggregates
	for role := range kpaByRole {
		d := deathsByRole[role]
		if d == 0 {
			d = 1
		}
		summary.KDAsByRole[role] = float64(kpaByRole[role]) / float64(d)
	}

	if totalDeaths == 0 {
		totalDeaths = 1
	}
	summary.TotalKDA = float64(totalKPA) / float64(totalDeaths)

	return summary, nil
}
