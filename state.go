package airvisual

import (
	"fmt"
	"net/url"
)

// States is an object returned from states endpoint
type States struct {
	Status string        `json:"status"`
	Data   []*StatesData `json:"data"`
}

// StatesData contains information of a state
type StatesData struct {
	State string `json:"state"`
}

// States list supported states in the specified country
func (c *Client) States(country string) (*States, error) {
	v := url.Values{}
	v.Add("key", c.APIKey)
	v.Add("country", country)

	endpoint := c.endpoint(statesEndpoint, v)

	var states States
	err := c.request(endpoint, &states)
	if err != nil {
		return &states, fmt.Errorf("unable to list states: %v", err)
	}

	return &states, nil
}
