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

type LeagueMatch struct {
	EndOfGameResult    string                   `json:"endOfGameResult"`
	GameDuration       int                      `json:"gameDuration"`
	GameStartTimestamp int                      `json:"gameStartTimeStamp"`
	MatchID            string                   `json:"matchId"`
	ParticipantPUUIDs  []string                 `json:"participantPUUIDs"`
	Participants       []LeagueMatchParticipant `json:"participants"`
	QueueId            int                      `json:"queueId"`
	Bans               []int                    `json:"bans"`
}

type LeagueMatchParticipant struct {
	Assists                     int    `json:"assists"`
	ChampionID                  int    `json:"championId"`
	ChampLevel                  int    `json:"champLevel"`
	Deaths                      int    `json:"deaths"`
	GoldEarned                  int    `json:"goldEarned"`
	Items                       []int  `json:"items"`
	Kills                       int    `json:"kills"`
	ProfileIconID               int    `json:"profileIcon"`
	PUUID                       string `json:"puuid"`
	RiotIDGameName              string `json:"riotIdGameName"`
	RiotIDTagline               string `json:"riotIdTagline"`
	Summoner1ID                 int    `json:"summoner1Id"`
	Summoner2ID                 int    `json:"summoner2Id"`
	SummonerLevel               int    `json:"summonerLevel"`
	TeamPosition                string `json:"teamPosition"`
	TotalDamageDealtToChampions int    `json:"totalDamageDealtToChampions"`
	TotalMinionsKilled          int    `json:"totalMinionsKilled"`
	VisionScore                 int    `json:"visionScore"`
}

type LeagueRank struct {
	QueueType    string `json:"queueType"`
	Tier         string `json:"tier"`
	Rank         string `json:"rank"`
	LeaguePoints int    `json:"leaguePoints"`
	Wins         int    `json:"wins"`
	Losses       int    `json:"losses"`
}

type LeagueProfilePage struct {
	ProfileIconID int           `json:"profileIconId"`
	GameName      string        `json:"gameName"`
	TagLine       string        `json:"tagLine"`
	Region        string        `json:"region"`
	Level         int           `json:"summonerLevel"`
	Ranks         []LeagueRank  `json:"ranks"`
	MasteryData   MasteryData   `json:"masteryData"`
	MatchData     []LeagueMatch `json:"lastGames"`
}
