package utils

import "winnable/internal/types"

func ToLeagueMatch(raw types.RawMatchResponse) types.LeagueMatch {
	res := types.LeagueMatch{
		EndOfGameResult:    raw.Info.EndOfGameResult,
		GameDuration:       raw.Info.GameDuration,
		GameStartTimestamp: raw.Info.GameStartTimestamp,
		GameVersion:        raw.Info.GameVersion,
		MatchID:            raw.Metadata.MatchID,
		ParticipantPUUIDs:  raw.Metadata.Participants,
		QueueId:            raw.Info.QueueID,
	}

	participantCount := len(res.ParticipantPUUIDs) + 1
	for i, p := range raw.Info.Participants {
		curr := types.LeagueMatchParticipant{
			Assists:                     p.Assists,
			ChampionID:                  p.ChampionID,
			ChampLevel:                  p.ChampLevel,
			Deaths:                      p.Deaths,
			GoldEarned:                  p.GoldEarned,
			Kills:                       p.Kills,
			ParticipantIndex:            i,
			ProfileIconID:               p.ProfileIcon,
			PUUID:                       p.Puuid,
			RiotIDGameName:              p.RiotIDGameName,
			RiotIDTagline:               p.RiotIDTagline,
			Summoner1ID:                 p.Summoner1ID,
			Summoner2ID:                 p.Summoner2ID,
			SummonerLevel:               p.SummonerLevel,
			TeamPosition:                p.TeamPosition,
			TotalDamageDealtToChampions: p.TotalDamageDealtToChampions,
			TotalMinionsKilled:          p.TotalMinionsKilled,
			VisionScore:                 p.VisionScore,
		}

		curr.Items = []int{
			p.Item0,
			p.Item1,
			p.Item2,
			p.Item3,
			p.Item4,
			p.Item5,
		}

		if i < participantCount/2 {
			curr.Team = 0
		} else {
			curr.Team = 1
		}

		res.Participants = append(res.Participants, curr)
	}

	for i, t := range raw.Info.Teams {
		for j, b := range t.Bans {
			res.Bans[i][j] = b.ChampionID
		}
	}

	if raw.Info.Teams[0].Win {
		res.WinningTeam = 0
	} else {
		res.WinningTeam = 1
	}

	return res
}
