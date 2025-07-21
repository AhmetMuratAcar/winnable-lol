package riot

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"winnable/internal/types"
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
		return "", fmt.Errorf("error creating GetSummonerPUUID request: %w", err)
	}
	req.Header.Set("X-Riot-Token", c.apiKey)

	res, err := c.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("GetSummonerPUUID API request failed: %w", err)
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

func (c *RiotClient) GetSummonerMastery(region, puuid string) (string, error) {
	region = strings.ToLower(region)
	baseEndpoint := "https://" + region + ".api.riotgames.com/lol/champion-mastery/v4/champion-masteries/by-puuid"
	endpoint := fmt.Sprintf(
		"%s/%s",
		baseEndpoint,
		puuid,
	)
	fmt.Printf("url: %s", endpoint)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return "", fmt.Errorf("error creating GetSummonerMastery request: %w", err)
	}
	req.Header.Set("X-Riot-Token", c.apiKey)

	res, err := c.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("GetSummonerMastery API request failed: %w", err)
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

	return "", nil
}
