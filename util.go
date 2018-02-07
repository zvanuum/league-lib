package leaguelib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// RegionToServicePlatform maps from game region to service proxy
var RegionToServicePlatform = map[string]string{
	Brazil:        "br1",
	EuropeNE:      "eun1",
	EuropeW:       "euw1",
	Japan:         "jp1",
	Korea:         "kr",
	LatinAmericaN: "la1",
	LatinAmericaS: "la2",
	NorthAmerica:  "na1",
	Oceania:       "oc1",
	Turkey:        "tr1",
	Russia:        "ru",
	"PBE":         "pbe1",
}

func unmarshalResponse(res *http.Response, target interface{}) error {
	var body []byte
	var err error

	if body, err = ioutil.ReadAll(res.Body); err == nil {
		if err2 := json.Unmarshal(body, &target); err2 != nil {
			return fmt.Errorf("Failed unmarshalling response: %s", err2)
		}

		return nil
	}

	return fmt.Errorf("Failed to read response body: %s", err)
}

func formatRequestURL(region string, api string, version int, endpoint string, query url.Values) url.URL {
	var requestURL url.URL

	requestURL.Scheme = "https"
	requestURL.Host = getHostByRegion(region)
	requestURL.Path = fmt.Sprintf("/lol/%s/v%d/%s", api, version, endpoint)
	requestURL.RawQuery = query.Encode()

	return requestURL
}

func getHostByRegion(region string) string {
	service := RegionToServicePlatform[region]
	if service == "" {
		// Will default to North America if unknown region is passed in
		service = "na1"
	}

	return fmt.Sprintf("%s.api.riotgames.com", service)
}
