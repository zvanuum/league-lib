package leaguelib

import (
	"log"
	"net/http"
	"net/url"
)

//TODO: Remove log statements!

// LeagueClient is a wrapper for contacting the Riot Games API
type LeagueClient struct {
	APIKey     string
	Client     *http.Client
	Region     string
	APIVersion int
}

// NewLeagueClient returns the reference to a new LeagueClient using the provided API key and region
func NewLeagueClient(apiKey string, region string) *LeagueClient {
	leagueClient := &LeagueClient{
		APIKey:     apiKey,
		Client:     &http.Client{},
		APIVersion: 3,
	}

	return leagueClient
}

func (leagueClient *LeagueClient) makeRequest(endpoint string, values url.Values) (*http.Request, error) {
	url := formatRequestURL(leagueClient.Region, "static-data", leagueClient.APIVersion, endpoint, values)
	query := url.Query()
	query.Add("api_key", leagueClient.APIKey)
	url.RawQuery = query.Encode()

	log.Printf("Creating request to: %s\n", url)

	return http.NewRequest("GET", url.String(), nil)
}
