package airvisual

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	baseEndpoint = "https://api.airvisual.com"

	countriesEndpoint      = "/v2/countries"
	statesEndpoint         = "/v2/states"
	citiesEndpoint         = "/v2/cities"
	nearestCityEndpoint    = "/v2/nearest_city"
	cityEndpoint           = "/v2/city"
	stationsEndpoint       = "/v2/stations"
	nearestStationEndpoint = "/v2/nearest_station"
	stationEndpoint        = "/v2/station"
	cityRankingEndpoint    = "/v2/city_ranking"
)

func (c *Client) endpoint(api string, v url.Values) string {
	return c.baseEndpoint + api + "?" + v.Encode()
}

func (c *Client) request(endpoint string, result interface{}) error {
	response, err := c.client.Get(endpoint)
	if err != nil {
		return fmt.Errorf("failed to fetch %s: %v", endpoint, err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected HTTP status %s", response.Status)
	}

	err = json.NewDecoder(response.Body).Decode(result)
	if err != nil {
		return fmt.Errorf("cannot decode JSON: %v", err)
	}

	return nil
}
