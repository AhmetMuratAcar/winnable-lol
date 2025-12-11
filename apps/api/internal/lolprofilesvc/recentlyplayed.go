package lolprofilesvc

import (
	"fmt"
	"sort"

	"github.com/AhmetMuratAcar/winnable-lol/apps/api/internal/types"
)

func UpdateRecentlyPlayedWithAndAgainst() {}

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
