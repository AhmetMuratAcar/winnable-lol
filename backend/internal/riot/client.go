package riot

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"winnable/internal/lolprofilesvc"
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

func (c *RiotClient) GetSummonerPUUID(reqBody types.RequestBody) (types.AccountResponse, error) {
	route := GetPuuidRegionRoute(reqBody.Region)
	endpoint := fmt.Sprintf(
		"https://%s.%s/riot/account/v1/accounts/by-riot-id/%s/%s",
		route,
		c.baseURL,
		url.PathEscape(reqBody.GameName),
		url.PathEscape(reqBody.TagLine),
	)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return types.AccountResponse{}, fmt.Errorf("error creating GetSummonerPUUID request: %w", err)
	}
	req.Header.Set("X-Riot-Token", c.apiKey)

	res, err := c.httpClient.Do(req)
	if err != nil {
		return types.AccountResponse{}, fmt.Errorf("GetSummonerPUUID API request failed: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(res.Body)

		err := &types.HTTPError{
			StatusCode: res.StatusCode,
			Body:       string(bodyBytes),
		}

		log.Printf("RIOT API ERROR: %v", err)
		return types.AccountResponse{}, err
	}

	var result types.AccountResponse
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return types.AccountResponse{}, fmt.Errorf("error decoding JSON: %w", err)
	}

	return result, nil
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

func (c *RiotClient) GetSummonerMatchIDs(puuid, region string, start int, count int) ([]string, error) {
	route, err := GetMatchDataRegionRoute(region)
	if err != nil {
		return nil, fmt.Errorf("provided region for GetSummonerMatchIDs does not exist: %w", err)
	}
	startStr := fmt.Sprintf("%d", start)
	endpoint := fmt.Sprintf(
		"https://%s.%s/lol/match/v5/matches/by-puuid/%s/ids?start=%s&count=%s",
		route,
		c.baseURL,
		puuid,
		startStr,
		strconv.Itoa(count),
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

func (c *RiotClient) GetMatchData(matchID, region string) (types.LeagueMatch, error) {
	route, err := GetMatchDataRegionRoute(region)
	if err != nil {
		return types.LeagueMatch{}, fmt.Errorf("provided region for GetMatchData does not exist: %w", err)
	}
	endpoint := fmt.Sprintf(
		"https://%s.%s/lol/match/v5/matches/%s",
		route,
		c.baseURL,
		matchID,
	)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return types.LeagueMatch{}, fmt.Errorf("error creating GetMatchData request: %w", err)
	}
	req.Header.Set("X-Riot-Token", c.apiKey)

	res, err := c.httpClient.Do(req)
	if err != nil {
		return types.LeagueMatch{}, fmt.Errorf("GetMatchData API request failed: %w", err)
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
		return types.LeagueMatch{}, err
	}

	var rawMatchData types.RawMatchResponse
	if err := json.NewDecoder((res.Body)).Decode(&rawMatchData); err != nil {
		return types.LeagueMatch{}, fmt.Errorf("error decoding JSON: %w", err)
	}

	result := lolprofilesvc.ToLeagueMatch(rawMatchData)
	return result, nil
}

func (c *RiotClient) GetSummonerIconAndLevel(puuid, region string) (int, int, error) {
	endpoint := fmt.Sprintf(
		"https://%s.%s/lol/summoner/v4/summoners/by-puuid/%s",
		region,
		c.baseURL,
		puuid,
	)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return -1, -1, fmt.Errorf("error creating GetSummonerIconAndLevel request: %w", err)
	}
	req.Header.Set("X-Riot-Token", c.apiKey)

	res, err := c.httpClient.Do(req)
	if err != nil {
		return -1, -1, fmt.Errorf("GetSummonerIconAndLevel API request failed: %w", err)
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
		return -1, -1, err
	}

	type summonerPartial struct {
		ProfileIconID int `json:"profileIconId"`
		SummonerLevel int `json:"summonerLevel"`
	}
	var result summonerPartial

	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return -1, -1, fmt.Errorf("error decoding response JSON: %w", err)
	}

	return result.ProfileIconID, result.SummonerLevel, nil
}

func (c *RiotClient) GetSummonerRanks(puuid, region string) ([]types.LeagueRank, error) {
	endpoint := fmt.Sprintf(
		"https://%s.%s/lol/league/v4/entries/by-puuid/%s",
		region,
		c.baseURL,
		puuid,
	)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating GetSummonerRank request: %w", err)
	}
	req.Header.Set("X-Riot-Token", c.apiKey)

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("GetSummonerRank API request failed: %w", err)
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

	var result []types.LeagueRank
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding JSON: %w", err)
	}

	return result, nil
}

func (c *RiotClient) GetLiveGame(puuid, region string) (types.LiveLeagueGame, error) {
	region = strings.ToLower(region)
	endpoint := fmt.Sprintf(
		"https://%s.%s/lol/spectator/v5/active-games/by-summoner/%s",
		region,
		c.baseURL,
		puuid,
	)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return types.LiveLeagueGame{}, fmt.Errorf("error creating GetLiveGame request: %w", err)
	}
	req.Header.Set("X-Riot-Token", c.apiKey)

	res, err := c.httpClient.Do(req)
	if err != nil {
		return types.LiveLeagueGame{}, fmt.Errorf("GetLiveGame API request failed: %w", err)
	}

	if res.StatusCode == http.StatusNotFound {
		return types.LiveLeagueGame{}, &types.RiotAPIError{
			StatusCode: res.StatusCode,
			Message:    "summoner not in a live game",
		}
	}

	if res.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(res.Body)
		return types.LiveLeagueGame{}, &types.RiotAPIError{
			StatusCode: res.StatusCode,
			Message:    string(body),
		}
	}

	var rawResponse types.RawLiveResponse
	if err := json.NewDecoder(res.Body).Decode(&rawResponse); err != nil {
		return types.LiveLeagueGame{}, fmt.Errorf("error decoding GetLiveGame API response: %w", err)
	}

	liveGame := lolprofilesvc.ToLiveLeagueGame(rawResponse)
	return liveGame, nil
}
