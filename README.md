# airvisual
[![GoDoc](https://godoc.org/github.com/johanavril/airvisual?status.svg)](https://godoc.org/github.com/johanavril/airvisual)
[![Go Report Card](https://goreportcard.com/badge/github.com/johanavril/airvisual)](https://goreportcard.com/report/github.com/johanavril/airvisual)
[![Build Status](https://travis-ci.org/johanavril/airvisual.svg?branch=master)](https://travis-ci.org/johanavril/airvisual)
[![Coverage Status](https://coveralls.io/repos/github/johanavril/airvisual/badge.svg?branch=master)](https://coveralls.io/github/johanavril/airvisual?branch=master)

## Description
airvisual is a Go wrapper for [AirVisual's API](https://api-docs.airvisual.com/). This package goal is to assist development on air quality application using AirVisual's API

## Installation
```
go get github.com/johanavril/airvisual
```

## Usage

### Importing
This package can be used by adding the following import to your `.go` files.
```go
import "github.com/johanavril/airvisual"
```

### Example
```go
func main() {
	client := airvisual.New(
		"API KEY",
		airvisual.WithHTTPClient(&http.Client{Timeout: 30 * time.Second}),
	)

	countries, err := client.Countries()
	if err != nil {
		log.Panicf("Failed to get supported Countries: %v\n", err)
	}

	countryName := countries[36].Country
	states, err := client.States(countryName)
	if err != nil {
		log.Panicf("Failed to get supported States: %v\n", err)
	}

	stateName := states[2].State
	cities, err := client.Cities(stateName, countryName)
	if err != nil {
		log.Panicf("Failed to get supported Cities: %v\n", err)
	}

	cityName := cities[0].City
	city, err := client.City(cityName, stateName, countryName)
	if err != nil {
		log.Panicf("Failed to get supported City: %v\n", err)
	}

	c, err := json.MarshalIndent(city, "", "\t")
	if err != nil {
		log.Panicf("Failed to marshal City: %v\n", err)
	}

	fmt.Printf("%s", string(c))
}
```

## Contributing
We are looking for any kind of contribution to improve this package. Create an issue or make a pull request if you found any improvement opportunity or a problem. 