package leaguelib

import (
	"fmt"
	"log"
	"net/url"

	"github.com/zachvanuum/league-lib/model"
)

// GetChampions is used to retrieve static data regarding champions (title, id)
func (leagueClient *LeagueClient) GetChampions() model.Champions {
	req, err := leagueClient.makeRequest("champions", nil)
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

// GetChampion is used to retrieve static data for a champion based on an id parameter
func (leagueClient *LeagueClient) GetChampion(id int64, values url.Values) model.Champion {
	req, err := leagueClient.makeRequest(fmt.Sprintf("%s/%d", "champions", id), values)
	if err != nil {
		log.Printf("Failed to create request: %s\n", err)
	}

	res, err := leagueClient.Client.Do(req)
	if err != nil {
		log.Printf("Failed to get response: %s\n", err)
	}
	defer res.Body.Close()

	var champion model.Champion
	err = unmarshalResponse(res, &champion)
	if err != nil {
		log.Printf("%s", err)
	}

	return champion
}
