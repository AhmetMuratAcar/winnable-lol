package riot

import (
	"net/http"
	"os"
	"time"

	"github.com/AhmetMuratAcar/winnable-lol/backend/internal/types"
)

type Client struct {
	httpClient *http.Client
	apiKey     string
}

// Default constructor
func NewClient() *Client {
	return NewClientWithHTTPClient(&http.Client{Timeout: 10 * time.Second})
}

// Constructor for test injection
func NewClientWithHTTPClient(httpClient *http.Client) *Client {
	return &Client{
		httpClient: httpClient,
		apiKey:     os.Getenv("RIOT_API_KEY"),
	}
}

func (c *Client) GetSummonerPUUID(req types.RequestBody) (puuid string, err error) {
	return "", nil
}

func (c *Client) GetSummonerMastery(puuid string) error {
	return nil
}

// func (c *Client) GetMatchData() error {
// 	return nil
// }
// func (c *Client) getMatchID() error {
// 	return nil
// }
