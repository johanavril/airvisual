package airvisual

import (
	"fmt"
	"net/url"
	"strconv"
)

// Stations is an object returned from stations endpoint
type Stations struct {
	Status string          `json:"status"`
	Data   []*StationsData `json:"data"`
}

// StationsData contains information of a station
type StationsData struct {
	Location *Location `json:"location"`
	Station  string    `json:"station"`
}

// Stations list supported active stations inside a specified city
func (c *Client) Stations(city, state, country string) (*Stations, error) {
	v := url.Values{}
	v.Add("key", c.APIKey)
	v.Add("country", country)
	v.Add("state", state)
	v.Add("city", city)

	endpoint := c.endpoint(stationsEndpoint, v)

	var stations Stations
	err := c.request(endpoint, &stations)
	if err != nil {
		return &stations, fmt.Errorf("unable to list supported stations: %v", err)
	}

	return &stations, nil
}

// Station is an object returned from station endpoint
type Station struct {
	Status string       `json:"status"`
	Data   *StationData `json:"data"`
}

// StationData contains data regarding forecast of a specific station
type StationData struct {
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

	var st Station
	err := c.request(endpoint, &st)
	if err != nil {
		return &st, fmt.Errorf("unable to retrieve station data: %v", err)
	}

	return &st, nil
}

// NearestStation is an object returned from nearest station endpoint
type NearestStation struct {
	Status string       `json:"status"`
	Data   *StationData `json:"data"`
}

// NearestStationIP return nearest station's data using IP address geolocation
func (c *Client) NearestStationIP() (*NearestStation, error) {
	v := url.Values{}
	v.Add("key", c.APIKey)

	endpoint := c.endpoint(nearestStationEndpoint, v)

	var nearestStation NearestStation
	err := c.request(endpoint, &nearestStation)
	if err != nil {
		return &nearestStation, fmt.Errorf("unable to retrieve nearest station by IP address geolocation: %v", err)
	}

	return &nearestStation, nil

}

// NearestStationGPS return nearest station's data using specified GPS coordinates
func (c *Client) NearestStationGPS(lat, lon float64) (*NearestStation, error) {
	v := url.Values{}
	v.Add("key", c.APIKey)
	v.Add("lat", strconv.FormatFloat(lat, 'f', -1, 64))
	v.Add("lon", strconv.FormatFloat(lon, 'f', -1, 64))

	endpoint := c.endpoint(nearestStationEndpoint, v)

	var nearestStation NearestStation
	err := c.request(endpoint, &nearestStation)
	if err != nil {
		return &nearestStation, fmt.Errorf("unable to retrieve nearest station by GPS coordinates: %v", err)
	}

	return &nearestStation, nil
}
