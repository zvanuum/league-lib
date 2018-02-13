package leaguelib

import (
	"fmt"
	"log"
	"net/url"

	"github.com/zachvanuum/league-lib/model"
)

// GetChampions is used to retrieve static data regarding champions (title, id)
func (leagueClient *LeagueClient) GetChampions(values url.Values) model.Champions {
	req, err := leagueClient.makeRequest("champions", values)
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

func (leagueClient *LeagueClient) GetItems(values url.Values) model.Items {
	req, err := leagueClient.makeRequest(fmt.Sprintf("%s", "items"), values)
	if err != nil {
		log.Printf("Failed to create request: %s\n", err)
	}

	res, err := leagueClient.Client.Do(req)
	if err != nil {
		log.Printf("Failed to get response: %s\n", err)
	}
	defer res.Body.Close()

	var items model.Items
	err = unmarshalResponse(res, &items)
	if err != nil {
		log.Printf("%s", err)
	}

	return items
}

func (leagueClient *LeagueClient) GetItem(id int64, values url.Values) model.Item {
	req, err := leagueClient.makeRequest(fmt.Sprintf("%s/%d", "items", id), values)
	if err != nil {
		log.Printf("Failed to create request: %s\n", err)
	}

	res, err := leagueClient.Client.Do(req)
	if err != nil {
		log.Printf("Failed to get response: %s\n", err)
	}
	defer res.Body.Close()

	var item model.Item
	err = unmarshalResponse(res, &item)
	if err != nil {
		log.Printf("%s", err)
	}

	return item
}

func (leagueClient *LeagueClient) GetLanguageStrings(values url.Values) model.LanguageStrings {
	req, err := leagueClient.makeRequest("language-strings", values)
	if err != nil {
		log.Printf("Failed to create request: %s\n", err)
	}

	res, err := leagueClient.Client.Do(req)
	if err != nil {
		log.Printf("Failed to get response: %s\n", err)
	}
	defer res.Body.Close()

	var languageStrings model.LanguageStrings
	err = unmarshalResponse(res, &languageStrings)
	if err != nil {
		log.Printf("%s", err)
	}

	return languageStrings
}

func (leagueClient *LeagueClient) GetLanguages() []string {
	req, err := leagueClient.makeRequest("languages", nil)
	if err != nil {
		log.Printf("Failed to create request: %s\n", err)
	}

	res, err := leagueClient.Client.Do(req)
	if err != nil {
		log.Printf("Failed to get response: %s\n", err)
	}
	defer res.Body.Close()

	var languages []string
	err = unmarshalResponse(res, &languages)
	if err != nil {
		log.Printf("%s", err)
	}

	return languages
}

func (leagueClient *LeagueClient) GetMaps(values url.Values) model.MapData {
	req, err := leagueClient.makeRequest("maps", values)
	if err != nil {
		log.Printf("Failed to create request: %s\n", err)
	}

	res, err := leagueClient.Client.Do(req)
	if err != nil {
		log.Printf("Failed to get response: %s\n", err)
	}
	defer res.Body.Close()

	var maps model.MapData
	err = unmarshalResponse(res, &maps)
	if err != nil {
		log.Printf("%s", err)
	}

	return maps
}

func (leagueClient *LeagueClient) GetMasteries(values url.Values) model.MasteryList {
	req, err := leagueClient.makeRequest("masteries", values)
	if err != nil {
		log.Printf("Failed to create request: %s\n", err)
	}

	res, err := leagueClient.Client.Do(req)
	if err != nil {
		log.Printf("Failed to get response: %s\n", err)
	}
	defer res.Body.Close()

	var masteries model.MasteryList
	err = unmarshalResponse(res, &masteries)
	if err != nil {
		log.Printf("%s", err)
	}

	return masteries
}

func (leagueClient *LeagueClient) GetMastery(id int64, values url.Values) model.Mastery {
	req, err := leagueClient.makeRequest(fmt.Sprintf("%s/%d", "masteries", id), values)
	if err != nil {
		log.Printf("Failed to create request: %s\n", err)
	}

	res, err := leagueClient.Client.Do(req)
	if err != nil {
		log.Printf("Failed to get response: %s\n", err)
	}
	defer res.Body.Close()

	var mastery model.Mastery
	err = unmarshalResponse(res, &mastery)
	if err != nil {
		log.Printf("%s", err)
	}

	return mastery
}

func (leagueClient *LeagueClient) GetProfileIcons(values url.Values) model.ProfileIcons {
	req, err := leagueClient.makeRequest("masteries", values)
	if err != nil {
		log.Printf("Failed to create request: %s\n", err)
	}

	res, err := leagueClient.Client.Do(req)
	if err != nil {
		log.Printf("Failed to get response: %s\n", err)
	}
	defer res.Body.Close()

	var profileIcons model.ProfileIcons
	err = unmarshalResponse(res, &profileIcons)
	if err != nil {
		log.Printf("%s", err)
	}

	return profileIcons
}

func (leagueClient *LeagueClient) GetRealm() model.Realm {
	req, err := leagueClient.makeRequest("realms", nil)
	if err != nil {
		log.Printf("Failed to create request: %s\n", err)
	}

	res, err := leagueClient.Client.Do(req)
	if err != nil {
		log.Printf("Failed to get response: %s\n", err)
	}
	defer res.Body.Close()

	var realm model.Realm
	err = unmarshalResponse(res, &realm)
	if err != nil {
		log.Printf("%s", err)
	}

	return realm
}

func (leagueClient *LeagueClient) GetRunes(values url.Values) model.RuneList {
	req, err := leagueClient.makeRequest("runes", values)
	if err != nil {
		log.Printf("Failed to create request: %s\n", err)
	}

	res, err := leagueClient.Client.Do(req)
	if err != nil {
		log.Printf("Failed to get response: %s\n", err)
	}
	defer res.Body.Close()

	var runes model.RuneList
	err = unmarshalResponse(res, &runes)
	if err != nil {
		log.Printf("%s", err)
	}

	return runes
}
