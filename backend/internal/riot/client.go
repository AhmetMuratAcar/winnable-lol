package riot

import (
	"net/http"
	"os"
	"time"
)

type Client struct {
	httpClient *http.Client
	apiKey     string
}

func NewClient() *Client {
	return &Client {
		httpClient: &http.Client{Timeout: 10 * time.Second},
		apiKey: os.Getenv("RIOT_API_KEY"),
	}
}
