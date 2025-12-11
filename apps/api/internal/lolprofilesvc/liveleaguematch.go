package lolprofilesvc

import (
	"fmt"
	"strings"

	"github.com/AhmetMuratAcar/winnable-lol/apps/api/internal/types"
)

func ToLiveLeagueGame(raw types.RawLiveResponse) types.LiveLeagueGame {
	res := types.LiveLeagueGame{
		QueueID:       raw.GameQueueConfigID,
		GameLength:    raw.GameLength,
		GameStartTime: raw.GameStartTime,
	}
	res.MatchID = fmt.Sprintf("%s_%d", raw.PlatformID, raw.GameID)

	for _, p := range raw.Participants {
		curr := types.LiveLeagueGameParticipant{
			PUUID:         p.Puuid,
			TeamID:        p.TeamID,
			Summoner1ID:   p.Spell1ID,
			Summoner2ID:   p.Spell2ID,
			ChampionID:    p.ChampionID,
			ProfileIconID: p.ProfileIconID,
		}

		riotIDParts := strings.Split(p.RiotID, "#")
		curr.GameName = riotIDParts[0]
		curr.TagLine = riotIDParts[1]

		curr.Runes.MainRuneID = p.Perks.PerkIds[0]
		curr.Runes.SubTreeID = p.Perks.PerkSubStyle

		res.Participants = append(res.Participants, curr)
	}

	for _, c := range raw.BannedChampions {
		curr := types.LiveLeagueGameBan{
			ChampionID: c.ChampionID,
			TeamID:     c.TeamID,
			PickTurn:   c.PickTurn,
		}

		res.Bans = append(res.Bans, curr)
	}

	return res
}
