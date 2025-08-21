package riot

import (
	"context"
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

	"winnable/internal/types"
	"winnable/internal/utils"

	"github.com/jackc/pgx/v5/pgxpool"
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

		err := &types.HTTPError{
			StatusCode: res.StatusCode,
			Body:       string(bodyBytes),
		}

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

func (c *RiotClient) GetSummonerMatchIDs(puuid string, start int, count int) ([]string, error) {
	// TODO: actually route to nearest server instead of defaulting all to americas
	baseEndpoint := "https://americas." + c.baseURL + "/lol/match/v5/matches/by-puuid"
	startStr := fmt.Sprintf("%d", start)
	endpoint := fmt.Sprintf(
		"%s/%s/ids?start=%s&count=%s",
		baseEndpoint,
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

func (c *RiotClient) GetMatchData(matchID string) (types.LeagueMatch, error) {
	baseEndpoint := "https://americas." + c.baseURL + "/lol/match/v5/matches"
	endpoint := fmt.Sprintf("%s/%s", baseEndpoint, matchID)

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

	result := utils.ToLeagueMatch(rawMatchData)
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

func FillCacheGaps(
	checklist types.CachedProfileCheckList,
	profile *types.LeagueProfilePage,
	client *RiotClient,
	ctx context.Context,
	pool *pgxpool.Pool,
) error {
	detachedCtx, cancel := context.WithTimeout(
		context.WithoutCancel(ctx),
		5*time.Second,
	)
	defer cancel()

	var err error
	if !checklist.Masteries {
		// checklist.Masteries is only false if NO masteries are present
		championMasteries, err := client.GetSummonerMastery(profile.Region, profile.PUUID)
		if err != nil {
			log.Printf(
				"Error requesting masteries in FillCacheGaps:\nPUUID:%s\nError: %v",
				profile.PUUID,
				err,
			)
		}

		for _, c := range championMasteries {
			profile.MasteryData.TotalMastery += c.ChampionLevel
			profile.MasteryData.TotalMasteryPoints += c.ChampionPoints
		}
		profile.MasteryData.ChampionsPlayed = len(championMasteries)
		profile.MasteryData.ChampionMasteries = championMasteries

		// TODO: async update DB with this info
	}

	if !checklist.Matches {
		numCachedMatches := len(profile.MatchData)
		matchIdIndexMap := make(map[string]int)
		if numCachedMatches > 0 {
			for i, m := range profile.MatchData {
				matchIdIndexMap[m.MatchID] = i
			}
		}

		startIndex := 0
		count := 20
		matchIDs, err := client.GetSummonerMatchIDs(profile.PUUID, startIndex, count)
		if err != nil {
			log.Printf(
				"Error requesting past match IDs in FillCacheGaps: \nPUUID%s\nError: %v",
				profile.PUUID,
				err,
			)
		}

		res := make([]types.LeagueMatch, 0, 20)
		toAdd := make([]types.LeagueMatch, 0, 20)
		for _, id := range matchIDs {
			if index, ok := matchIdIndexMap[id]; ok {
				res = append(res, profile.MatchData[index])
			} else {
				matchData, err := client.GetMatchData(id)
				if err != nil {
					log.Printf(
						"Error fetching matchID %s in FillCacheGaps\nError: %v",
						id,
						err,
					)
					continue
				}

				res = append(res, matchData)
				toAdd = append(toAdd, matchData)
			}
		}
		profile.MatchData = res
		
		go func(batch []types.LeagueMatch) {
			if err := utils.AddMatchData(detachedCtx, pool, batch); err != nil {
				log.Printf("async AddMatchData error in FillCacheGaps: %v", err)
			}
		}(toAdd)
	}

	if !checklist.Ranks {
		profile.Ranks, err = client.GetSummonerRanks(profile.PUUID, profile.Region)
		if err != nil {
			log.Printf("error fetching summoner ranks in FillCacheGaps. Error: %v", err)
		}

		// TODO: async update DB with this info
	}

	if !checklist.ProfileIcon || !checklist.Level {
		profile.ProfileIconID, profile.Level, err = client.GetSummonerIconAndLevel(
			profile.PUUID,
			profile.Region,
		)
		if err != nil {
			log.Printf("error fetching summoner icon and level in FillCacheGaps. Error: %v", err)
		}

		// TODO: async update DB with this info
	}

	return nil
}
