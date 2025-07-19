package types

type RequestBody struct {
	Region   string `json:"region"`
	GameName string `json:"gameName"`
	TagLine  string `json:"tagLine"`
}
