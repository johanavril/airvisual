package airvisual

import (
	"fmt"
	"net/url"
	"strconv"
)

// Cities is an object returned from cities endpoint
type Cities struct {
	Status string        `json:"status"`
	Data   []*CitiesData `json:"data"`
}

// CitiesData contains information of a city
type CitiesData struct {
	City string `json:"city"`
}

// Cities list supported cities in the specified state
func (c *Client) Cities(state, country string) (*Cities, error) {
	v := url.Values{}
	v.Add("key", c.APIKey)
	v.Add("country", country)
	v.Add("state", state)

	endpoint := c.endpoint(citiesEndpoint, v)

	var cities Cities
	err := c.request(endpoint, &cities)
	if err != nil {
		return &cities, fmt.Errorf("unable to list cities: %v", err)
	}

	return &cities, nil
}

// City is an object returned from city endpoint
type City struct {
	Status string    `json:"status"`
	Data   *CityData `json:"data"`
}

// CityData contains data regarding forecast of a specific city
type CityData struct {
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

	var ci City
	err := c.request(endpoint, &ci)
	if err != nil {
		return &ci, fmt.Errorf("unable to retrieve city data: %v", err)
	}

	return &ci, nil
}

// NearestCity is an object returned from nearest city endpoint
type NearestCity struct {
	Status string    `json:"status"`
	Data   *CityData `json:"data"`
}

// NearestCityIP return nearest city's data using IP address geolocation
func (c *Client) NearestCityIP() (*NearestCity, error) {
	v := url.Values{}
	v.Add("key", c.APIKey)

	endpoint := c.endpoint(nearestCityEndpoint, v)

	var nearestCity NearestCity
	err := c.request(endpoint, &nearestCity)
	if err != nil {
		return &nearestCity, fmt.Errorf("unable to retrieve nearest city by IP address geolocation: %v", err)
	}

	return &nearestCity, nil
}

// NearestCityGPS return nearest city's data using specified GPS coordinates
func (c *Client) NearestCityGPS(lat, lon float64) (*NearestCity, error) {
	v := url.Values{}
	v.Add("key", c.APIKey)
	v.Add("lat", strconv.FormatFloat(lat, 'f', -1, 64))
	v.Add("lon", strconv.FormatFloat(lon, 'f', -1, 64))

	endpoint := c.endpoint(nearestCityEndpoint, v)

	var nearestCity NearestCity
	err := c.request(endpoint, &nearestCity)
	if err != nil {
		return &nearestCity, fmt.Errorf("unable to retrieve nearest city by GPS coordinates: %v", err)
	}

	return &nearestCity, nil
}

// CityRanking is an object returned from city ranking endpoint
type CityRanking struct {
	Status string             `json:"status"`
	Data   []*CityRankingData `json:"data"`
}

// CityRankingData contains ranking information of a city
type CityRankingData struct {
	City    string   `json:"city"`
	State   string   `json:"state"`
	Country string   `json:"country"`
	Ranking *Ranking `json:"ranking"`
}

// CityRanking return sorted array of selected major cities in the world from highest to lowest AQI
func (c *Client) CityRanking() (*CityRanking, error) {
	v := url.Values{}
	v.Add("key", c.APIKey)

	endpoint := c.endpoint(cityRankingEndpoint, v)

	var cityRanking CityRanking
	err := c.request(endpoint, &cityRanking)
	if err != nil {
		return &cityRanking, fmt.Errorf("unable to list city ranking: %v", err)
	}

	return &cityRanking, nil
}
