package utils

import (
	"fmt"
	"sort"
	"winnable/internal/types"
)

func ConstructRecentlyPlayedWithAndAgainst(
	matches []types.LeagueMatch,
	selfPUUID string,
) (
	playedWithSummoners []types.PlayedSummoner,
	playedAgainstSummoners []types.PlayedSummoner,
	err error,
) {
	if len(matches) == 0 {
		return nil, nil, fmt.Errorf("no matches provided")
	}

	// key: PUUID
	playedWith := make(map[string]types.PlayedSummoner)
	playedAgainst := make(map[string]types.PlayedSummoner)

	// PUUID -> latest game end time
	lastSeenWith := make(map[string]int)
	lastSeenAgainst := make(map[string]int)

	foundSelf := false
	for _, m := range matches {
		if m.GameEndedInEarlySurrender {
			// remake
			continue
		}

		endTs := int(m.GameStartTimestamp) + int(m.GameDuration)*1000
		selfTeam := -1
		for _, p := range m.Participants {
			if p.PUUID == selfPUUID {
				selfTeam = p.Team
				foundSelf = true
				break
			}
		}

		if selfTeam == -1 {
			// should not happen but accounting for it just in case
			continue
		}

		didWin := (m.WinningTeam == selfTeam)
		for _, p := range m.Participants {
			if p.PUUID == selfPUUID {
				continue
			}

			var bucket map[string]types.PlayedSummoner
			if p.Team == selfTeam {
				bucket = playedWith
			} else {
				bucket = playedAgainst
			}

			key := p.PUUID
			entry := bucket[key]
			entry.GamesPlayed++
			entry.GameName = p.RiotIDGameName
			entry.TagLine = p.RiotIDTagline

			if didWin == (selfTeam == p.Team) {
				entry.Wins++
			} else {
				entry.Losses++
			}

			if p.Team == selfTeam {
				if endTs > lastSeenWith[key] {
					lastSeenWith[key] = endTs
					entry.ProfileIconID = p.ProfileIconID
				}
			} else {
				if endTs > lastSeenAgainst[key] {
					lastSeenAgainst[key] = endTs
					entry.ProfileIconID = p.ProfileIconID
				}
			}
			
			bucket[key] = entry
		}
	}

	if !foundSelf {
		return nil, nil, fmt.Errorf("self PUUID %s not found in any match", selfPUUID)
	}

	toSortedList := func(m map[string]types.PlayedSummoner, seen map[string]int) []types.PlayedSummoner {
		type kv struct {
			key string
			val types.PlayedSummoner
		}
		pairs := make([]kv, 0, len(m))
		for k, v := range m {
			pairs = append(pairs, kv{key: k, val: v})
		}

		sort.Slice(pairs, func(i, j int) bool {
			// sorting based on 4 keys:
			// player with most games
			if pairs[i].val.GamesPlayed != pairs[j].val.GamesPlayed {
				return pairs[i].val.GamesPlayed > pairs[j].val.GamesPlayed
			}

			// player with most wins
			if pairs[i].val.Wins != pairs[j].val.Wins {
				return pairs[i].val.Wins > pairs[j].val.Wins
			}

			// player in most recent game
			if seen[pairs[i].key] != seen[pairs[j].key] {
				return seen[pairs[i].key] > seen[pairs[j].key]
			}

			// lexicographically descending
			return pairs[i].val.GameName < pairs[j].val.GameName
		})

		if len(pairs) > 10 {
			pairs = pairs[:10]
		}

		out := make([]types.PlayedSummoner, 0, len(pairs))
		for _, kv := range pairs {
			out = append(out, kv.val)
		}
		return out
	}
	return toSortedList(playedWith, lastSeenWith), toSortedList(playedAgainst, lastSeenAgainst), nil
}

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
