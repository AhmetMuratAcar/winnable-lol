package riot

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
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
		baseURL:    "api.riotgames.com",
	}
}

func (c *RiotClient) GetSummonerPUUID(reqBody types.RequestBody) (puuid string, err error) {
	// TODO: actually route to nearest server instead of defaulting all to americas
	baseEndpoint := "https://americas." + c.baseURL + "/riot/account/v1/accounts/by-riot-id"
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
		err := fmt.Errorf(
			"non-200 response: %s\n%s",
			res.Status,
			string(bodyBytes),
		)

		log.Printf("RIOT API ERROR: %v", err)
		return "", err
	}

	var result types.AccountResponse
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("error decoding JSON: %w", err)
	}

	return result.Puuid, nil
}

func (c *RiotClient) GetSummonerMastery(region, puuid string) ([]types.ChampionMastery, error) {
	region = strings.ToLower(region)
	baseEndpoint := "https://" + region + "." + c.baseURL + "/lol/champion-mastery/v4/champion-masteries/by-puuid"
	endpoint := fmt.Sprintf(
		"%s/%s",
		baseEndpoint,
		puuid,
	)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating GetSummonerMastery request: %w", err)
	}
	req.Header.Set("X-Riot-Token", c.apiKey)

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("GetSummonerMastery API request failed: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(res.Body)
		err := fmt.Errorf(
			"non-200 response: %s\n%s",
			res.Status,
			string(bodyBytes),
		)

		log.Printf("RIOT API ERROR: %v", err)
		return nil, err
	}

	var championMasteries []types.ChampionMastery
	if err := json.NewDecoder(res.Body).Decode(&championMasteries); err != nil {
		return nil, fmt.Errorf("error decoding JSON: %w", err)
	}

	return championMasteries, nil
}

func (c *RiotClient) GetSummonerMatchIDs(puuid string, start int) ([]string, error) {
	// TODO: actually route to nearest server instead of defaulting all to americas
	baseEndpoint := "https://americas." + c.baseURL + "/lol/match/v5/matches/by-puuid"
	count := "20"
	startStr := fmt.Sprintf("%d", start)
	endpoint := fmt.Sprintf(
		"%s/%s/ids?start=%s&count=%s",
		baseEndpoint,
		puuid,
		startStr,
		count,
	)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating GetSummonerMatchIDs request: %w", err)
	}
	req.Header.Set("X-Riot-Token", c.apiKey)

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("GetSummonerMatchIDs API request failed: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(res.Body)
		err := fmt.Errorf(
			"non-200 response: %s\n%s",
			res.Status,
			string(bodyBytes),
		)

		log.Printf("RIOT API ERROR: %v", err)
		return nil, err
	}

	var result []string
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding JSON: %w", err)
	}

	return result, nil
}

func (c *RiotClient) GetMatchData(matchID string) (types.LeagueMatch, error) {
	var result types.LeagueMatch
	return result, nil 
}