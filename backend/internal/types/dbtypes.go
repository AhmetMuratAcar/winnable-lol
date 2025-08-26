package types

import "time"

type PUUIDCacheCheck struct {
	Found         bool
	PUUID         string
	Stale         bool
	IsPopulated   bool
	LastUpdated   time.Time
	ProfileIconID int
	Level         int
}

type CachedProfileCheckList struct {
	PUUID       bool
	ProfileIcon bool
	Level       bool
	Ranks       bool
	Masteries   bool
	Matches     bool
}

type MatchRow struct {
	MatchID         string
	EndOfGameResult string
	GameDurationSec int
	GameStart       time.Time
	GameVersion     string
	QueueID         int
	WinningTeam     int
	BansBlue        []int
	BansRed         []int
}

type MatchParticipantRow struct {
	MatchID              string
	PUUID                string
	ParticipantIndex     int
	Team                 int
	ChampionID           int
	ChampLevel           int
	Kills                int
	Deaths               int
	Assists              int
	GoldEarned           int
	TotalDamageToChamps  int
	TotalMinionsKilled   int
	VisionScore          int
	Items                []int
	Summoner1ID          int
	Summoner2ID          int
	TeamPosition         string
	RiotIDGameName       string
	RiotIDTagLine        string
	SummonerLevelAtMatch int
	ProfileIconAtMatch   int
}

type SummonerRow struct {
	PUUID              string
	Region             string
	GameName           string
	TagLine            string
	ProfileIconID      int
	SummonerLevel      int
	TotalMastery       int
	TotalMasteryPoints int
	ChampionsPlayed    int
	CreatedAt          time.Time
	UpdatedAt          time.Time
	IsPopulated        bool
}

type RankRow struct {
	PUUID        string
	QueueType    string
	Tier         string
	Division     string
	LeaguePoints int
	Wins         int
	Losses       int
}

type ChampionMasteryRow struct {
	PUUID          string
	ChampionID     int
	ChampionLevel  int
	ChampionPoints int
}
