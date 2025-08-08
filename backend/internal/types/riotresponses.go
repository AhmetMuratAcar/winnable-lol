package types

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

// TODO: figure out exactly what to grab from riot's API for each match
type LeagueMatch struct{}
type LeagueProfilePage struct {
	ProfileIconId int           `json:"profileIconId"`
	GameName      string        `json:"gameName"`
	TagLine       string        `json:"tagLine"`
	Region        string        `json:"region"`
	Level         int           `json:"summonerLevel"`
	Rank          string        `json:"rank"`
	MasteryData   MasteryData   `json:"masteryData"`
	LastGames     []LeagueMatch `json:"lastGames"`
}
