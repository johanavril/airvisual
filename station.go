package airvisual

import (
	"fmt"
	"net/url"
	"strconv"
)

// Stations contains information of a station
type Stations struct {
	Location *Location `json:"location"`
	Station  string    `json:"station"`
}

// Stations list supported active stations inside a specified city
func (c *Client) Stations(city, state, country string) ([]*Stations, error) {
	v := url.Values{}
	v.Add("key", c.APIKey)
	v.Add("country", country)
	v.Add("state", state)
	v.Add("city", city)

	endpoint := c.endpoint(stationsEndpoint, v)

	payload := struct {
		Status string      `json:"status"`
		Data   []*Stations `json:"data"`
	}{}
	err := c.request(endpoint, &payload)
	if err != nil {
		return nil, fmt.Errorf("unable to list stations: %v", err)
	}
	if payload.Status != "success" {
		return nil, fmt.Errorf("unable to list stations: %v", payload.Status)
	}

	return payload.Data, nil
}

// Station contains data regarding forecast of a specific station
type Station struct {
	Name      string      `json:"name"`
	City      string      `json:"city"`
	State     string      `json:"state"`
	Country   string      `json:"country"`
	Location  *Location   `json:"location"`
	Forecasts []*Forecast `json:"forecasts"`
	Current   *Current    `json:"current"`
	History   *History    `json:"history"`
}

// Station return specified station's data object
func (c *Client) Station(station, city, state, country string) (*Station, error) {
	v := url.Values{}
	v.Add("key", c.APIKey)
	v.Add("country", country)
	v.Add("state", state)
	v.Add("city", city)
	v.Add("station", station)

	endpoint := c.endpoint(stationEndpoint, v)

	payload := struct {
		Status string   `json:"status"`
		Data   *Station `json:"data"`
	}{}
	err := c.request(endpoint, &payload)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve station data: %v", err)
	}
	if payload.Status != "success" {
		return nil, fmt.Errorf("unable to retrieve station data: %v", payload.Status)
	}

	return payload.Data, nil
}

// NearestStationIP return nearest station's data using IP address geolocation
func (c *Client) NearestStationIP() (*Station, error) {
	v := url.Values{}
	v.Add("key", c.APIKey)

	endpoint := c.endpoint(nearestStationEndpoint, v)

	payload := struct {
		Status string   `json:"status"`
		Data   *Station `json:"data"`
	}{}
	err := c.request(endpoint, &payload)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve nearest station by IP address geolocation: %v", err)
	}
	if payload.Status != "success" {
		return nil, fmt.Errorf("unable to retrieve nearest station by IP address geolocation: %v", payload.Status)
	}

	return payload.Data, nil

}

// NearestStationGPS return nearest station's data using specified GPS coordinates
func (c *Client) NearestStationGPS(lat, lon float64) (*Station, error) {
	v := url.Values{}
	v.Add("key", c.APIKey)
	v.Add("lat", strconv.FormatFloat(lat, 'f', -1, 64))
	v.Add("lon", strconv.FormatFloat(lon, 'f', -1, 64))

	endpoint := c.endpoint(nearestStationEndpoint, v)

	payload := struct {
		Status string   `json:"status"`
		Data   *Station `json:"data"`
	}{}
	err := c.request(endpoint, &payload)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve nearest station by GPS coordinates: %v", err)
	}
	if payload.Status != "success" {
		return nil, fmt.Errorf("unable to retrieve nearest station by GPS coordinates: %v", payload.Status)
	}

	return payload.Data, nil
}
