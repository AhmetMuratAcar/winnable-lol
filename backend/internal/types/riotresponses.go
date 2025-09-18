package types

import "time"

type AccountResponse struct {
	Puuid    string `json:"puuid"`
	GameName string `json:"gameName"`
	TagLine  string `json:"tagLine"`
}

type ChampionMastery struct {
	ChampionID     int `json:"championId"`
	ChampionLevel  int `json:"championLevel"`
	ChampionPoints int `json:"championPoints"`
}

type MasteryData struct {
	ChampionMasteries  []ChampionMastery `json:"championMasteries"`
	TotalMastery       int               `json:"totalMastery"`
	TotalMasteryPoints int               `json:"totalMasteryPoints"`
	ChampionsPlayed    int               `json:"championsPlayed"`
}

type LeagueMatch struct {
	EndOfGameResult           string                   `json:"endOfGameResult"`
	GameDuration              int                      `json:"gameDuration"`
	GameEndedInEarlySurrender bool                     `json:"gameEndedInEarlySurrender"`
	GameStartTimestamp        int                      `json:"gameStartTimeStamp"`
	GameVersion               string                   `json:"gameVersion"`
	MatchID                   string                   `json:"matchId"`
	ParticipantPUUIDs         []string                 `json:"participantPUUIDs"`
	Participants              []LeagueMatchParticipant `json:"participants"`
	QueueId                   int                      `json:"queueId"`
	Bans                      [][]int                  `json:"bans"`
	WinningTeam               int                      `json:"winningTeam"`
}

type LeagueMatchParticipant struct {
	Assists                     int           `json:"assists"`
	ChampionID                  int           `json:"championId"`
	ChampLevel                  int           `json:"champLevel"`
	Deaths                      int           `json:"deaths"`
	GoldEarned                  int           `json:"goldEarned"`
	Items                       []int         `json:"items"`
	Kills                       int           `json:"kills"`
	ParticipantIndex            int           `json:"participantIndex"`
	ProfileIconID               int           `json:"profileIcon"`
	PUUID                       string        `json:"puuid"`
	RiotIDGameName              string        `json:"riotIdGameName"`
	RiotIDTagline               string        `json:"riotIdTagline"`
	Runes                       SummonerRunes `json:"runes"`
	Summoner1ID                 int           `json:"summoner1Id"`
	Summoner2ID                 int           `json:"summoner2Id"`
	SummonerLevel               int           `json:"summonerLevel"`
	Team                        int           `json:"team"`
	TeamPosition                string        `json:"teamPosition"`
	TotalDamageDealtToChampions int           `json:"totalDamageDealtToChampions"`
	TotalMinionsKilled          int           `json:"totalMinionsKilled"`
	VisionScore                 int           `json:"visionScore"`
}

type SummonerRunes struct {
	StatPerks     StatPerks         `json:"statPerks"`
	MainTree      MainRuneTree      `json:"mainTree"`
	SecondaryTree SecondaryRuneTree `json:"secondaryTree"`
}

type StatPerks struct {
	Defense int `json:"defense"`
	Flex    int `json:"flex"`
	Offense int `json:"offense"`
}

type MainRuneTree struct {
	Keystone int `json:"keystone"`
	Rune1    int `json:"rune1"`
	Rune2    int `json:"rune2"`
	Rune3    int `json:"rune3"`
}

type SecondaryRuneTree struct {
	Rune1 int `json:"rune1"`
	Rune2 int `json:"rune2"`
}

type LeagueRank struct {
	QueueType    string `json:"queueType"`
	Tier         string `json:"tier"`
	Rank         string `json:"rank"`
	LeaguePoints int    `json:"leaguePoints"`
	Wins         int    `json:"wins"`
	Losses       int    `json:"losses"`
}

type PlayedSummoner struct {
	GameName      string `json:"gameName"`
	TagLine       string `json:"tagLine"`
	GamesPlayed   int    `json:"gamesPlayed"`
	Wins          int    `json:"wins"`
	Losses        int    `json:"losses"`
	ProfileIconID int    `json:"profileIconID"`
}

type LeagueMatchSummary struct {
	ChampionID    int    `json:"championID"`
	OppChampionID int    `json:"oppChampionID"`
	Role          string `json:"role"`
	Kills         int    `json:"kills"`
	Deaths        int    `json:"deaths"`
	Assists       int    `json:"assists"`
	DidWin        bool   `json:"didWin"`
	QueueID       int    `json:"queueID"`
}

type WinLoss struct {
	Wins   int `json:"wins"`
	Losses int `json:"losses"`
}

type GamesSummary struct {
	MatchSummaries []LeagueMatchSummary `json:"matchSummaries"`
	KDAsByRole     map[string]float64   `json:"KDAsByRole"`
	RecordByRole   map[string]WinLoss   `json:"recordByRole"`
}

type LeagueProfilePage struct {
	PUUID         string           `json:"PUUID"`
	ProfileIconID int              `json:"profileIconId"`
	GameName      string           `json:"gameName"`
	TagLine       string           `json:"tagLine"`
	Region        string           `json:"region"`
	Level         int              `json:"summonerLevel"`
	LastUpdated   time.Time        `json:"lastUpdated"`
	Ranks         []LeagueRank     `json:"ranks"`
	MasteryData   MasteryData      `json:"masteryData"`
	MatchData     []LeagueMatch    `json:"matchData"`
	PlayedWith    []PlayedSummoner `json:"recentlyPlayedWith"`
	PlayedAgainst []PlayedSummoner `json:"recentlyPlayedAgainst"`
	RecentGames   GamesSummary     `json:"recentGames"`
}

// Removed the challenges and PlayerScore data from the RawMatchResponse struct because the
// Riot docs straight up lie and define some fields as int when they can sometimes
// be sent over as a float. Don't need it, just nuked it all to avoid the headache.
type RawMatchResponse struct {
	Metadata struct {
		DataVersion  string   `json:"dataVersion"`
		MatchID      string   `json:"matchId"`
		Participants []string `json:"participants"`
	} `json:"metadata"`
	Info struct {
		EndOfGameResult    string `json:"endOfGameResult"`
		GameCreation       int    `json:"gameCreation"`
		GameDuration       int    `json:"gameDuration"`
		GameEndTimestamp   int    `json:"gameEndTimestamp"`
		GameID             int    `json:"gameId"`
		GameMode           string `json:"gameMode"`
		GameName           string `json:"gameName"`
		GameStartTimestamp int    `json:"gameStartTimestamp"`
		GameType           string `json:"gameType"`
		GameVersion        string `json:"gameVersion"`
		MapID              int    `json:"mapId"`
		Participants       []struct {
			AllInPings                  int    `json:"allInPings"`
			AssistMePings               int    `json:"assistMePings"`
			Assists                     int    `json:"assists"`
			BaronKills                  int    `json:"baronKills"`
			BasicPings                  int    `json:"basicPings"`
			ChampExperience             int    `json:"champExperience"`
			ChampLevel                  int    `json:"champLevel"`
			ChampionID                  int    `json:"championId"`
			ChampionName                string `json:"championName"`
			ChampionTransform           int    `json:"championTransform"`
			CommandPings                int    `json:"commandPings"`
			ConsumablesPurchased        int    `json:"consumablesPurchased"`
			DamageDealtToBuildings      int    `json:"damageDealtToBuildings"`
			DamageDealtToObjectives     int    `json:"damageDealtToObjectives"`
			DamageDealtToTurrets        int    `json:"damageDealtToTurrets"`
			DamageSelfMitigated         int    `json:"damageSelfMitigated"`
			DangerPings                 int    `json:"dangerPings"`
			Deaths                      int    `json:"deaths"`
			DetectorWardsPlaced         int    `json:"detectorWardsPlaced"`
			DoubleKills                 int    `json:"doubleKills"`
			DragonKills                 int    `json:"dragonKills"`
			EligibleForProgression      bool   `json:"eligibleForProgression"`
			EnemyMissingPings           int    `json:"enemyMissingPings"`
			EnemyVisionPings            int    `json:"enemyVisionPings"`
			FirstBloodAssist            bool   `json:"firstBloodAssist"`
			FirstBloodKill              bool   `json:"firstBloodKill"`
			FirstTowerAssist            bool   `json:"firstTowerAssist"`
			FirstTowerKill              bool   `json:"firstTowerKill"`
			GameEndedInEarlySurrender   bool   `json:"gameEndedInEarlySurrender"`
			GameEndedInSurrender        bool   `json:"gameEndedInSurrender"`
			GetBackPings                int    `json:"getBackPings"`
			GoldEarned                  int    `json:"goldEarned"`
			GoldSpent                   int    `json:"goldSpent"`
			HoldPings                   int    `json:"holdPings"`
			IndividualPosition          string `json:"individualPosition"`
			InhibitorKills              int    `json:"inhibitorKills"`
			InhibitorTakedowns          int    `json:"inhibitorTakedowns"`
			InhibitorsLost              int    `json:"inhibitorsLost"`
			Item0                       int    `json:"item0"`
			Item1                       int    `json:"item1"`
			Item2                       int    `json:"item2"`
			Item3                       int    `json:"item3"`
			Item4                       int    `json:"item4"`
			Item5                       int    `json:"item5"`
			Item6                       int    `json:"item6"`
			ItemsPurchased              int    `json:"itemsPurchased"`
			KillingSprees               int    `json:"killingSprees"`
			Kills                       int    `json:"kills"`
			Lane                        string `json:"lane"`
			LargestCriticalStrike       int    `json:"largestCriticalStrike"`
			LargestKillingSpree         int    `json:"largestKillingSpree"`
			LargestMultiKill            int    `json:"largestMultiKill"`
			LongestTimeSpentLiving      int    `json:"longestTimeSpentLiving"`
			MagicDamageDealt            int    `json:"magicDamageDealt"`
			MagicDamageDealtToChampions int    `json:"magicDamageDealtToChampions"`
			MagicDamageTaken            int    `json:"magicDamageTaken"`
			NeedVisionPings             int    `json:"needVisionPings"`
			NeutralMinionsKilled        int    `json:"neutralMinionsKilled"`
			NexusKills                  int    `json:"nexusKills"`
			NexusLost                   int    `json:"nexusLost"`
			NexusTakedowns              int    `json:"nexusTakedowns"`
			ObjectivesStolen            int    `json:"objectivesStolen"`
			ObjectivesStolenAssists     int    `json:"objectivesStolenAssists"`
			OnMyWayPings                int    `json:"onMyWayPings"`
			ParticipantID               int    `json:"participantId"`
			PentaKills                  int    `json:"pentaKills"`
			Perks                       struct {
				StatPerks struct {
					Defense int `json:"defense"`
					Flex    int `json:"flex"`
					Offense int `json:"offense"`
				} `json:"statPerks"`
				Styles []struct {
					Description string `json:"description"`
					Selections  []struct {
						Perk int `json:"perk"`
						Var1 int `json:"var1"`
						Var2 int `json:"var2"`
						Var3 int `json:"var3"`
					} `json:"selections"`
					Style int `json:"style"`
				} `json:"styles"`
			} `json:"perks"`
			PhysicalDamageDealt            int    `json:"physicalDamageDealt"`
			PhysicalDamageDealtToChampions int    `json:"physicalDamageDealtToChampions"`
			PhysicalDamageTaken            int    `json:"physicalDamageTaken"`
			Placement                      int    `json:"placement"`
			PlayerAugment1                 int    `json:"playerAugment1"`
			PlayerAugment2                 int    `json:"playerAugment2"`
			PlayerAugment3                 int    `json:"playerAugment3"`
			PlayerAugment4                 int    `json:"playerAugment4"`
			PlayerAugment5                 int    `json:"playerAugment5"`
			PlayerAugment6                 int    `json:"playerAugment6"`
			PlayerSubteamID                int    `json:"playerSubteamId"`
			ProfileIcon                    int    `json:"profileIcon"`
			PushPings                      int    `json:"pushPings"`
			Puuid                          string `json:"puuid"`
			QuadraKills                    int    `json:"quadraKills"`
			RetreatPings                   int    `json:"retreatPings"`
			RiotIDGameName                 string `json:"riotIdGameName"`
			RiotIDTagline                  string `json:"riotIdTagline"`
			Role                           string `json:"role"`
			SightWardsBoughtInGame         int    `json:"sightWardsBoughtInGame"`
			Spell1Casts                    int    `json:"spell1Casts"`
			Spell2Casts                    int    `json:"spell2Casts"`
			Spell3Casts                    int    `json:"spell3Casts"`
			Spell4Casts                    int    `json:"spell4Casts"`
			SubteamPlacement               int    `json:"subteamPlacement"`
			Summoner1Casts                 int    `json:"summoner1Casts"`
			Summoner1ID                    int    `json:"summoner1Id"`
			Summoner2Casts                 int    `json:"summoner2Casts"`
			Summoner2ID                    int    `json:"summoner2Id"`
			SummonerID                     string `json:"summonerId"`
			SummonerLevel                  int    `json:"summonerLevel"`
			SummonerName                   string `json:"summonerName"`
			TeamEarlySurrendered           bool   `json:"teamEarlySurrendered"`
			TeamID                         int    `json:"teamId"`
			TeamPosition                   string `json:"teamPosition"`
			TimeCCingOthers                int    `json:"timeCCingOthers"`
			TimePlayed                     int    `json:"timePlayed"`
			TotalAllyJungleMinionsKilled   int    `json:"totalAllyJungleMinionsKilled"`
			TotalDamageDealt               int    `json:"totalDamageDealt"`
			TotalDamageDealtToChampions    int    `json:"totalDamageDealtToChampions"`
			TotalDamageShieldedOnTeammates int    `json:"totalDamageShieldedOnTeammates"`
			TotalDamageTaken               int    `json:"totalDamageTaken"`
			TotalEnemyJungleMinionsKilled  int    `json:"totalEnemyJungleMinionsKilled"`
			TotalHeal                      int    `json:"totalHeal"`
			TotalHealsOnTeammates          int    `json:"totalHealsOnTeammates"`
			TotalMinionsKilled             int    `json:"totalMinionsKilled"`
			TotalTimeCCDealt               int    `json:"totalTimeCCDealt"`
			TotalTimeSpentDead             int    `json:"totalTimeSpentDead"`
			TotalUnitsHealed               int    `json:"totalUnitsHealed"`
			TripleKills                    int    `json:"tripleKills"`
			TrueDamageDealt                int    `json:"trueDamageDealt"`
			TrueDamageDealtToChampions     int    `json:"trueDamageDealtToChampions"`
			TrueDamageTaken                int    `json:"trueDamageTaken"`
			TurretKills                    int    `json:"turretKills"`
			TurretTakedowns                int    `json:"turretTakedowns"`
			TurretsLost                    int    `json:"turretsLost"`
			UnrealKills                    int    `json:"unrealKills"`
			VisionClearedPings             int    `json:"visionClearedPings"`
			VisionScore                    int    `json:"visionScore"`
			VisionWardsBoughtInGame        int    `json:"visionWardsBoughtInGame"`
			WardsKilled                    int    `json:"wardsKilled"`
			WardsPlaced                    int    `json:"wardsPlaced"`
			Win                            bool   `json:"win"`
		} `json:"participants"`
		PlatformID string `json:"platformId"`
		QueueID    int    `json:"queueId"`
		Teams      []struct {
			Bans []struct {
				ChampionID int `json:"championId"`
				PickTurn   int `json:"pickTurn"`
			} `json:"bans"`
			Feats struct {
				EPICMONSTERKILL struct {
					FeatState int `json:"featState"`
				} `json:"EPIC_MONSTER_KILL"`
				FIRSTBLOOD struct {
					FeatState int `json:"featState"`
				} `json:"FIRST_BLOOD"`
				FIRSTTURRET struct {
					FeatState int `json:"featState"`
				} `json:"FIRST_TURRET"`
			} `json:"feats"`
			Objectives struct {
				Atakhan struct {
					First bool `json:"first"`
					Kills int  `json:"kills"`
				} `json:"atakhan"`
				Baron struct {
					First bool `json:"first"`
					Kills int  `json:"kills"`
				} `json:"baron"`
				Champion struct {
					First bool `json:"first"`
					Kills int  `json:"kills"`
				} `json:"champion"`
				Dragon struct {
					First bool `json:"first"`
					Kills int  `json:"kills"`
				} `json:"dragon"`
				Horde struct {
					First bool `json:"first"`
					Kills int  `json:"kills"`
				} `json:"horde"`
				Inhibitor struct {
					First bool `json:"first"`
					Kills int  `json:"kills"`
				} `json:"inhibitor"`
				RiftHerald struct {
					First bool `json:"first"`
					Kills int  `json:"kills"`
				} `json:"riftHerald"`
				Tower struct {
					First bool `json:"first"`
					Kills int  `json:"kills"`
				} `json:"tower"`
			} `json:"objectives"`
			TeamID int  `json:"teamId"`
			Win    bool `json:"win"`
		} `json:"teams"`
		TournamentCode string `json:"tournamentCode"`
	} `json:"info"`
}
