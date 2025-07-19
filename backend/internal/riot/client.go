package riot

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/AhmetMuratAcar/winnable-lol/backend/internal/types"
)

type RiotClient struct {
	httpClient *http.Client
	apiKey     string
	baseURL    string
}

// Default constructor
func NewClient() *RiotClient {
	return NewClientWithHTTPClient(&http.Client{Timeout: 10 * time.Second})
}

// Constructor for test injection
func NewClientWithHTTPClient(httpClient *http.Client) *RiotClient {
	return &RiotClient{
		httpClient: httpClient,
		apiKey:     os.Getenv("RIOT_API_KEY"),
		baseURL:    "api.riotgames.com/",
	}
}

func (c *RiotClient) GetSummonerPUUID(reqBody types.RequestBody) (puuid string, err error) {
	baseEndpoint := "https://americas." + c.baseURL + "riot/account/v1/accounts/by-riot-id"
	endpoint := fmt.Sprintf(
		"%s/%s/%s",
		baseEndpoint,
		url.PathEscape(reqBody.GameName),
		url.PathEscape(reqBody.TagLine),
	)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("X-Riot-Token", c.apiKey)

	res, err := c.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("request failed: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(res.Body)
		return "", fmt.Errorf(
				"non-200 response: %s\n%s", 
				res.Status, 
				string(bodyBytes),
			)
	}

	var result types.AccountResponse
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("error decoding JSON: %w", err)
	}

	return result.Puuid, nil
}

func (c *RiotClient) GetSummonerMastery(puuid string) error {
	return nil
}

// func (c *Client) GetMatchData() error {
// 	return nil
// }
// func (c *Client) getMatchID() error {
// 	return nil
// }
