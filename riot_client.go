package leaguelib

import "net/http"

type LeagueClient struct {
	APIKey string
	Client http.Client
	// Region string
}

// , region string
func NewClient(apiKey string) *LeagueClient {
	leagueClient := &LeagueClient{
		APIKey: apiKey,
	}

	return leagueClient
}
