package airvisual

import (
	"fmt"
	"net/url"
)

// Countries contains name of supported country
type Countries struct {
	Country string `json:"country"`
}

// Countries list supported countries
func (c *Client) Countries() ([]*Countries, error) {
	v := url.Values{}
	v.Add("key", c.APIKey)

	endpoint := c.endpoint(countriesEndpoint, v)

	payload := struct {
		Status string       `json:"status"`
		Data   []*Countries `json:"data"`
	}{}
	err := c.request(endpoint, &payload)
	if err != nil {
		return nil, fmt.Errorf("unable to list countries: %v", err)
	}
	if payload.Status != "success" {
		return nil, fmt.Errorf("unable to list countries: %v", payload.Status)
	}

	return payload.Data, nil
}
