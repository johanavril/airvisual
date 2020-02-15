package airvisual

import (
	"fmt"
	"net/url"
	"strconv"
)

// Cities contains name of supported city
type Cities struct {
	City string `json:"city"`
}

// Cities list supported cities in the specified state
func (c *Client) Cities(state, country string) ([]*Cities, error) {
	v := url.Values{}
	v.Add("key", c.APIKey)
	v.Add("country", country)
	v.Add("state", state)

	endpoint := c.endpoint(citiesEndpoint, v)

	payload := struct {
		Status string    `json:"status"`
		Data   []*Cities `json:"data"`
	}{}
	err := c.request(endpoint, &payload)
	if err != nil {
		return nil, fmt.Errorf("unable to list cities: %v", err)
	}
	if payload.Status != "success" {
		return nil, fmt.Errorf("unable to list cities: %v", payload.Status)
	}

	return payload.Data, nil
}

// City contains data regarding forecast of a specific city
type City struct {
	City      string      `json:"city"`
	State     string      `json:"state"`
	Country   string      `json:"country"`
	Location  *Location   `json:"location"`
	Forecasts []*Forecast `json:"forecasts"`
	Current   *Current    `json:"current"`
	History   *History    `json:"history"`
}

// City return specified city's data object
func (c *Client) City(city, state, country string) (*City, error) {
	v := url.Values{}
	v.Add("key", c.APIKey)
	v.Add("country", country)
	v.Add("state", state)
	v.Add("city", city)

	endpoint := c.endpoint(cityEndpoint, v)

	payload := struct {
		Status string `json:"status"`
		Data   *City  `json:"data"`
	}{}
	err := c.request(endpoint, &payload)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve city data: %v", err)
	}
	if payload.Status != "success" {
		return nil, fmt.Errorf("unable to retrieve city data: %v", payload.Status)
	}

	return payload.Data, nil
}

// NearestCityIP return nearest city's data using IP address geolocation
func (c *Client) NearestCityIP() (*City, error) {
	v := url.Values{}
	v.Add("key", c.APIKey)

	endpoint := c.endpoint(nearestCityEndpoint, v)

	payload := struct {
		Status string `json:"status"`
		Data   *City  `json:"data"`
	}{}
	err := c.request(endpoint, &payload)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve nearest city by IP address geolocation: %v", err)
	}
	if payload.Status != "success" {
		return nil, fmt.Errorf("unable to retrieve nearest city by IP address geolocation: %v", payload.Status)
	}

	return payload.Data, nil
}

// NearestCityGPS return nearest city's data using specified GPS coordinates
func (c *Client) NearestCityGPS(lat, lon float64) (*City, error) {
	v := url.Values{}
	v.Add("key", c.APIKey)
	v.Add("lat", strconv.FormatFloat(lat, 'f', -1, 64))
	v.Add("lon", strconv.FormatFloat(lon, 'f', -1, 64))

	endpoint := c.endpoint(nearestCityEndpoint, v)

	payload := struct {
		Status string `json:"status"`
		Data   *City  `json:"data"`
	}{}
	err := c.request(endpoint, &payload)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve nearest city by GPS coordinates: %v", err)
	}
	if payload.Status != "success" {
		return nil, fmt.Errorf("unable to retrieve nearest city by GPS coordinates: %v", payload.Status)
	}

	return payload.Data, nil
}

// CityRanking contains ranking information of a city
type CityRanking struct {
	City    string   `json:"city"`
	State   string   `json:"state"`
	Country string   `json:"country"`
	Ranking *Ranking `json:"ranking"`
}

// CityRanking return sorted array of selected major cities in the world from highest to lowest AQI
func (c *Client) CityRanking() ([]*CityRanking, error) {
	v := url.Values{}
	v.Add("key", c.APIKey)

	endpoint := c.endpoint(cityRankingEndpoint, v)

	payload := struct {
		Status string         `json:"status"`
		Data   []*CityRanking `json:"data"`
	}{}
	err := c.request(endpoint, &payload)
	if err != nil {
		return nil, fmt.Errorf("unable to list city ranking: %v", err)
	}
	if payload.Status != "success" {
		return nil, fmt.Errorf("unable to list city ranking: %v", payload.Status)
	}

	return payload.Data, nil
}
