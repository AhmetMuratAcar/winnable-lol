package types

// This is the format of the initial request from the frontend
type RequestBody struct {
	Region   string `json:"region"`
	GameName string `json:"gameName"`
	TagLine  string `json:"tagLine"`
}
