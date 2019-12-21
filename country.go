package airvisual

import (
	"fmt"
	"net/url"
)

// Countries is an object returned from countries endpoint
type Countries struct {
	Status string           `json:"status"`
	Data   []*CountriesData `json:"data"`
}

// CountriesData contains information of a country
type CountriesData struct {
	Country string `json:"country"`
}

// Countries list supported countries
func (c *Client) Countries() (*Countries, error) {
	v := url.Values{}
	v.Add("key", c.APIKey)

	endpoint := c.endpoint(countriesEndpoint, v)

	var countries Countries
	err := c.request(endpoint, &countries)
	if err != nil {
		return &countries, fmt.Errorf("unable to list countries: %v", err)
	}

	return &countries, nil
}
