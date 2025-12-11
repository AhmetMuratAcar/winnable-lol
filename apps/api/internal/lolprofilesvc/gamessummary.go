package lolprofilesvc

import (
	"fmt"
	"maps"

	"github.com/AhmetMuratAcar/winnable-lol/apps/api/internal/types"
)

func UpdateGamesSummary() {}

func ConstructGamesSummary(matches []types.LeagueMatch, PUUID string) (types.GamesSummary, error) {
	if len(matches) == 0 {
		return types.GamesSummary{}, fmt.Errorf("no matches provided")
	}

	summary := types.GamesSummary{
		TotalsByRole:   make(map[string]types.RoleSummary),
		TotalsByQueue:  make(map[int]types.RoleSummary),
		MatchSummaries: []types.LeagueMatchSummary{},
	}
	roleTotals := make(map[string]types.RoleSummary)
	queueTotals := make(map[int]types.RoleSummary)

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
		didWin := m.WinningTeam == user.Team
		currSummary := types.LeagueMatchSummary{
			ChampionID:    user.ChampionID,
			OppChampionID: -1, // defaulting for now
			Role:          user.TeamPosition,
			Kills:         user.Kills,
			Deaths:        user.Deaths,
			Assists:       user.Assists,
			QueueID:       m.QueueId,
			DidWin:        didWin,
		}

		// finding lane opponent
		for _, o := range m.Participants {
			if o.TeamPosition == user.TeamPosition && o.PUUID != user.PUUID && o.Team != user.Team {
				currSummary.OppChampionID = o.ChampionID
				break
			}
		}
		summary.MatchSummaries = append(summary.MatchSummaries, currSummary)

		// totals calcs
		qt := queueTotals[m.QueueId]
		updateTotals(&qt, user.Kills, user.Deaths, user.Assists, didWin)
		queueTotals[m.QueueId] = qt

		rt := roleTotals[user.TeamPosition]
		updateTotals(&rt, user.Kills, user.Deaths, user.Assists, didWin)
		roleTotals[user.TeamPosition] = rt

		updateTotals(&summary.TotalsAll, user.Kills, user.Deaths, user.Assists, didWin)
	}

	maps.Copy(summary.TotalsByQueue, queueTotals)
	maps.Copy(summary.TotalsByRole, roleTotals)

	return summary, nil
}

func updateTotals(rs *types.RoleSummary, kills, deaths, assists int, didWin bool) {
	rs.Kills += kills
	rs.Deaths += deaths
	rs.Assists += assists
	rs.Games++
	if didWin {
		rs.Wins++
	} else {
		rs.Losses++
	}
}
