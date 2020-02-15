package airvisual

import (
	"fmt"
	"net/url"
)

// States contains name of supported State
type States struct {
	State string `json:"state"`
}

// States list supported states in the specified country
func (c *Client) States(country string) ([]*States, error) {
	v := url.Values{}
	v.Add("key", c.APIKey)
	v.Add("country", country)

	endpoint := c.endpoint(statesEndpoint, v)

	payload := struct {
		Status string    `json:"status"`
		Data   []*States `json:"data"`
	}{}
	err := c.request(endpoint, &payload)
	if err != nil {
		return nil, fmt.Errorf("unable to list states: %v", err)
	}
	if payload.Status != "success" {
		return nil, fmt.Errorf("unable to list states: %v", payload.Status)
	}

	return payload.Data, nil
}
