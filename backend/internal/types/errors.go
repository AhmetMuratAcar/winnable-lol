package types

import "fmt"

type HTTPError struct {
	StatusCode int
	Body       string
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("non-200 response: %d %s", e.StatusCode, e.Body)
}

type RiotAPIError struct {
	StatusCode int
	Message    string
}

func (e *RiotAPIError) Error() string {
	return fmt.Sprintf("Riot API error: %d - %s", e.StatusCode, e.Message)
}
