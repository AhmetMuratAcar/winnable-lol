package utils

import "winnable/internal/types"

func ToLeagueMatch(raw types.RawMatchResponse) types.LeagueMatch {
	res := types.LeagueMatch{
		EndOfGameResult:           raw.Info.EndOfGameResult,
		GameDuration:              raw.Info.GameDuration,
		GameEndedInEarlySurrender: raw.Info.Participants[0].GameEndedInEarlySurrender,
		GameStartTimestamp:        raw.Info.GameStartTimestamp,
		GameVersion:               raw.Info.GameVersion,
		MatchID:                   raw.Metadata.MatchID,
		ParticipantPUUIDs:         raw.Metadata.Participants,
		QueueId:                   raw.Info.QueueID,
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

		curr.Runes = types.SummonerRunes{
			StatPerks: types.StatPerks{
				Defense: p.Perks.StatPerks.Defense,
				Flex:    p.Perks.StatPerks.Flex,
				Offense: p.Perks.StatPerks.Offense,
			},
			MainTree: types.MainRuneTree{
				Keystone: p.Perks.Styles[0].Selections[0].Perk,
				Rune1:    p.Perks.Styles[0].Selections[1].Perk,
				Rune2:    p.Perks.Styles[0].Selections[2].Perk,
				Rune3:    p.Perks.Styles[0].Selections[3].Perk,
			},
			SecondaryTree: types.SecondaryRuneTree{
				Rune1: p.Perks.Styles[1].Selections[0].Perk,
				Rune2: p.Perks.Styles[1].Selections[1].Perk,
			},
		}

		curr.Items = []int{
			p.Item0,
			p.Item1,
			p.Item2,
			p.Item3,
			p.Item4,
			p.Item5,
			p.Item6,
		}

		if i < participantCount/2 {
			curr.Team = 0
		} else {
			curr.Team = 1
		}

		res.Participants = append(res.Participants, curr)
	}

	for _, t := range raw.Info.Teams {
		row := make([]int, 0, len(t.Bans))
		for _, b := range t.Bans {
			row = append(row, b.ChampionID)
		}
		res.Bans = append(res.Bans, row)
	}

	if raw.Info.Teams[0].Win {
		res.WinningTeam = 0
	} else {
		res.WinningTeam = 1
	}

	return res
}
