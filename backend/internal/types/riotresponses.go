package types

type AccountResponse struct {
	Puuid     string `json:"puuid"`
	GameName  string `json:"gameName"`
	TagLine   string `json:"tagLine"`
}

type ChampionMastery struct {
	ChampionID     int `json:"championId"`
    ChampionLevel  int `json:"championLevel"`
    ChampionPoints int `json:"championPoints"`
}