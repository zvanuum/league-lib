package leaguelib

import (
	"log"
	"net/http"

	"github.com/zachvanuum/league-lib/model"
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

func (leagueClient *LeagueClient) makeRequest(endpoint string) (*http.Request, error) {
	url := formatRequestURL(leagueClient.Region, leagueClient.APIVersion, endpoint, nil)
	query := url.Query()
	query.Add("api_key", leagueClient.APIKey)
	url.RawQuery = query.Encode()

	log.Printf("Creating request to: %s\n", url)

	return http.NewRequest("GET", url.String(), nil)
}

func (leagueClient *LeagueClient) GetChampions() model.Champions {
	req, err := leagueClient.makeRequest("champions")
	if err != nil {
		log.Printf("Failed to create request: %s\n", err)
	}

	res, err := leagueClient.Client.Do(req)
	if err != nil {
		log.Printf("Failed to get response: %s\n", err)
	}
	defer res.Body.Close()

	var champions model.Champions
	err = unmarshalResponse(res, &champions)
	if err != nil {
		log.Printf("%s", err)
	}

	return champions
}
