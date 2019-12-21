# airvisual
[![GoDoc](https://godoc.org/github.com/johanavril/airvisual?status.svg)](https://godoc.org/github.com/johanavril/airvisual)

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
	if countries.Status != "success" {
		log.Panicf("Countries API reject the request: %s\n", countries.Status)
	}

	countryName := countries.Data[36].Country
	states, err := client.States(countryName)
	if err != nil {
		log.Panicf("Failed to get supported States: %v\n", err)
	}
	if states.Status != "success" {
		log.Panicf("States API reject the request: %s\n", states.Status)
	}

	stateName := states.Data[2].State
	cities, err := client.Cities(stateName, countryName)
	if err != nil {
		log.Panicf("Failed to get supported Cities: %v\n", err)
	}
	if cities.Status != "success" {
		log.Panicf("Cities API reject the request: %s\n", cities.Status)
	}

	cityName := cities.Data[0].City
	city, err := client.City(cityName, stateName, countryName)
	if err != nil {
		log.Panicf("Failed to get supported City: %v\n", err)
	}
	if city.Status != "success" {
		log.Panicf("City API reject the request: %s\n", city.Status)
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