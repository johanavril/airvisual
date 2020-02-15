package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/johanavril/airvisual"
)

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
